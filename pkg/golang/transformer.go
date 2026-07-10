package golang

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	mathrand "math/rand"
	"strconv"
	"strings"

	"gobfuscator/options"
	"gobfuscator/pkg/utils"
)

// Transformer handles the AST modifications based on options.
type Transformer struct {
	opts      options.Options
	nameMap   map[string]string
	generated map[ast.Node]bool
}

// NewTransformer creates a new Transformer.
func NewTransformer(opts options.Options) *Transformer {
	return &Transformer{
		opts:      opts,
		nameMap:   make(map[string]string),
		generated: make(map[ast.Node]bool),
	}
}

// Transform applies obfuscation rules to the given file AST.
func (t *Transformer) Transform(f *ast.File) {
	// First pass: collect declarations to rename
	if t.opts.RenameVariables {
		t.collectDeclarations(f)
	}

	// Second pass: mutate AST
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return true
		}

		if t.opts.ObfuscateNumbers {
			t.obfuscateNumbers(n)
		}

		if t.opts.RenameVariables {
			t.renameIdentifiers(n)
		}

		return true
	})
	
	if t.opts.ObfuscateStrings {
	    t.obfuscateStringsPass(f)
	}

	// Third pass: inject dead code
	if t.opts.AddDeadCode {
		t.injectDeadCodePass(f)
	}

	// Fourth pass: flatten flow
	if t.opts.FlattenFlow {
		t.flattenFlowPass(f)
	}

	// Fifth pass: hide API
	if t.opts.HideAPI {
		t.hideAPIPass(f)
	}

	// Sixth pass: Semantic Junk
	if t.opts.SemanticJunk {
		t.semanticJunkPass(f)
	}

	// Seventh pass: Anti-Debug
	if t.opts.AntiDebug {
		t.antiDebugPass(f)
	}

	// Eighth pass: Opaque Predicates
	if t.opts.OpaquePreds {
		t.opaquePredsPass(f)
	}

	// Ninth pass: Mixed Boolean-Arithmetic (MBA)
	if t.opts.MBA {
		t.mbaPass(f)
	}
}

func (t *Transformer) collectDeclarations(f *ast.File) {
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ValueSpec:
			for _, name := range x.Names {
				if name.Name != "_" {
					t.nameMap[name.Name] = t.getOrGenerateName(name.Name)
				}
			}
		case *ast.AssignStmt:
			if x.Tok == token.DEFINE {
				for _, expr := range x.Lhs {
					if ident, ok := expr.(*ast.Ident); ok && ident.Name != "_" {
						if _, exists := t.nameMap[ident.Name]; !exists {
							t.nameMap[ident.Name] = t.getOrGenerateName(ident.Name)
						}
					}
				}
			}
		case *ast.FuncDecl:
			if x.Name.Name != "main" && x.Name.Name != "init" && !ast.IsExported(x.Name.Name) {
				t.nameMap[x.Name.Name] = t.getOrGenerateName(x.Name.Name)
			}
			if x.Type.Params != nil {
				for _, field := range x.Type.Params.List {
					for _, name := range field.Names {
						if name.Name != "_" {
							t.nameMap[name.Name] = t.getOrGenerateName(name.Name)
						}
					}
				}
			}
		}
		return true
	})
}

func (t *Transformer) getOrGenerateName(original string) string {
	if newName, ok := t.opts.GlobalNameMap[original]; ok {
		return newName
	}
	newName := generateRandomName()
	t.opts.GlobalNameMap[original] = newName
	return newName
}

func (t *Transformer) renameIdentifiers(n ast.Node) {
	if ident, ok := n.(*ast.Ident); ok {
		if newName, exists := t.nameMap[ident.Name]; exists {
			ident.Name = newName
		}
	}
}

