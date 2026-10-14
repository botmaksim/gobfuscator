# Gobfuscator: AST-Driven Source Code Obfuscation Suite

[![Go 1.22+](https://img.shields.io/badge/Go-1.22%2B-00ADD8.svg)](https://go.dev/)
[![Security](https://img.shields.io/badge/Security-Code%20Protection-red.svg)](#)
[![Compiler](https://img.shields.io/badge/Compiler-AST%20Analysis-blue.svg)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Gobfuscator is a multi-pass source code obfuscation utility implemented in Go. By parsing Abstract Syntax Trees, the tool applies AST-level code transformations including symbol mangling, string literal encryption, dead code injection, and control flow flattening to deter reverse engineering and automated static analysis.

---

## Table of Contents
1. [Obfuscation Pipeline](#obfuscation-pipeline)
2. [Transformation Passes](#transformation-passes)
3. [Prerequisites](#prerequisites)
4. [Build & Usage](#build--usage)
5. [Directory Structure](#directory-structure)
6. [License](#license)

---

## Obfuscation Pipeline

1. **Parsing**: Target source code is parsed into an Abstract Syntax Tree using standard AST parser routines.
2. **Analysis**: Symbol definitions, variable scopes, and control branches are mapped into an immutable symbol table.
3. **Mutation Passes**: Configured transformation passes rewrite AST nodes sequentially.
4. **Code Generation**: The transformed AST is emitted as valid, fully executable source code.

---

## Transformation Passes

* **Identifier Scrambling**: Replaces variable, function, and parameter names with uninformative cryptographic hashes.
* **String Encryption**: Encodes string constants into byte arrays decoded at runtime.
* **Dead Code Injection**: Generates mathematically dead conditional branches to increase reverse engineering complexity.
* **Control Flow Flattening**: Replaces linear execution paths with state-variable switch dispatch loops.

---

## Prerequisites

* Go 1.22 or higher

---

## Build & Usage

```bash
git clone git@github.com:botmaksim/gobfuscator.git
cd gobfuscator

go build -o gobfuscator main.go
./gobfuscator -help
```

---

## Directory Structure

```text
gobfuscator/
├── main.go                  # Command-line entry point and flags
├── go.mod                   # Go module definition
├── core/                    # AST parsing and visitor abstractions
├── pkg/                     # Concrete transformation passes
└── options/                 # Obfuscation configuration options
```

---

## License

This project is open-source software licensed under the [MIT License](LICENSE).
