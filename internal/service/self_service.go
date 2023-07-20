package service

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const rootDir = "./"

type SelfService interface {
	Find(ctx context.Context, substr string) ([]string, error)
}

type selfService struct{}

func NewSelfService() SelfService {
	return &selfService{}
}

func (ss *selfService) Find(ctx context.Context, substr string) ([]string, error) {
	fset := token.NewFileSet()

	identifiers := make([]string, 0)
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			sourceCode, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			astFile, err := parser.ParseFile(fset, "", string(sourceCode), parser.AllErrors)
			if err != nil {
				return err
			}

			ast.Inspect(astFile, func(node ast.Node) bool {
				if ident, ok := node.(*ast.Ident); ok {
					if strings.Contains(ident.Name, substr) {
						identifiers = append(identifiers, ident.Name)
					}
				}
				return true
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return identifiers, nil
}
