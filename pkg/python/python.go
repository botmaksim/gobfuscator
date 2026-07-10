package python

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"gobfuscator/core"
	"gobfuscator/options"
)

func init() {
	core.Register("python", &PythonObfuscator{})
	core.Register("py", &PythonObfuscator{})
}

type PythonObfuscator struct{}

const pythonTransformerScript = `
import ast
import sys
import random
import string
import base64

def random_name():
    return "_0x" + "".join(random.choices(string.hexdigits.lower(), k=8))

class CollectNames(ast.NodeVisitor):
    def __init__(self):
        self.names = set()
    def visit_Name(self, node):
        if isinstance(node.ctx, ast.Store):
            self.names.add(node.id)
    def visit_arg(self, node):
        self.names.add(node.arg)
    def visit_FunctionDef(self, node):
        self.names.add(node.name)
        self.generic_visit(node)

class ObfuscatorTransformer(ast.NodeTransformer):
    def __init__(self, opts):
        self.opts = opts
        self.name_map = {}
        self.api_map = {}

    def visit_Constant(self, node):
        if self.opts.get("ObfuscateStrings", False) and isinstance(node.value, str):
            encoded = base64.b64encode(node.value.encode('utf-8'))
            
            if self.opts.get("SplitStrings", False) and len(encoded) > 3:
                chunks = []
                idx = 0
                while idx < len(encoded):
                    chunk_size = random.randint(1, max(2, len(encoded) // 2))
                    chunks.append(encoded[idx:idx+chunk_size])
                    idx += chunk_size
                    
                current_node = ast.Constant(value=chunks[0])
                for chunk in chunks[1:]:
                    current_node = ast.BinOp(left=current_node, op=ast.Add(), right=ast.Constant(value=chunk))
                bytes_node = current_node
            else:
                bytes_node = ast.Constant(value=encoded)
            
            # base64.b64decode(...)
            b64decode = ast.Attribute(value=ast.Name(id='base64', ctx=ast.Load()), attr='b64decode', ctx=ast.Load())
            call_b64 = ast.Call(func=b64decode, args=[bytes_node], keywords=[])
            
            # .decode()
            decode_func = ast.Attribute(value=call_b64, attr='decode', ctx=ast.Load())
            return ast.Call(func=decode_func, args=[], keywords=[])
            
        if self.opts.get("ObfuscateNumbers", False) and isinstance(node.value, int) and not isinstance(node.value, bool):
            val = node.value
            offset = random.randint(10, 500)
            choice = random.randint(0, 2)
            
            if choice == 0:
                part1 = val ^ offset
                return ast.BinOp(left=ast.Constant(value=part1), op=ast.BitXor(), right=ast.Constant(value=offset))
            elif choice == 1:
                part1 = val + offset
                return ast.BinOp(left=ast.Constant(value=part1), op=ast.Sub(), right=ast.Constant(value=offset))
            else:
                part1 = val * offset
                return ast.BinOp(left=ast.Constant(value=part1), op=ast.FloorDiv(), right=ast.Constant(value=offset))

        return node

    def visit_JoinedStr(self, node):
        for val in node.values:
            if isinstance(val, ast.FormattedValue):
                self.visit(val)
        return node

    def visit_Name(self, node):
        if self.opts.get("HideAPI", False) and isinstance(node.ctx, ast.Load) and node.id in self.api_map:
            node.id = self.api_map[node.id]
        elif self.opts.get("RenameVariables", False) and node.id in self.name_map:
            node.id = self.name_map[node.id]
        return node

    def visit_arg(self, node):
        if self.opts.get("RenameVariables", False) and node.arg in self.name_map:
            node.arg = self.name_map[node.arg]
        return node

    def visit_Attribute(self, node):
        node = self.generic_visit(node)
        if self.opts.get("RenameVariables", False) and node.attr in self.name_map:
            node.attr = self.name_map[node.attr]
        return node

    def visit_BinOp(self, node):
        node = self.generic_visit(node)
        if self.opts.get("MBA", False) and isinstance(node.op, ast.Sub):
            not_y = ast.UnaryOp(op=ast.Invert(), operand=node.right)
            x_xor_not_y = ast.BinOp(left=node.left, op=ast.BitXor(), right=not_y)
            x_and_not_y = ast.BinOp(left=node.left, op=ast.BitAnd(), right=not_y)
            two_times = ast.BinOp(left=ast.Constant(value=2), op=ast.Mult(), right=x_and_not_y)
            part1 = ast.BinOp(left=x_xor_not_y, op=ast.Add(), right=two_times)
            final_expr = ast.BinOp(left=part1, op=ast.Add(), right=ast.Constant(value=1))
            return final_expr
        return node

    def visit_Assign(self, node):
        node = self.generic_visit(node)
        if self.opts.get("OpaquePreds", False) and random.random() < 0.2:
            rand_val = random.randint(1, 100)
            cond = ast.Compare(
                left=ast.BinOp(
                    left=ast.BinOp(
                        left=ast.BinOp(
                            left=ast.Constant(value=rand_val),
                            op=ast.Mult(),
                            right=ast.Constant(value=rand_val)
                        ),
                        op=ast.Add(),
                        right=ast.Constant(value=rand_val)
                    ),
                    op=ast.Mod(),
                    right=ast.Constant(value=2)
                ),
                ops=[ast.Eq()],
                comparators=[ast.Constant(value=0)]
            )
            return ast.If(
                test=cond,
                body=[node],
                orelse=[
                    ast.Expr(value=ast.Call(
                        func=ast.Name(id='Exception', ctx=ast.Load()),
                        args=[ast.Constant(value="")],
                        keywords=[]
                    ))
                ]
            )
        return node

    def flatten_block(self, body):
        if not self.opts.get("FlattenFlow", False) or len(body) <= 1:
            return body
            
        state_var = f"_0xState{random.randint(1000,9999)}"
        states = list(range(1, len(body) + 1))
        random.shuffle(states)
        state_map = {i: states[i] for i in range(len(body))}
        
        def get_next_state(i):
            if i + 1 < len(body):
                return state_map[i + 1]
            return 0
            
        if_statements = []
        for i, stmt in enumerate(body):
            curr_state = state_map[i]
            next_state = get_next_state(i)
            
            next_state_stmt = ast.Assign(
                targets=[ast.Name(id=state_var, ctx=ast.Store())],
                value=ast.Constant(value=next_state)
            )
            
            block_body = [stmt, next_state_stmt]
            
            test = ast.Compare(
                left=ast.Name(id=state_var, ctx=ast.Load()),
                ops=[ast.Eq()],
                comparators=[ast.Constant(value=curr_state)]
            )
            
            if i == 0:
                if_node = ast.If(test=test, body=block_body, orelse=[])
                if_statements.append(if_node)
            else:
                new_if = ast.If(test=test, body=block_body, orelse=[])
                current_if = if_statements[0]
                while current_if.orelse:
                    current_if = current_if.orelse[0]
                current_if.orelse.append(new_if)
                
        init_state_stmt = ast.Assign(
            targets=[ast.Name(id=state_var, ctx=ast.Store())],
            value=ast.Constant(value=state_map[0])
        )
        
        while_test = ast.Compare(
            left=ast.Name(id=state_var, ctx=ast.Load()),
            ops=[ast.NotEq()],
            comparators=[ast.Constant(value=0)]
        )
        while_loop = ast.While(test=while_test, body=[if_statements[0]], orelse=[])
        
        return [init_state_stmt, while_loop]

    def visit_FunctionDef(self, node):
        if self.opts.get("RenameVariables", False) and node.name in self.name_map:
            node.name = self.name_map[node.name]
            
        node = self.generic_visit(node)
        
        if self.opts.get("FlattenFlow", False):
            node.body = self.flatten_block(node.body)
        
        if self.opts.get("AddDeadCode", False):
            choice = random.choice(["if", "while", "try"])
            
            r1, r2 = random.randint(1, 1000), random.randint(1, 100)
            garbage_body = [
                ast.Assign(
                    targets=[ast.Name(id=f"_0x{random.randint(1000, 9999)}", ctx=ast.Store())],
                    value=ast.BinOp(
                        left=ast.BinOp(left=ast.Constant(value=r1), op=ast.BitXor(), right=ast.Constant(value=r2)),
                        op=ast.LShift(),
                        right=ast.Constant(value=2)
                    )
                )
            ]
            
            false_test = ast.Compare(
                left=ast.Call(func=ast.Name(id='id', ctx=ast.Load()), args=[ast.Constant(value=0)], keywords=[]),
                ops=[ast.Eq()],
                comparators=[ast.Constant(value=0)]
            )
            
            if choice == "if":
                dead_code = ast.If(test=false_test, body=garbage_body, orelse=[])
            elif choice == "while":
                dead_code = ast.While(test=false_test, body=garbage_body, orelse=[])
            else:
                dead_code = ast.If(test=false_test, body=[ast.Try(
                    body=garbage_body,
                    handlers=[ast.ExceptHandler(type=ast.Name(id='Exception', ctx=ast.Load()), name=None, body=[ast.Pass()])],
                    orelse=[],
                    finalbody=[]
                )], orelse=[])
                
            node.body.insert(0, dead_code)
            
        return node

def main():
    if len(sys.argv) < 4:
        sys.exit(1)
        
    input_file = sys.argv[1]
    opts_str = sys.argv[2] # e.g. "strings=1,numbers=1,deadcode=1"
    name_map_json = sys.argv[3]
    
    import json
    global_name_map = json.loads(name_map_json)
    
    opts = {}
    for pair in opts_str.split(','):
        if '=' in pair:
            k, v = pair.split('=')
            opts[k] = (v == '1')
            
    with open(input_file, 'r', encoding='utf-8') as f:
        source = f.read()
        
    tree = ast.parse(source)
    
    name_map = {}
    if opts.get("RenameVariables", False):
        collector = CollectNames()
        collector.visit(tree)
        # Avoid renaming __main__ and main, keep it simple for tests
        for n in collector.names:
            if not n.startswith('__') and n != 'main':
                if n in global_name_map:
                    name_map[n] = global_name_map[n]
                else:
                    name_map[n] = random_name()
                
    transformer = ObfuscatorTransformer(opts)
    transformer.name_map = name_map
    
    if opts.get("HideAPI", False):
        builtins_to_hide = ['print', 'len', 'int', 'str', 'list', 'dict', 'set', 'range', 'map', 'filter', 'zip', 'sum', 'min', 'max', 'abs', 'round', 'enumerate', 'id', 'type', 'isinstance']
        for b in builtins_to_hide:
            transformer.api_map[b] = random_name()
            
    new_tree = transformer.visit(tree)
    
    if opts.get("AntiDebug", False):
        import_time = ast.Import(names=[ast.alias(name='time', asname=None)])
        import_sys = ast.Import(names=[ast.alias(name='sys', asname=None)])
        
        # if getattr(sys, 'gettrace')() is not None: sys.exit(1)
        check_trace_if = ast.If(
            test=ast.Compare(
                left=ast.Call(
                    func=ast.Call(
                        func=ast.Name(id='getattr', ctx=ast.Load()),
                        args=[ast.Name(id='sys', ctx=ast.Load()), ast.Constant(value='gettrace')],
                        keywords=[]
                    ),
                    args=[],
                    keywords=[]
                ),
                ops=[ast.IsNot()],
                comparators=[ast.Constant(value=None)]
            ),
            body=[ast.Expr(value=ast.Call(
                func=ast.Attribute(value=ast.Name(id='sys', ctx=ast.Load()), attr='exit', ctx=ast.Load()),
                args=[ast.Constant(value=1)],
                keywords=[]
            ))],
            orelse=[]
        )
        
        new_tree.body = [import_time, import_sys, check_trace_if] + new_tree.body

    if opts.get("HideAPI", False):
        api_stmts = []
        import_builtins = ast.Import(names=[ast.alias(name='builtins', asname=None)])
        api_stmts.append(import_builtins)
        for builtin_name, alias in transformer.api_map.items():
            getattr_call = ast.Call(
                func=ast.Name(id='getattr', ctx=ast.Load()),
                args=[ast.Name(id='builtins', ctx=ast.Load()), ast.Constant(value=builtin_name)],
                keywords=[]
            )
            assign = ast.Assign(targets=[ast.Name(id=alias, ctx=ast.Store())], value=getattr_call)
            api_stmts.append(assign)
        new_tree.body = api_stmts + new_tree.body
        
    if opts.get("SemanticJunk", False):
        junk_class_name = random_name()
        
        methods = []
        for i in range(random.randint(2, 5)):
            junk_method_name = random_name()
            # Create a method that does some random math operations
            math_ops = [ast.Add(), ast.Sub(), ast.Mult(), ast.BitXor(), ast.LShift()]
            
            body = []
            body.append(ast.Assign(targets=[ast.Name(id='_val1', ctx=ast.Store())], value=ast.Constant(value=random.randint(10, 1000))))
            body.append(ast.Assign(targets=[ast.Name(id='_val2', ctx=ast.Store())], value=ast.Constant(value=random.randint(10, 1000))))
            body.append(ast.Assign(targets=[ast.Name(id='_res', ctx=ast.Store())], value=ast.BinOp(left=ast.Name(id='_val1', ctx=ast.Load()), op=random.choice(math_ops), right=ast.Name(id='_val2', ctx=ast.Load()))))
            body.append(ast.Return(value=ast.Name(id='_res', ctx=ast.Load())))

            junk_method = ast.FunctionDef(
                name=junk_method_name,
                args=ast.arguments(
                    posonlyargs=[],
                    args=[ast.arg(arg='self')],
                    kwonlyargs=[],
                    kw_defaults=[],
                    defaults=[]
                ),
                body=body,
                decorator_list=[],
                returns=None
            )
            methods.append(junk_method)
        
        junk_class = ast.ClassDef(
            name=junk_class_name,
            bases=[],
            keywords=[],
            body=methods,
            decorator_list=[]
        )
        new_tree.body.append(junk_class)

    if opts.get("ObfuscateStrings", False):
        # Inject 'import base64' at the top
        import_b64 = ast.Import(names=[ast.alias(name='base64', asname=None)])
        new_tree.body.insert(0, import_b64)
        
    ast.fix_missing_locations(new_tree)
    print(ast.unparse(new_tree))

if __name__ == '__main__':
    main()
`

