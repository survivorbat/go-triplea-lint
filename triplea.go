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
	tests := findTestNodes(pass)

	for testFile, testFunctions := range tests {
		for _, testFunc := range testFunctions {
			nameComponents := testNameComponents(testFunc.Name.Name)

			actIndex, ok := findAct(testFunc.Body.List, nameComponents)
			if !ok {
				continue
			}

			actStatement := testFunc.Body.List[actIndex]

			startPos := testFunc.Body.List[0].Pos()
			endPos := testFunc.Body.List[len(testFunc.Body.List)-1].Pos()

			commentsBeforeAct := findCommentsBetween(testFile.Comments, startPos, actStatement.Pos())
			commentsAfterAct := findCommentsBetween(testFile.Comments, actStatement.Pos(), endPos)

			if len(commentsBeforeAct) == 0 {
				pass.Report(analysis.Diagnostic{
					Pos:     actStatement.Pos(),
					End:     actStatement.End(),
					Message: "// Act statement expected",
				})
			}

			if len(testFunc.Body.List) > actIndex+1 {
				afterActStatement := testFunc.Body.List[actIndex+1]
				if len(commentsAfterAct) == 0 {
					pass.Report(analysis.Diagnostic{
						Pos:     afterActStatement.Pos(),
						End:     afterActStatement.End(),
						Message: "// Assert statement expected",
					})
				}
			}
		}
	}

	return nil, nil
}

func findCommentsBetween(comments []*ast.CommentGroup, from token.Pos, to token.Pos) []*ast.CommentGroup {
	result := make([]*ast.CommentGroup, 0, len(comments))

	for _, comment := range comments {
		if comment.Pos() > from && comment.Pos() < to {
			result = append(result, comment)
		}
	}

	return result
}

func findTestNodes(pass *analysis.Pass) map[*ast.File][]*ast.FuncDecl {
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

			return true
		})
	}

	return actIndex, actIndex != -1
}

func testNameComponents(name string) []string {
	return strings.Split(strings.TrimPrefix(name, "Test"), "_")
}