func (t *Transformer) obfuscateStringsPass(f *ast.File) {
	if t.generated == nil {
		t.generated = make(map[ast.Node]bool)
	}
	
	decFuncName := generateRandomName()
	
	// Inject decryptor function
	src := fmt.Sprintf(`package main
	func %s(d []byte, k byte) string {
		for i := 0; i < len(d); i++ { d[i] ^= k }
		return string(d)
	}`, decFuncName)
	f2, err := parser.ParseFile(token.NewFileSet(), "", src, 0)
	if err == nil {
		f.Decls = append(f.Decls, f2.Decls[0])
	}
	t.opts.GlobalNameMap["_currentDecFunc"] = decFuncName
	
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil { return true }
		if t.generated[n] { return false }
		
		switch x := n.(type) {
		case *ast.AssignStmt:
			for i, rhs := range x.Rhs {
				x.Rhs[i] = t.tryObfuscateString(rhs)
			}
		case *ast.BinaryExpr:
			x.X = t.tryObfuscateString(x.X)
			x.Y = t.tryObfuscateString(x.Y)
		case *ast.CallExpr:
			for i, arg := range x.Args {
				x.Args[i] = t.tryObfuscateString(arg)
			}
		case *ast.CompositeLit:
			for i, elt := range x.Elts {
				x.Elts[i] = t.tryObfuscateString(elt)
			}
		case *ast.KeyValueExpr:
			x.Value = t.tryObfuscateString(x.Value)
			x.Key = t.tryObfuscateString(x.Key)
		case *ast.ReturnStmt:
			for i, res := range x.Results {
				x.Results[i] = t.tryObfuscateString(res)
			}
		case *ast.IndexExpr:
			x.Index = t.tryObfuscateString(x.Index)
		case *ast.SendStmt:
			x.Value = t.tryObfuscateString(x.Value)
		}
		return true
	})
}

func (t *Transformer) tryObfuscateString(expr ast.Expr) ast.Expr {
	if lit, ok := expr.(*ast.BasicLit); ok && lit.Kind == token.STRING {
		val, err := strconv.Unquote(lit.Value)
		if err == nil && val != "" {
			if t.opts.SplitStrings && len(val) > 3 {
				return t.createSplitDecryptCall(val)
			}
			return t.createDecryptCall(val)
		}
	}
	return expr
}

func (t *Transformer) createSplitDecryptCall(val string) ast.Expr {
	var chunks []string
	idx := 0
	for idx < len(val) {
		chunkSize := mathrand.Intn(len(val)/2+1) + 1
		if idx+chunkSize > len(val) {
			chunkSize = len(val) - idx
		}
		chunks = append(chunks, val[idx:idx+chunkSize])
		idx += chunkSize
	}
	
	var current ast.Expr = t.createDecryptCall(chunks[0])
	for i := 1; i < len(chunks); i++ {
		current = &ast.BinaryExpr{
			X:  current,
			Op: token.ADD,
			Y:  t.createDecryptCall(chunks[i]),
		}
	}
	return current
}

func (t *Transformer) createDecryptCall(val string) ast.Expr {
	key := byte(42) // Simple fixed key for demonstration
	var elements []ast.Expr
	
	for i := 0; i < len(val); i++ {
		encByte := val[i] ^ key
		elements = append(elements, &ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%d", encByte)})
	}
	
	byteSlice := &ast.CompositeLit{
		Type: &ast.ArrayType{Elt: &ast.Ident{Name: "byte"}},
		Elts: elements,
	}
	
	decFuncName := t.opts.GlobalNameMap["_currentDecFunc"]
	if decFuncName == "" {
		decFuncName = "_0xDec"
	}
	
	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: decFuncName},
		Args: []ast.Expr{
			byteSlice,
			&ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%d", key)},
		},
	}
	
	t.generated[call] = true
	t.generated[byteSlice] = true
	return call
}

