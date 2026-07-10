package cpp

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gobfuscator/core"
	"gobfuscator/options"
	"gobfuscator/pkg/utils"
)

func init() {
	core.Register("cpp", &CppObfuscator{})
	core.Register("c++", &CppObfuscator{})
}

type CppObfuscator struct{}

func (c *CppObfuscator) Obfuscate(source []byte, opts options.Options) ([]byte, error) {
	tmpDir, err := ioutil.TempDir("", "gobfuscator_cpp_")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	inputPath := filepath.Join(tmpDir, "input.cpp")

	err = ioutil.WriteFile(inputPath, source, 0644)
	if err != nil {
		return nil, err
	}

	code := string(source)


	if opts.RenameVariables {
		vars := utils.ExtractVariables(code, "cpp")
		for _, v := range vars {
			if v == "main" {
				continue
			}
			
			newName, ok := opts.GlobalNameMap[v]
			if !ok {
				newName = utils.RandomName()
				opts.GlobalNameMap[v] = newName
			}
			
			reVar := regexp.MustCompile(`\b` + regexp.QuoteMeta(v) + `\b`)
			code = reVar.ReplaceAllString(code, newName)
		}
	}

	if opts.MBA {
		// x + y -> ((x ^ y) + 2 * (x & y))
		reAdd := regexp.MustCompile(`\b([a-zA-Z_]\w*|\d+)\s*\+\s*([a-zA-Z_]\w*|\d+)\b`)
		code = reAdd.ReplaceAllString(code, "(($1 ^ $2) + 2 * ($1 & $2))")

		// x - y -> ((x ^ ~$y) + 2 * (x & ~$y) + 1)
		reSub := regexp.MustCompile(`\b([a-zA-Z_]\w*|\d+)\s*-\s*([a-zA-Z_]\w*|\d+)\b`)
		code = reSub.ReplaceAllString(code, "(($1 ^ ~$2) + 2 * ($1 & ~$2) + 1)")
	}

	if opts.AddDeadCode {
		reFunc := regexp.MustCompile(`\)\s*\{`)
		code = reFunc.ReplaceAllStringFunc(code, func(match string) string {
			pred := utils.GenerateOpaquePredicate("cpp", nil)
			body := utils.GenerateDeadCodeBody("cpp", string(source))
			body = strings.ReplaceAll(body, "\n", "\n        ")
			return fmt.Sprintf(") {\n    if(%s) {\n        %s\n    }\n", pred, body)
		})
	}

	if opts.OpaquePreds {
		// Inject an opaque predicate that evaluates to false at the start of code blocks
		reBlock := regexp.MustCompile(`(\)\s*\{|else\s*\{)`)
		code = reBlock.ReplaceAllStringFunc(code, func(match string) string {
			if rand.Float32() > 0.3 {
				return match
			}
			randVal := rand.Intn(100) + 1
			pred := fmt.Sprintf("((%d * %d + %d) %% 2 != 0)", randVal, randVal, randVal)
			return fmt.Sprintf("%s\n    if (%s) { while(1); }\n", match, pred)
		})
	}

	if opts.ObfuscateStrings {
		re := regexp.MustCompile(`"([^"\\]*)"`)
		
		poly := utils.NewPolymorphicStringObfuscator()
		decryptor := poly.GenerateCppDecryptor()
		
		hasStrings := false
		code = re.ReplaceAllStringFunc(code, func(match string) string {
			hasStrings = true
			content := match[1 : len(match)-1]
			encBytes := poly.Encrypt(content)
			var hexArr string
			for _, b := range encBytes {
				hexArr += fmt.Sprintf("\\x%02x", b)
			}
			return fmt.Sprintf("%s(\"%s\", %d)", poly.FuncName, hexArr, len(content))
		})
		
		if hasStrings {
			code = "#include <string>\n#include <stdlib.h>\n" + decryptor + "\n" + code
		}
	}

	if opts.AntiDebug {
		adID := rand.Int31()
		adCode := fmt.Sprintf(`
#include <chrono>
#include <stdlib.h>
class _0xAD%x {
    std::chrono::high_resolution_clock::time_point s;
public:
    _0xAD%x() { s = std::chrono::high_resolution_clock::now(); }
    ~_0xAD%x() { 
        auto e = std::chrono::high_resolution_clock::now();
        if (std::chrono::duration_cast<std::chrono::seconds>(e - s).count() > 5) {
            exit(1);
        }
    }
} _0x_global_ad_%x;
`, adID, adID, adID, adID)
		code = adCode + code
	}

	// Macro Dictionary Obfuscation (API Hiding & Junk)
	if opts.HideAPI {
		var macros strings.Builder
		keywords := []string{"int", "float", "double", "char", "if", "else", "return", "while", "for", "void", "cout", "cin", "printf", "malloc", "vector", "string", "endl"}
		for _, kw := range keywords {
			macroName := fmt.Sprintf("_0xMac%x", rand.Int31())
			macros.WriteString(fmt.Sprintf("#define %s %s\n", macroName, kw))
			
			// Lookbehind to avoid matching #if, #else, #include <string>, etc.
			reKw := regexp.MustCompile(`(?m)(^|[^#<"\w])\b` + kw + `\b`)
			code = reKw.ReplaceAllString(code, "${1}"+macroName)
		}
		// Insert macros after all includes
		reInc := regexp.MustCompile(`(?m)^(?:#include.*[\r\n]+)+`)
		match := reInc.FindString(code)
		if match != "" {
			code = strings.Replace(code, match, match+macros.String()+"\n", 1)
		} else {
			code = macros.String() + "\n" + code
		}
	}

	// Always mangle whitespace for C++
	code = utils.MangleWhitespace(code)

	return []byte(code), nil
}
