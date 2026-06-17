package linters

import (
	"fmt"
	"go/ast"
	"go/token"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	Arrange = "Arrange"
	Act     = "Act"
	Assert  = "Assert"
)

func run(pass *analysis.Pass) (any, error) {
	testFiles := findTestFunctions(pass)

	// TODO: Cleanup

	for testFile, testFuncs := range testFiles {
		for _, testFunc := range testFuncs {
			for _, testBlock := range findTestBlocks(pass, testFunc) {
				// Empty test, not our problem
				if len(testBlock.List) == 0 {
					continue
				}

				actCandidates := actCandidates(testFunc.Name.Name)

				actIndex, ok := findAct(pass, testBlock.List, actCandidates)
				if !ok {
					// If no act was found, we assume we just couldn't figure out what the test was about, so we don't report anything
					continue
				}

				actStatement := testBlock.List[actIndex]

				// Notice: This does include the // Act itself
				testComments := findCommentsBetweenLines(pass, testFile.Comments, testBlock.Pos(), testBlock.End())
				commentsBeforeAct := findCommentsBetweenLines(pass, testComments, testBlock.Pos(), actStatement.End())
				commentsAfterAct := findCommentsBetweenLines(pass, testComments, actStatement.End(), testBlock.End())

				detectDuplicates(pass, testComments, testBlock.Pos(), testBlock.End(), Arrange)
				detectDuplicates(pass, testComments, testBlock.Pos(), testBlock.End(), Act)
				detectDuplicates(pass, testComments, testBlock.Pos(), testBlock.End(), Assert)

				// Check if any comments are before the act
				if len(commentsBeforeAct) == 0 || !hasLineWithPrefix(commentsBeforeAct[len(commentsBeforeAct)-1].Text(), Act) {
					pass.Report(analysis.Diagnostic{
						Pos:     actStatement.Pos(),
						End:     actStatement.End(),
						Message: fmt.Sprintf("// %s statement expected", Act),
					})
				}

				arrangeExceptions, ok := isArrangeRequired(pass, actIndex, testBlock.List)
				if ok {
					if len(commentsBeforeAct) < 1 || !hasLineWithPrefix(commentsBeforeAct[0].Text(), Arrange) {
						pass.Report(analysis.Diagnostic{
							Pos:     testBlock.List[arrangeExceptions].Pos(),
							End:     testBlock.List[arrangeExceptions].End(),
							Message: fmt.Sprintf("// %s statement expected", Arrange),
						})
					}
				}

				// Check if statements exist after the Act
				if len(testBlock.List) > actIndex+1 {
					afterActStatement := testBlock.List[actIndex+1]
					if len(commentsAfterAct) == 0 || !hasLineWithPrefix(commentsAfterAct[0].Text(), Assert) {
						pass.Report(analysis.Diagnostic{
							Pos:     afterActStatement.Pos(),
							End:     afterActStatement.End(),
							Message: fmt.Sprintf("// %s statement expected", Assert),
						})
					}
				} else {
					// If not, something is wrong
					pass.Report(analysis.Diagnostic{
						Pos:     actStatement.Pos(),
						End:     actStatement.End(),
						Message: fmt.Sprintf("// %s statement expected", Assert),
					})
				}
			}
		}
	}

	return nil, nil
}

// detectDuplicates looks for duplicate statements
func detectDuplicates(pass *analysis.Pass, comments []*ast.CommentGroup, from token.Pos, to token.Pos, prefix string) {
	var count int

	for _, comment := range comments {
		if hasLineWithPrefix(comment.Text(), prefix) {
			count++
		}
	}

	if count > 1 {
		pass.Report(analysis.Diagnostic{
			Pos:     from,
			End:     to,
			Message: fmt.Sprintf("Duplicate %s statement", prefix),
		})
	}
}

// isArrangeRequired determines whether the code should have an arrange statement. This
// statement is only necessary if there is non-setup code above the Act statement. The first
// return value indicates how many statements are considered exceptions.
func isArrangeRequired(pass *analysis.Pass, actIndex int, statements []ast.Stmt) (int, bool) {
	for index, statement := range statements {
		if isArrangeException(pass, statement) {
			continue
		}

		if index == actIndex {
			return 0, false
		}

		return index, true
	}

	return 0, false
}

// findCommentsBetweenLines filters comments between the 2 given positions, used primarily to split
// comments between the start, act and end of a test
func findCommentsBetweenLines(pass *analysis.Pass, comments []*ast.CommentGroup, from token.Pos, to token.Pos) []*ast.CommentGroup {
	result := make([]*ast.CommentGroup, 0, len(comments))

	fromLine := pass.Fset.Position(from).Line
	toLine := pass.Fset.Position(to).Line

	for _, comment := range comments {
		commentLine := pass.Fset.Position(comment.Pos()).Line
		if commentLine > fromLine && commentLine < toLine {
			result = append(result, comment)
		}
	}

	return result
}

// findTestFunctions finds all the functions that run tests
func findTestFunctions(pass *analysis.Pass) map[*ast.File][]*ast.FuncDecl {
	tests := make(map[*ast.File][]*ast.FuncDecl, len(pass.Files))

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			funcNode, ok := node.(*ast.FuncDecl)
			if !ok {
				return true
			}

			// Skip non-test functions
			if !strings.HasPrefix(funcNode.Name.Name, "Test") {
				return true
			}

			tests[file] = append(tests[file], funcNode)

			return true
		})
	}

	return tests
}