func (t *Transformer) obfuscateNumbers(n ast.Node) {
	if t.generated == nil {
		t.generated = make(map[ast.Node]bool)
	}
	if t.generated[n] {
		return
	}

	switch x := n.(type) {
	case *ast.AssignStmt:
		for i, expr := range x.Rhs {
			if lit, ok := expr.(*ast.BasicLit); ok && lit.Kind == token.INT {
				x.Rhs[i] = t.createMathExpr(lit.Value)
			}
		}
	case *ast.BinaryExpr:
		if lit, ok := x.X.(*ast.BasicLit); ok && lit.Kind == token.INT {
			x.X = t.createMathExpr(lit.Value)
		}
		if lit, ok := x.Y.(*ast.BasicLit); ok && lit.Kind == token.INT {
			x.Y = t.createMathExpr(lit.Value)
		}
	case *ast.CallExpr:
		for i, arg := range x.Args {
			if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.INT {
				x.Args[i] = t.createMathExpr(lit.Value)
			}
		}
	}
}

func (t *Transformer) createMathExpr(original string) ast.Expr {
	val, err := strconv.Atoi(original)
	if err != nil {
		return &ast.BasicLit{Kind: token.INT, Value: original}
	}

	offset := mathrand.Intn(500) + 10
	var binExpr *ast.BinaryExpr

	choice := mathrand.Intn(3)
	if choice == 0 {
		// val = part1 ^ offset
		part1 := val ^ offset
		binExpr = &ast.BinaryExpr{
			X:  &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(part1)},
			Op: token.XOR,
			Y:  &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(offset)},
		}
	} else if choice == 1 {
		// val = part1 - offset
		part1 := val + offset
		binExpr = &ast.BinaryExpr{
			X:  &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(part1)},
			Op: token.SUB,
			Y:  &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(offset)},
		}
	} else {
		// val = part1 / offset (assuming exact multiple)
		part1 := val * offset
		binExpr = &ast.BinaryExpr{
			X:  &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(part1)},
			Op: token.QUO,
			Y:  &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(offset)},
		}
	}
	
	parenExpr := &ast.ParenExpr{X: binExpr}

	if t.generated == nil {
		t.generated = make(map[ast.Node]bool)
	}

	var markGenerated func(node ast.Node)
	markGenerated = func(node ast.Node) {
		if node == nil {
			return
		}
		t.generated[node] = true
		ast.Inspect(node, func(n ast.Node) bool {
			if n != nil {
				t.generated[n] = true
			}
			return true
		})
	}

	markGenerated(parenExpr)

	return parenExpr
}

func (t *Transformer) injectDeadCodePass(f *ast.File) {
	var targetBlocks []*ast.BlockStmt

	ast.Inspect(f, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Body != nil {
			targetBlocks = append(targetBlocks, fn.Body)
		}
		return true
	})

	for _, block := range targetBlocks {
		if len(block.List) == 0 {
			continue
		}

		predStr := utils.GenerateOpaquePredicate("go", nil)
		expr, err := parser.ParseExpr(predStr)
		if err != nil {
			expr = &ast.Ident{Name: "false"}
		}

		r1, r2 := mathrand.Intn(1000), mathrand.Intn(100)
		garbageBody := []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.Ident{Name: "_"}},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{
					&ast.BinaryExpr{
						X: &ast.BinaryExpr{
							X:  &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(r1)},
							Op: token.XOR,
							Y:  &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(r2)},
						},
						Op: token.SHL,
						Y:  &ast.BasicLit{Kind: token.INT, Value: "2"},
					},
				},
			},
		}

		var deadCodeBlock ast.Stmt
		if mathrand.Intn(2) == 0 {
			deadCodeBlock = &ast.IfStmt{
				Cond: expr,
				Body: &ast.BlockStmt{List: garbageBody},
			}
		} else {
			deadCodeBlock = &ast.ForStmt{
				Cond: expr,
				Body: &ast.BlockStmt{List: garbageBody},
			}
		}

		block.List = append([]ast.Stmt{deadCodeBlock}, block.List...)
	}
}