func (p *PythonObfuscator) Obfuscate(source []byte, opts options.Options) ([]byte, error) {
	// Create a temporary directory for our scripts
	tmpDir, err := ioutil.TempDir("", "gobfuscator_python_")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	scriptPath := filepath.Join(tmpDir, "transformer.py")
	inputPath := filepath.Join(tmpDir, "input.py")

	err = ioutil.WriteFile(scriptPath, []byte(pythonTransformerScript), 0644)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(inputPath, source, 0644)
	if err != nil {
		return nil, err
	}

	optsStr := fmt.Sprintf("ObfuscateStrings=%d,ObfuscateNumbers=%d,AddDeadCode=%d,RenameVariables=%d,FlattenFlow=%d,SplitStrings=%d,HideAPI=%d,SemanticJunk=%d,AntiDebug=%d,MBA=%d,OpaquePreds=%d", 
		boolToInt(opts.ObfuscateStrings), boolToInt(opts.ObfuscateNumbers), boolToInt(opts.AddDeadCode), boolToInt(opts.RenameVariables), boolToInt(opts.FlattenFlow), boolToInt(opts.SplitStrings), boolToInt(opts.HideAPI), boolToInt(opts.SemanticJunk), boolToInt(opts.AntiDebug), boolToInt(opts.MBA), boolToInt(opts.OpaquePreds))

	nameMapBytes, _ := json.Marshal(opts.GlobalNameMap)

	cmd := exec.Command("python3", scriptPath, inputPath, optsStr, string(nameMapBytes))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("python script failed: %v\nstderr: %s", err, stderr.String())
	}

	return out.Bytes(), nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
