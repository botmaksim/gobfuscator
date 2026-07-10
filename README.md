# Gobfuscator

Gobfuscator is an advanced, multi-language code obfuscator designed to protect intellectual property and deter reverse engineering. It supports multiple languages and applies a variety of code transformations, ranging from simple variable renaming to complex control flow flattening and mathematical rewriting.

## Features and Capabilities

The obfuscator operates through multiple discrete passes, which can be toggled via command-line flags.

- **Variable and Function Renaming (`--rename`)**: Scans the code for all declared identifiers — variables, function definitions, method parameters — and replaces every occurrence with randomized hex-style names (e.g., `_0x3f1a2c7b`). For C++, Java, and JavaScript the engine extracts typed declarations (`int`, `void`, `function`, etc.) as well as function signatures. For Ruby and Python it additionally extracts `def` function definitions. All renames are consistent within a single run: the same original name always maps to the same obfuscated name.

- **Polymorphic String Encryption (`--strings`)**: Encrypts string literals using randomized algorithms (XOR, ADD, SUB, NOT) and injects dynamic decryption logic. For Go, this generates inline XOR decode functions directly into the AST. For Ruby, the entire script is compressed with zlib, base64-encoded, and wrapped in a one-liner `eval(Zlib::Inflate.inflate(Base64.decode64(...)))`. Use `--ruby-no-eval` to skip this final wrapping step.

- **String Splitting (`--split-strings`)**: Breaks large string literals into randomized, smaller chunks that are concatenated at runtime, defeating simple string extraction tools.

- **Integer Literal Obfuscation (`--numbers`)**: Replaces numeric constants with randomized mathematical identities evaluated at runtime (e.g., `value = (value ^ offset) ^ offset`).

- **Control Flow Flattening (`--flatten`)**: Transforms linear execution paths into complex state machines, completely destroying the original control graph. In the unreachable fallthrough branch, an infinite loop with an opaque condition and garbage arithmetic body is injected to further mislead static analysis.

- **Mixed Boolean-Arithmetic (`--mba`)**: Rewrites standard arithmetic operations using computationally equivalent but syntactically convoluted bitwise logic (e.g., `a + b` becomes `((a ^ b) + 2 * (a & b))`).

- **Opaque Predicates (`--opaque-preds`)**: Injects tautological conditional branches after function definitions. The true branch contains the actual code; the false branch is an infinite loop whose condition is itself a second independently generated opaque predicate. The loop body contains garbage arithmetic computations to make it indistinguishable from real loop logic. See the section below for the exact structure.

- **Dead Code Injection (`--deadcode`)**: Generates and inserts structurally valid but functionally meaningless code blocks wrapped in opaque predicates.

- **Semantic Junk (`--semantic-junk`)**: Creates entirely fake structures, classes, and methods tailored to the target programming language to artificially inflate cyclomatic complexity.

- **API Hiding (`--hide-api`)**: Wraps common standard library calls (like print statements) in dynamically generated wrapper functions or classes.

- **Anti-Debugging (`--anti-debug`)**: Injects language-specific countermeasures (timing checks, debugger interrupts, call stack verifications) to hinder dynamic analysis.

## Opaque Infinite Loop Structure

A distinctive technique applied in both opaque predicates and CFF fallthrough paths is a layered infinite loop pattern:

```
// C++ / Java / JavaScript example
if ((N1 * N1 + N1) % 2 != 0) {
    // original code
} else {
    while ((N2 * N2 + N2) % 2 != 0) {
        _0xGarbage = rand1 + rand2;
    }
}

// Go (AST)
if cond1 {
    // original code
} else {
    for {
        _ = rand1 + rand2
    }
}

// Python (AST)
if cond1:
    # original code
else:
    while cond1:
        _0xFake = rand1 + rand2
```

Both predicates are generated independently from the template `((N * N + N) % 2 != 0)` with distinct random constants, ensuring neither is a literal `true`. The loop body discards the result of an arithmetic expression, making the block look like a legitimate computation to a naive reader.

## Supported Languages & Auto-Detection

Gobfuscator features automatic language detection based on file extensions. When parsing a file or traversing a directory, the engine identifies the language and routes it to the appropriate obfuscation module.