func (t *Transformer) flattenFlowPass(f *ast.File) {
	var targetBlocks []*ast.BlockStmt

	ast.Inspect(f, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Body != nil {
			targetBlocks = append(targetBlocks, fn.Body)
		}
		return true
	})

	for _, block := range targetBlocks {
		if len(block.List) == 0 {
			continue
		}
		
		stateVar := generateRandomName()
		
		initStmt := &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: stateVar}},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "1"}},
		}
		
		endStmt := &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: stateVar}},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "0"}},
		}
		
		caseBody := append(block.List, endStmt)
		
		caseClause := &ast.CaseClause{
			List: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "1"}},
			Body: caseBody,
		}
		
		switchStmt := &ast.SwitchStmt{
			Tag: &ast.Ident{Name: stateVar},
			Body: &ast.BlockStmt{List: []ast.Stmt{caseClause}},
		}
		
		forStmt := &ast.ForStmt{
			Cond: &ast.BinaryExpr{
				X:  &ast.Ident{Name: stateVar},
				Op: token.GTR,
				Y:  &ast.BasicLit{Kind: token.INT, Value: "0"},
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{switchStmt}},
		}
		
		hasReturn := false
		// Find if the block belongs to a function that has return values
		ast.Inspect(f, func(n ast.Node) bool {
			if fn, ok := n.(*ast.FuncDecl); ok && fn.Body == block {
				if fn.Type.Results != nil && len(fn.Type.Results.List) > 0 {
					hasReturn = true
				}
			}
			return true
		})
		
		block.List = []ast.Stmt{initStmt, forStmt}
		
		if hasReturn {
			panicStmt := &ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.Ident{Name: "panic"},
					Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"unreachable"`}},
				},
			}
			block.List = append(block.List, panicStmt)
		}
	}
}

func (t *Transformer) hideAPIPass(f *ast.File) {
	apiMap := make(map[string]string)
	
	ast.Inspect(f, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
				if id, ok := sel.X.(*ast.Ident); ok {
					if id.Name == "fmt" || id.Name == "os" || id.Name == "strings" || id.Name == "math" || id.Name == "bytes" || id.Name == "regexp" {
						// Only alias exported functions (start with uppercase)
						if len(sel.Sel.Name) > 0 && sel.Sel.Name[0] >= 'A' && sel.Sel.Name[0] <= 'Z' {
							fullName := id.Name + "." + sel.Sel.Name
							alias, exists := apiMap[fullName]
							if !exists {
								alias = generateRandomName()
								apiMap[fullName] = alias
							}
							call.Fun = &ast.Ident{Name: alias}
						}
					}
				}
			}
		}
		return true
	})
	
	if len(apiMap) > 0 {
		var specs []ast.Spec
		for fullName, alias := range apiMap {
			parts := strings.Split(fullName, ".")
			pkgName := parts[0]
			funcName := parts[1]
			
			spec := &ast.ValueSpec{
				Names: []*ast.Ident{{Name: alias}},
				Values: []ast.Expr{
					&ast.SelectorExpr{
						X:   &ast.Ident{Name: pkgName},
						Sel: &ast.Ident{Name: funcName},
					},
				},
			}
			specs = append(specs, spec)
		}
		decl := &ast.GenDecl{
			Tok:   token.VAR,
			Specs: specs,
		}
		importIdx := 0
		for i, d := range f.Decls {
			if gen, ok := d.(*ast.GenDecl); ok && gen.Tok == token.IMPORT {
				importIdx = i + 1
			} else {
				break
			}
		}
		newDecls := make([]ast.Decl, 0, len(f.Decls)+1)
		newDecls = append(newDecls, f.Decls[:importIdx]...)
		newDecls = append(newDecls, decl)
		newDecls = append(newDecls, f.Decls[importIdx:]...)
		f.Decls = newDecls
	}
}

func (t *Transformer) semanticJunkPass(f *ast.File) {
	numJunk := mathrand.Intn(3) + 2
	var junkDecls []ast.Decl
	
	for i := 0; i < numJunk; i++ {
		structName := generateRandomName()
		methodName := generateRandomName()
		
		// type _0xStructName struct { _0xField1 int; _0xField2 string }
		field1 := generateRandomName()
		field2 := generateRandomName()
		
		structDecl := &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: &ast.Ident{Name: structName},
					Type: &ast.StructType{
						Fields: &ast.FieldList{
							List: []*ast.Field{
								{
									Names: []*ast.Ident{{Name: field1}},
									Type:  &ast.Ident{Name: "int"},
								},
								{
									Names: []*ast.Ident{{Name: field2}},
									Type:  &ast.Ident{Name: "string"},
								},
							},
						},
					},
				},
			},
		}
		
		// func (s *_0xStructName) _0xMethodName() int { return s._0xField1 + 42 }
		methodDecl := &ast.FuncDecl{
			Recv: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: "s"}},
						Type:  &ast.StarExpr{X: &ast.Ident{Name: structName}},
					},
				},
			},
			Name: &ast.Ident{Name: methodName},
			Type: &ast.FuncType{
				Params: &ast.FieldList{},
				Results: &ast.FieldList{
					List: []*ast.Field{
						{Type: &ast.Ident{Name: "int"}},
					},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{
							&ast.BinaryExpr{
								X: &ast.SelectorExpr{
									X:   &ast.Ident{Name: "s"},
									Sel: &ast.Ident{Name: field1},
								},
								Op: token.ADD,
								Y:  &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(mathrand.Intn(1000))},
							},
						},
					},
				},
			},
		}
		
		junkDecls = append(junkDecls, structDecl, methodDecl)
	}
	
	f.Decls = append(f.Decls, junkDecls...)
}

func (t *Transformer) antiDebugPass(f *ast.File) {
	// Look for the main function to inject timing check
	var mainFunc *ast.FuncDecl
	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.Name == "main" {
			mainFunc = fn
			break
		}
	}
	
	if mainFunc != nil {
		// start := time.Now()
		// defer func() { if time.Since(start) > 5 * time.Second { os.Exit(1) } }()
		
		startName := generateRandomName()
		
		startAssign := &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: startName}},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "time"},
						Sel: &ast.Ident{Name: "Now"},
					},
				},
			},
		}
		
		deferStmt := &ast.DeferStmt{
			Call: &ast.CallExpr{
				Fun: &ast.FuncLit{
					Type: &ast.FuncType{Params: &ast.FieldList{}},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.IfStmt{
								Cond: &ast.BinaryExpr{
									X: &ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X:   &ast.Ident{Name: "time"},
											Sel: &ast.Ident{Name: "Since"},
										},
										Args: []ast.Expr{&ast.Ident{Name: startName}},
									},
									Op: token.GTR,
									Y: &ast.BinaryExpr{
										X:  &ast.BasicLit{Kind: token.INT, Value: "5"},
										Op: token.MUL,
										Y: &ast.SelectorExpr{
											X:   &ast.Ident{Name: "time"},
											Sel: &ast.Ident{Name: "Second"},
										},
									},
								},
								Body: &ast.BlockStmt{
									List: []ast.Stmt{
										&ast.ExprStmt{
											X: &ast.CallExpr{
												Fun: &ast.SelectorExpr{
													X:   &ast.Ident{Name: "os"},
													Sel: &ast.Ident{Name: "Exit"},
												},
												Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "1"}},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		
		newBody := []ast.Stmt{startAssign, deferStmt}
		newBody = append(newBody, mainFunc.Body.List...)
		mainFunc.Body.List = newBody
		var importDecl *ast.GenDecl
		for _, decl := range f.Decls {
			if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.IMPORT {
				importDecl = gen
				break
			}
		}

		if importDecl != nil {
			importDecl.Specs = append(importDecl.Specs, 
				&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: `"time"`}},
				&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: `"os"`}},
			)
		} else {
			importDecl = &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: `"time"`}},
					&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: `"os"`}},
				},
			}
			f.Decls = append([]ast.Decl{importDecl}, f.Decls...)
		}
	}
}

func (t *Transformer) opaquePredsPass(f *ast.File) {
	var blocks []*ast.BlockStmt
	ast.Inspect(f, func(n ast.Node) bool {
		if block, ok := n.(*ast.BlockStmt); ok {
			blocks = append(blocks, block)
		}
		return true
	})

	for _, block := range blocks {
		if len(block.List) == 0 {
			continue
		}

		// Prevent modifying Switch or Select bodies which must contain only Case/Comm clauses
		isSwitchBody := false
		for _, stmt := range block.List {
			if _, ok := stmt.(*ast.CaseClause); ok {
				isSwitchBody = true
				break
			}
			if _, ok := stmt.(*ast.CommClause); ok {
				isSwitchBody = true
				break
			}
		}
		if isSwitchBody {
			continue
		}
		
		// Only wrap ~20% of blocks to avoid insane code bloat
		if mathrand.Float32() > 0.2 {
			continue
		}

		// Random integer for the opaque predicate: (x * x + x) % 2 == 0
		randVal := mathrand.Intn(100) + 1
		randValStr := strconv.Itoa(randVal)

		cond := &ast.BinaryExpr{
			X: &ast.BinaryExpr{
				X: &ast.BinaryExpr{
					X: &ast.BinaryExpr{
						X:  &ast.BasicLit{Kind: token.INT, Value: randValStr},
						Op: token.MUL,
						Y:  &ast.BasicLit{Kind: token.INT, Value: randValStr},
					},
					Op: token.ADD,
					Y:  &ast.BasicLit{Kind: token.INT, Value: randValStr},
				},
				Op: token.REM,
				Y:  &ast.BasicLit{Kind: token.INT, Value: "2"},
			},
			Op: token.EQL,
			Y:  &ast.BasicLit{Kind: token.INT, Value: "0"},
		}

		trueBranch := &ast.BlockStmt{List: block.List}
		falseBranch := &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.Ident{Name: "panic"},
						Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"unreachable"`}},
					},
				},
			},
		}

		ifStmt := &ast.IfStmt{
			Cond: cond,
			Body: trueBranch,
			Else: falseBranch,
		}

		block.List = []ast.Stmt{ifStmt}
	}
}

