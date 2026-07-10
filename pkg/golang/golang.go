package golang

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"

	"gobfuscator/core"
	"gobfuscator/options"
)

func init() {
	core.Register("go", &GoObfuscator{})
}

// GoObfuscator implements the core.Obfuscator interface for Go source code.
type GoObfuscator struct{}

// Obfuscate applies transformations to the provided Go source code.
func (g *GoObfuscator) Obfuscate(source []byte, opts options.Options) ([]byte, error) {
	fset := token.NewFileSet()
	// Parse the source code to an AST
	f, err := parser.ParseFile(fset, "", source, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// Apply transformations
	transformer := NewTransformer(opts)
	transformer.Transform(f)

	// Format back to source code
	var buf bytes.Buffer
	err = printer.Fprint(&buf, fset, f)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
