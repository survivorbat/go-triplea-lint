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
	testFuncs := findTestFunctions(pass)
	testFuncs = expandSubtests(pass, testFuncs)

	for testFile, testFunctions := range testFuncs {
		for _, testFunc := range testFunctions {
			// Empty test, not our problem
			if len(testFunc.Body.List) == 0 {
				continue
			}

			nameComponents := testNameComponents(testFunc.Name.Name)

			actIndex, ok := findAct(testFunc.Body.List, nameComponents)
			if !ok {
				continue
			}

			actStatement := testFunc.Body.List[actIndex]

			startPos := testFunc.Pos()
			endPos := testFunc.End()

			// NoticE: This does include the // Act itself
			commentsBeforeAct := findCommentsBetween(testFile.Comments, startPos, actStatement.Pos())
			commentsAfterAct := findCommentsBetween(testFile.Comments, actStatement.Pos(), endPos)

			// If the actIndex is 0, it means it is the first and only call in the function. If it is 1 and the
			// call at 0 is t.Parallel OR t := suite.T(), it also means that no arrange is necessary. If both conditions are
			// false, then an arrange is required
			if isArrangeRequired(pass, actIndex, testFunc.Body.List) {
				if len(commentsBeforeAct) < 2 {
					pass.Report(analysis.Diagnostic{
						Pos:     testFunc.Body.List[0].Pos(),
						End:     testFunc.Body.List[0].End(),
						Message: "// Arrange statement expected",
					})
				} else if !strings.Contains(commentsBeforeAct[0].Text(), "Arrange") {
					pass.Report(analysis.Diagnostic{
						Pos:     testFunc.Body.List[0].Pos(),
						End:     testFunc.Body.List[0].End(),
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
			if len(testFunc.Body.List) > actIndex+1 {
				afterActStatement := testFunc.Body.List[actIndex+1]
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

	return nil, nil
}

// isArrangeRequired determines whether the code should have an arrange statement. This
// statement is only necessary if there is non-setup code above the Act statemet.
func isArrangeRequired(pass *analysis.Pass, actIndex int, statements []ast.Stmt) bool {
	switch {
	case actIndex == 0:
		return false
	case actIndex == 1 && isArrangeException(pass, statements[0]):
		return false
	default:
		return true
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

func expandSubtests(pass *analysis.Pass, testFiles map[*ast.File][]*ast.FuncDecl) map[*ast.File][]*ast.FuncDecl {
	output := make(map[*ast.File][]*ast.FuncDecl, len(testFiles))

	for testFile, testFunctions := range testFiles {
		for _, testFunc := range testFunctions {
			subTests, ok := findSubTests(pass, testFunc.Body.List)
			if !ok {
				output[testFile] = append(output[testFile], testFunc)
				continue
			}

			output[testFile] = append(output[testFile], subTests...)
		}
	}

	return output
}

// findSubTests recursively walks through the given statements hunting for subtests, if any are found they
// are returned. The boolean indicates if any were found.
func findSubTests(pass *analysis.Pass, statements []ast.Stmt) ([]*ast.FuncDecl, bool) {
	var result []*ast.FuncDecl

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

		ast.Print(pass.Fset, selectorExp)

		typeT, ok := pass.TypesInfo.Types[selectorExp]
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
