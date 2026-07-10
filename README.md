# Gobfuscator

Gobfuscator is an advanced, multi-language code obfuscator designed to protect intellectual property and deter reverse engineering. It supports multiple languages and applies a variety of code transformations, ranging from simple variable renaming to complex control flow flattening and mathematical rewriting.

## Features and Capabilities

The obfuscator operates through multiple discrete passes, which can be toggled via command-line flags.

*   **Variable Renaming (`--rename`)**: Globally renames variables, functions, and classes across single files or entire directories to meaningless, randomized identifiers.
*   **Polymorphic String Encryption (`--strings`)**: Encrypts string literals using randomized algorithms (XOR, ADD, SUB, NOT) and injects dynamic decryption logic.
*   **String Splitting (`--split-strings`)**: Breaks large string literals into randomized, smaller chunks that are concatenated at runtime, defeating simple string extraction tools.
*   **Integer Literal Obfuscation (`--numbers`)**: Replaces numeric constants with randomized mathematical identities evaluated at runtime (e.g., `value = (value ^ offset) ^ offset`).
*   **Control Flow Flattening (`--flatten`)**: Transforms linear execution paths into complex state machines, completely destroying the original control graph.
*   **Mixed Boolean-Arithmetic (`--mba`)**: Rewrites standard arithmetic operations using computationally equivalent but syntactically convoluted bitwise logic (e.g., converting `a + b` to `((a ^ b) + 2 * (a & b))`).
*   **Opaque Predicates (`--opaque-preds`)**: Injects tautological conditional statements that always evaluate to true, padding the code with unreachable, misleading execution branches.
*   **Dead Code Injection (`--deadcode`)**: Generates and inserts structurally valid but functionally meaningless code blocks wrapped in opaque predicates.
*   **Semantic Junk (`--semantic-junk`)**: Creates entirely fake structures, classes, and methods tailored to the target programming language to artificially inflate cyclomatic complexity.
*   **API Hiding (`--hide-api`)**: Wraps common standard library calls (like print statements) in dynamically generated wrapper functions or classes.
*   **Anti-Debugging (`--anti-debug`)**: Injects language-specific countermeasures (timing checks, debugger interrupts, call stack verifications) to hinder dynamic analysis.

## Supported Languages & Auto-Detection

Gobfuscator features automatic language detection based on file extensions. When parsing a file or traversing a directory, the engine identifies the language and routes it to the appropriate obfuscation module.

Supported extensions:
*   **Go**: `.go`
*   **C++**: `.cpp`, `.cc`, `.cxx`, `.h`, `.hpp`
*   **Java**: `.java`
*   **JavaScript**: `.js`, `.cjs`, `.mjs`
*   **Python**: `.py`
*   **Ruby**: `.rb`

You can bypass auto-detection and force the usage of a specific engine using the `--lang` flag (e.g., `--lang=cpp`).

## Installation

Ensure you have Go installed on your system, then build the binary:

```bash
go build -o gobfuscator
```

## Usage

Gobfuscator can process individual files or entire project directories.

### Basic Usage

Process a single file with auto-detected language settings:
```bash
./gobfuscator script.py -o obfuscated_script.py
```

Process an entire directory. The output directory structure will mirror the input directory:
```bash
./gobfuscator ./src_dir -o ./build_dir
```

### Advanced Usage

By default, all obfuscation passes are enabled. You can disable specific passes to tune performance or debug execution.

Disable anti-debugging and control flow flattening:
```bash
./gobfuscator main.go -o build.go --anti-debug=false --flatten=false
```

Force a specific language engine if auto-detection fails or is undesired:
```bash
./gobfuscator code.txt -o obfuscated_code.txt --lang=java
```

## Command-Line Flags

| Flag | Description | Default |
| :--- | :--- | :--- |
| `-o`, `--output` | Destination file or directory for obfuscated code | Required |
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
| `--opaque-preds` | Enable complex logical opaque predicates | true |

## Mechanics

Gobfuscator works in multiple passes. When processing a directory, it first performs a global analysis pass to collect identifiers for cross-file consistent renaming. The second pass processes the source files through the respective language engines.

Code is processed at the AST level where possible (Go, Python) or via robust regex-based manipulation strategies for languages where AST parsing is prohibitive (C++, Java, Ruby, JavaScript). String and number transformations utilize a unique polymorphic generation strategy per build, ensuring varying signatures for identical input source files.
