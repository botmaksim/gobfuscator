package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"gobfuscator/core"
	"gobfuscator/options"
	"gobfuscator/pkg/utils"

	// Register engines
	_ "gobfuscator/pkg/cpp"
	_ "gobfuscator/pkg/golang"
	_ "gobfuscator/pkg/java"
	_ "gobfuscator/pkg/javascript"
	_ "gobfuscator/pkg/python"
	_ "gobfuscator/pkg/ruby"
)

func detectLang(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".go":
		return "go"
	case ".cpp", ".cc", ".cxx", ".h", ".hpp":
		return "cpp"
	case ".java":
		return "java"
	case ".js", ".cjs", ".mjs":
		return "js"
	case ".py":
		return "python"
	case ".rb":
		return "ruby"
	}
	return ""
}

func main() {
	var lang string
	var output string
	var rename bool
	var stringsOpt bool
	var numbers bool
	var deadcode bool
	var flatten bool
	var split bool
	var hide bool
	var junk bool
	var antidebug bool
	var mba bool
	var opaque bool

	var rootCmd = &cobra.Command{
		Use:   "gobfuscator [input_path]",
		Short: "Advanced Multi-Language Code Obfuscator",
		Long: `Gobfuscator is an advanced, multi-pass code obfuscator designed to protect intellectual property.
It supports multiple languages (Go, Python, C++, Java, JavaScript, Ruby) and offers various
transformations including Control Flow Flattening, Mixed Boolean-Arithmetic, and String Encryption.

The obfuscator can process single files or recursively process entire directories.
Language detection is automatic based on file extensions, unless explicitly overridden.`,
		Example: `  # Obfuscate a single file with auto-detected language
  gobfuscator main.go -o main_obfuscated.go

  # Obfuscate an entire project directory
  gobfuscator ./src -o ./obfuscated_src

  # Disable specific protections
  gobfuscator script.py -o script_obf.py --anti-debug=false --flatten=false

  # Force a specific language engine
  gobfuscator code.txt -o obfuscated.txt --lang=cpp`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			inputPath := args[0]
			stat, err := os.Stat(inputPath)
			if err != nil {
				fmt.Printf("[!] Error accessing input path: %v\n", err)
				os.Exit(1)
			}

			opts := options.Options{
				RenameVariables:  rename,
				ObfuscateStrings: stringsOpt,
				ObfuscateNumbers: numbers,
				AddDeadCode:      deadcode,
				FlattenFlow:      flatten,
				SplitStrings:     split,
				HideAPI:          hide,
				SemanticJunk:     junk,
				AntiDebug:        antidebug,
				MBA:              mba,
				OpaquePreds:      opaque,
				GlobalNameMap:    make(map[string]string),
			}

			fmt.Println("[*] Gobfuscator Engine Initialized")

			if !stat.IsDir() {
				processSingleFile(inputPath, output, lang, opts)
			} else {
				if output == "" {
					fmt.Println("[!] Error: Output directory (-o) must be specified when processing a directory.")
					os.Exit(1)
				}
				processDirectory(inputPath, output, lang, opts)
			}
		},
	}

	rootCmd.Flags().StringVar(&lang, "lang", "", "Force specific language engine (go, cpp, java, js, python, ruby). If omitted, relies on auto-detection.")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "Destination file or directory for obfuscated code")
	rootCmd.Flags().BoolVar(&rename, "rename", true, "Enable global variable and function renaming")
	rootCmd.Flags().BoolVar(&stringsOpt, "strings", true, "Enable polymorphic string encryption")
	rootCmd.Flags().BoolVar(&numbers, "numbers", true, "Enable integer literal obfuscation")
	rootCmd.Flags().BoolVar(&deadcode, "deadcode", true, "Enable opaque dead code injection")
	rootCmd.Flags().BoolVar(&flatten, "flatten", true, "Enable control flow flattening state machines")
	rootCmd.Flags().BoolVar(&split, "split-strings", true, "Enable string splitting and dynamic concatenation")
	rootCmd.Flags().BoolVar(&hide, "hide-api", true, "Enable standard library API aliasing")
	rootCmd.Flags().BoolVar(&junk, "semantic-junk", true, "Enable semantic junk structure generation")
	rootCmd.Flags().BoolVar(&antidebug, "anti-debug", true, "Enable anti-debugging and timing countermeasures")
	rootCmd.Flags().BoolVar(&mba, "mba", true, "Enable Mixed Boolean-Arithmetic mathematical rewriting")
	rootCmd.Flags().BoolVar(&opaque, "opaque-preds", true, "Enable complex logical opaque predicates")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func processSingleFile(inputFile, outputFile, forcedLang string, opts options.Options) {
	lang := forcedLang
	if lang == "" {
		lang = detectLang(inputFile)
		if lang != "" {
			fmt.Printf("[*] Auto-detected language: %s\n", lang)
		}
	} else {
		fmt.Printf("[*] Forced language engine: %s\n", lang)
	}

	source, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("[!] Error reading input file: %v\n", err)
		os.Exit(1)
	}

	engine, exists := core.Get(lang)
	if !exists {
		fmt.Printf("[!] Unsupported or undetected language: %s\n", lang)
		os.Exit(1)
	}

	if opts.RenameVariables {
		vars := utils.ExtractVariables(string(source), lang)
		for _, v := range vars {
			if _, ok := opts.GlobalNameMap[v]; !ok {
				opts.GlobalNameMap[v] = utils.RandomName()
			}
		}
	}

	obfuscated, err := engine.Obfuscate(source, opts)
	if err != nil {
		fmt.Printf("[!] Error during obfuscation: %v\n", err)
		os.Exit(1)
	}

	if outputFile == "" {
		fmt.Println(string(obfuscated))
	} else {
		err = ioutil.WriteFile(outputFile, obfuscated, 0644)
		if err != nil {
			fmt.Printf("[!] Error writing output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("[+] Successfully obfuscated %s to %s\n", inputFile, outputFile)
	}
}

func processDirectory(inputDir, outputDir, forcedLang string, opts options.Options) {
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Printf("[!] Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("[*] Scanning directory for supported files...")

	// Pass 1: Collect globals
	var filesToProcess []string
	err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		lang := forcedLang
		if lang == "" {
			lang = detectLang(path)
		}
		if lang == "" {
			return nil // Skip unsupported files
		}
		
		filesToProcess = append(filesToProcess, path)

		if opts.RenameVariables {
			source, err := ioutil.ReadFile(path)
			if err == nil {
				vars := utils.ExtractVariables(string(source), lang)
				for _, v := range vars {
					if _, ok := opts.GlobalNameMap[v]; !ok {
						opts.GlobalNameMap[v] = utils.RandomName()
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("[!] Error walking input directory during pass 1: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[*] Found %d files to process\n", len(filesToProcess))

	// Pass 2: Obfuscate files
	for _, path := range filesToProcess {
		lang := forcedLang
		if lang == "" {
			lang = detectLang(path)
		}

		engine, exists := core.Get(lang)
		if !exists {
			continue // Should not happen as we checked in Pass 1
		}

		source, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("[!] Warning: Skipping %s due to read error: %v\n", path, err)
			continue
		}

		obfuscated, err := engine.Obfuscate(source, opts)
		if err != nil {
			fmt.Printf("[!] Warning: Failed to obfuscate %s: %v\n", path, err)
			continue
		}

		relPath, _ := filepath.Rel(inputDir, path)
		outPath := filepath.Join(outputDir, relPath)

		os.MkdirAll(filepath.Dir(outPath), 0755)
		err = ioutil.WriteFile(outPath, obfuscated, 0644)
		if err != nil {
			fmt.Printf("[!] Error writing output file %s: %v\n", outPath, err)
		} else {
			fmt.Printf("[+] Obfuscated: %s -> %s (Lang: %s)\n", path, outPath, lang)
		}
	}
	fmt.Println("[*] Directory processing complete")
}
