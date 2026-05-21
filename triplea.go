package linters

import (
	"go/ast"
	"go/token"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "triplea",
		Doc:  "TODO",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	testFiles := findTestFunctions(pass)

	for testFile, testFuncs := range testFiles {
		for _, testFunc := range testFuncs {
			for _, testBlock := range findTestBlocks(pass, testFunc) {
				// Empty test, not our problem
				if len(testBlock.List) == 0 {
					continue
				}

				nameComponents := testNameComponents(testFunc.Name.Name)

				actIndex, ok := findAct(testBlock.List, nameComponents)
				if !ok {
					continue
				}

				actStatement := testBlock.List[actIndex]

				// NoticE: This does include the // Act itself
				commentsBeforeAct := findCommentsBetween(testFile.Comments, testBlock.Pos(), actStatement.Pos())
				commentsAfterAct := findCommentsBetween(testFile.Comments, actStatement.Pos(), testBlock.End())

				// If the actIndex is 0, it means it is the first and only call in the function. If it is 1 and the
				// call at 0 is t.Parallel OR t := suite.T(), it also means that no arrange is necessary. If both conditions are
				// false, then an arrange is required
				arrangeExceptions, ok := isArrangeRequired(pass, actIndex, testBlock.List)
				if ok {
					if len(commentsBeforeAct) < 1 || !strings.Contains(commentsBeforeAct[0].Text(), "Arrange") {
						pass.Report(analysis.Diagnostic{
							Pos:     testBlock.List[arrangeExceptions].Pos(),
							End:     testBlock.List[arrangeExceptions].End(),
							Message: "// Arrange statement expected",
						})
					}
				}

				// Check if any comments are before the act
				if len(commentsBeforeAct) == 0 {
					pass.Report(analysis.Diagnostic{
						Pos:     actStatement.Pos(),
						End:     actStatement.End(),
						Message: "// Act statement expected",
					})
				} else {
					closestComment := commentsBeforeAct[len(commentsBeforeAct)-1]
					if !strings.Contains(closestComment.Text(), "Act") {
						pass.Report(analysis.Diagnostic{
							Pos:     actStatement.Pos(),
							End:     actStatement.End(),
							Message: "// Act statement expected",
						})
					}
				}

				// Check if statements exist after the Act
				if len(testBlock.List) > actIndex+1 {
					afterActStatement := testBlock.List[actIndex+1]
					if len(commentsAfterAct) == 0 {
						pass.Report(analysis.Diagnostic{
							Pos:     afterActStatement.Pos(),
							End:     afterActStatement.End(),
							Message: "// Assert statement expected",
						})
					} else if !strings.Contains(commentsAfterAct[0].Text(), "Assert") {
						pass.Report(analysis.Diagnostic{
							Pos:     afterActStatement.Pos(),
							End:     afterActStatement.End(),
							Message: "// Assert statement expected",
						})
					}
				} else {
					// If not, something is wrong
					pass.Report(analysis.Diagnostic{
						Pos:     actStatement.Pos(),
						End:     actStatement.End(),
						Message: "// Assert statement expected",
					})
				}
			}
		}
	}

	return nil, nil
}

// isArrangeRequired determines whether the code should have an arrange statement. This
// statement is only necessary if there is non-setup code above the Act statement. The first
// return value indicates how many statements are considered exceptions.
//
// Future-proof implementation
func isArrangeRequired(pass *analysis.Pass, actIndex int, statements []ast.Stmt) (int, bool) {
	switch {
	case actIndex == 0:
		return 0, false
	case actIndex == 1 && isArrangeException(pass, statements[0]):
		return 1, false
	default:
		return 0, true
	}
}

// findCommentsBetween filters comments between the 2 given positions, used primarily to split
// comments between the start, act and end of a test
func findCommentsBetween(comments []*ast.CommentGroup, from token.Pos, to token.Pos) []*ast.CommentGroup {
	result := make([]*ast.CommentGroup, 0, len(comments))

	for _, comment := range comments {
		if comment.Pos() > from && comment.Pos() < to {
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

func findAct(statements []ast.Stmt, functionNames []string) (int, bool) {
	actIndex := -1

	for index, statement := range statements {
		ast.Inspect(statement, func(node ast.Node) bool {
			callExp, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			identExp, iOk := callExp.Fun.(*ast.Ident)
			selectorExp, sOk := callExp.Fun.(*ast.SelectorExpr)
			if !iOk && !sOk {
				return true
			}

			if iOk && slices.Contains(functionNames, identExp.Name) {
				actIndex = index
				return false
			} else if sOk && slices.Contains(functionNames, selectorExp.Sel.Name) {
				actIndex = index
				return false
			}

			// TODO: Add more logic here, desperately searching for the Act

			return true
		})
	}

	return actIndex, actIndex != -1
}

// isArrangeException returns whether the given statement is allowed to be outside of the Arrange statement
func isArrangeException(pass *analysis.Pass, statement ast.Stmt) bool {
	return isTParallel(pass, statement) || isSuiteT(pass, statement)
}

// isSuiteT returns whether the given statement equals an assignment such as `t := suite.T()` at the top
// of a suite test.
func isSuiteT(pass *analysis.Pass, statement ast.Stmt) bool {
	return isCall(pass, statement, "T", "github.com/stretchr/testify/suite.T")
}

// isTParallel returns whether the given statement is a call to testing.T.Parallel
func isTParallel(pass *analysis.Pass, statement ast.Stmt) bool {
	return isCall(pass, statement, "Parallel", "*testing.T")
}

func isCall(pass *analysis.Pass, statement ast.Stmt, selector string, typeName string) bool {
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

		isCall = selectorExp.Sel.Name == selector && typeT.Type.String() == typeName

		return false
	})

	return isCall
}

// testNameComponents removes the Test prefix and splits the test name into pieces.
func testNameComponents(name string) []string {
	return strings.Split(strings.TrimPrefix(name, "Test"), "_")
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

	return typeT.Type.String() == "github.com/stretchr/testify/suite.T"
}