Supported extensions:
- **Go**: `.go`
- **C++**: `.cpp`, `.cc`, `.cxx`, `.h`, `.hpp`
- **Java**: `.java`
- **JavaScript**: `.js`, `.cjs`, `.mjs`
- **Python**: `.py`
- **Ruby**: `.rb`

You can bypass auto-detection and force a specific engine with `--lang` (e.g., `--lang=cpp`).

## Installation

Ensure you have Go installed on your system, then build the binary:

```bash
go build -o gobfuscator
```

## Usage

Gobfuscator can process individual files or entire project directories.

### Default Output

The `-o` flag is optional. When omitted:

- **Single file**: the obfuscated output is written to the **current working directory** with the same filename as the input.
- **Directory**: a new directory named `<inputdir>_obf` is created in the **current working directory**, mirroring the full subdirectory structure of the input.

### Basic Usage

Process a single file (output written to `./script.py` in cwd):
```bash
./gobfuscator path/to/script.py
```

Process a single file with an explicit output path:
```bash
./gobfuscator script.py -o obfuscated_script.py
```

Process an entire directory (output goes to `./src_obf/` in cwd):
```bash
./gobfuscator ./src
```

Process an entire directory with an explicit output path:
```bash
./gobfuscator ./src_dir -o ./build_dir
```

### Advanced Usage

Disable anti-debugging and control flow flattening:
```bash
./gobfuscator main.go -o build.go --anti-debug=false --flatten=false
```

Force a specific language engine:
```bash
./gobfuscator code.txt -o obfuscated_code.txt --lang=java
```

Obfuscate a Ruby script but keep output readable (no eval wrapping):
```bash
./gobfuscator script.rb -o obfuscated.rb --ruby-no-eval
```

## Command-Line Flags

| Flag | Description | Default |
| :--- | :--- | :--- |
| `-o`, `--output` | Output file or directory. Defaults to `cwd/<filename>` for files, `cwd/<dir>_obf` for directories | Auto |
| `--lang` | Force specific language engine (go, cpp, java, js, python, ruby) | Auto-detected |
| `--rename` | Enable global variable and function renaming | true |
| `--strings` | Enable polymorphic string encryption | true |
| `--numbers` | Enable integer literal obfuscation | true |
| `--deadcode` | Enable opaque dead code injection | true |
| `--flatten` | Enable control flow flattening state machines | true |
| `--split-strings` | Enable string splitting and dynamic concatenation | true |
| `--hide-api` | Enable standard library API aliasing | true |
| `--semantic-junk` | Enable semantic junk structure generation | true |
| `--anti-debug` | Enable anti-debugging and timing countermeasures | true |
| `--mba` | Enable Mixed Boolean-Arithmetic mathematical rewriting | true |
| `--opaque-preds` | Inject opaque conditional branches with infinite loop fallback | true |
| `--ruby-no-eval` | Disable zlib+base64 encoding and eval wrapping for Ruby output | false |

## Mechanics

Gobfuscator works in multiple passes. When processing a directory, it first performs a global analysis pass to collect identifiers for cross-file consistent renaming. The second pass processes source files through the respective language engines.

Code is processed at the AST level where possible (Go, Python) or via robust regex-based manipulation for languages where AST parsing is prohibitive (C++, Java, Ruby, JavaScript). String and number transformations use a unique polymorphic generation strategy per build, ensuring varying obfuscation signatures for identical input sources.

### Identifier Extraction

The renaming engine uses language-specific patterns to discover all renameable identifiers:

- **Go**: Full AST traversal — captures `ValueSpec`, `:=` short assignments, function declarations, and function parameters.
- **Python**: Regex over `def` function definitions and bare assignment targets.
- **JavaScript**: Matches `var`, `let`, `const`, and `function` declarations.
- **C++ / Java**: Matches typed variable declarations (`int`, `char`, `float`, `double`, `void`, `boolean`, `long`, `auto`, `String`) and function signatures.
- **Ruby**: Matches `def` function definitions and bare assignment targets.

### Ruby String Encoding

When `--strings` is enabled and `--ruby-no-eval` is not set, Ruby output undergoes a two-stage transformation: the obfuscated source is compressed with zlib, base64-encoded, and the final file contains only a single-line eval decoder. This renders the source entirely opaque even to string-search tools.

Passing `--ruby-no-eval` skips this final step, producing readable (though structurally obfuscated) Ruby code. This is useful for debugging or when running in environments that restrict `eval`.
