package generation

import (
	"github.com/k0kubun/pp"
	"go/ast"
	"go/parser"
	"go/token"
)

type XX string

const (
	DEFAULT_OUTPUT_DIR       = "."
	DEFAULT_OUTPUT_FILE_NAME = "./examples/interface.go"
)

type (
	fileParser struct {

		interfaceTypes []*interfaceType
	}
	interfaceType struct {
		name       string
		interfaces subInterface
		functions  []*method
	}

	subInterface struct {
		Name      string
		Functions []*method
	}

	method struct {
		name    string
		params  map[string]string
		returns map[string]string
	}
)

// Parse parse func comment.
func Parse() error {
	var (
		pkgName        = ""
		importStrings  = []string{}
		interfaceNames = []string{}
	)

	var fs = token.NewFileSet()
	// this is a comment

	f, err := parser.ParseFile(fs, DEFAULT_OUTPUT_FILE_NAME, nil, 0)
	if err != nil {
		return err
	}

	pp.Println(f)
	pkgName = f.Name.Name

	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.ImportSpec:
					importStrings = append(importStrings, spec.(*ast.ImportSpec).Path.Value)
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)

					interfaceNames = append(interfaceNames, typeSpec.Name.Name)
					switch typeSpec.Type.(type) {
					case *ast.InterfaceType:
						interfaceType := typeSpec.Type.(*ast.InterfaceType)
						for _, field := range interfaceType.Methods.List {
							//pp.Println(field)
						}
					default:
						continue
					}

				default:
					continue
				}
			}
		}
	}

	println(pkgName)
	return nil
}