func (t *Transformer) mbaPass(f *ast.File) {
	ast.Inspect(f, func(n ast.Node) bool {
		if bin, ok := n.(*ast.BinaryExpr); ok {
			// We only safely obfuscate SUB (-) to avoid breaking string concatenation (+)
			// x - y => (x ^ ^y) + 2 * (x & ^y) + 1
			if bin.Op == token.SUB {
				x := bin.X
				y := bin.Y

				notY := &ast.UnaryExpr{
					Op: token.XOR, // Bitwise NOT in Go is ^
					X:  y,
				}

				xXorNotY := &ast.BinaryExpr{
					X:  x,
					Op: token.XOR,
					Y:  notY,
				}

				xAndNotY := &ast.BinaryExpr{
					X:  x,
					Op: token.AND,
					Y:  notY,
				}

				twoTimesXAndNotY := &ast.BinaryExpr{
					X:  &ast.BasicLit{Kind: token.INT, Value: "2"},
					Op: token.MUL,
					Y:  &ast.ParenExpr{X: xAndNotY}, // (x & ^y)
				}

				// (x ^ ^y) + 2 * (x & ^y)
				part1 := &ast.BinaryExpr{
					X:  &ast.ParenExpr{X: xXorNotY},
					Op: token.ADD,
					Y:  twoTimesXAndNotY,
				}

				// part1 + 1
				finalExpr := &ast.BinaryExpr{
					X:  part1,
					Op: token.ADD,
					Y:  &ast.BasicLit{Kind: token.INT, Value: "1"},
				}

				bin.X = finalExpr
				bin.Op = token.ADD // Change operation to add 0 just to satisfy the structure if needed? No, we can just change bin.X to finalExpr, bin.Op to ADD, bin.Y to 0
				bin.Y = &ast.BasicLit{Kind: token.INT, Value: "0"}
			}
		}
		return true
	})
}

func generateRandomName() string {
	b := make([]byte, 4)
	rand.Read(b)
	return "_0x" + hex.EncodeToString(b)
}