// findTestBlocks attempts to find all test blocks, which could either be the function block itself
// or multiple subtests
func findTestBlocks(pass *analysis.Pass, testFunc *ast.FuncDecl) []*ast.BlockStmt {
	subTests, ok := findSubTests(pass, testFunc.Body.List)
	if ok {
		return subTests
	}

	return []*ast.BlockStmt{testFunc.Body}
}

// findSubTests recursively walks through the given statements hunting for subtests, if any are found they
// are returned. The boolean indicates if any were found.
func findSubTests(pass *analysis.Pass, statements []ast.Stmt) ([]*ast.BlockStmt, bool) {
	var result []*ast.BlockStmt

	for _, statement := range statements {
		ast.Inspect(statement, func(node ast.Node) bool {
			callExp, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			selectorExp, ok := callExp.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// Look for calls to the function Run on a suite/t
			if selectorExp.Sel.Name != "Run" || !isSubtestRun(pass, selectorExp.X) {
				return true
			}

			// If the second argument isn't a function then it might just be nil, which is questionable but outside the scope of this linter
			subTestFunction, ok := callExp.Args[1].(*ast.FuncLit)
			if !ok {
				return true
			}

			subTests, ok := findSubTests(pass, subTestFunction.Body.List)
			if ok {
				result = append(result, subTests...)
				return true
			}

			result = append(result, subTestFunction.Body)

			return true
		})
	}

	return result, len(result) > 0
}

func findAct(pass *analysis.Pass, statements []ast.Stmt, actCandidates [][]string) (int, bool) {
	foundSelectors := make(map[int]*ast.SelectorExpr)
	foundIdentifiers := make(map[int]*ast.Ident)

	for index, statement := range statements {
		ast.Inspect(statement, func(node ast.Node) bool {
			callExp, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			identifier, ok := callExp.Fun.(*ast.Ident)
			if ok {
				foundIdentifiers[index] = identifier
				return true
			}

			selector, ok := callExp.Fun.(*ast.SelectorExpr)
			if ok {
				foundSelectors[index] = selector
				return true
			}

			return true
		})
	}

	for _, candidateGroup := range actCandidates {
		for statementIndex, identifier := range foundIdentifiers {
			if slices.Contains(candidateGroup, strings.ToLower(identifier.Name)) {
				return statementIndex, true
			}
		}

		for statementIndex, selector := range foundSelectors {
			if slices.Contains(candidateGroup, strings.ToLower(selector.Sel.Name)) {
				return statementIndex, true
			}
		}
	}

	return 0, false
}

// isArrangeException returns whether the given statement is allowed to be outside of the Arrange statement
func isArrangeException(pass *analysis.Pass, statement ast.Stmt) bool {
	return isTParallel(pass, statement) || isTAssign(pass, statement)
}

// isTAssign returns whether the given statement equals an assignment to a variable of type *testing.T
func isTAssign(pass *analysis.Pass, statement ast.Stmt) bool {
	var isCall bool

	ast.Inspect(statement, func(node ast.Node) bool {
		assignExp, ok := node.(*ast.AssignStmt)
		if !ok {
			return false
		}

		if len(assignExp.Lhs) != 1 {
			return false
		}

		varIdentifier, ok := assignExp.Lhs[0].(*ast.Ident)
		if !ok {
			return false
		}

		identifierType, ok := pass.TypesInfo.Defs[varIdentifier]
		if !ok {
			return false
		}

		if identifierType == nil || identifierType.Type() == nil {
			return false
		}

		isCall = identifierType.Type().String() == "*testing.T"

		return false
	})

	return isCall
}

// isTParallel returns whether the given statement is a call to testing.T.Parallel
func isTParallel(pass *analysis.Pass, statement ast.Stmt) bool {
	var isCall bool

	ast.Inspect(statement, func(node ast.Node) bool {
		callExp, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		selectorExp, ok := callExp.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		typeT, ok := pass.TypesInfo.Types[selectorExp.X]
		if !ok {
			return true
		}

		isCall = selectorExp.Sel.Name == "Parallel" && typeT.Type.String() == "*testing.T"

		return false
	})

	return isCall
}

// isSubtestRun checks whether the expression is either:
//
// - testing.T.Run
// - github.com/stretchr/testify/suite.T.Run
func isSubtestRun(pass *analysis.Pass, expression ast.Expr) bool {
	return isTestingT(pass, expression) || isSuite(pass, expression)
}

// isTestingT checks whether the expression is testing.T.Run
func isTestingT(pass *analysis.Pass, expression ast.Expr) bool {
	typeT, ok := pass.TypesInfo.Types[expression]
	if !ok {
		return true
	}

	return typeT.Type.String() == "*testing.T"
}

// isTestingT checks whether the expression is github.com/stretchr/testify/suite.T.Run
func isSuite(pass *analysis.Pass, expression ast.Expr) bool {
	typeT, ok := pass.TypesInfo.Types[expression]
	if !ok {
		return true
	}

	// TODO : Consider whether this is the best approach
	return strings.HasSuffix(typeT.Type.String(), "Suite")
}
