package java

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"gobfuscator/core"
	"gobfuscator/options"
	"gobfuscator/pkg/utils"
)

func init() {
	core.Register("java", &JavaObfuscator{})
}

type JavaObfuscator struct{}

func (j *JavaObfuscator) Obfuscate(source []byte, opts options.Options) ([]byte, error) {
	code := string(source)


	if opts.RenameVariables {
		vars := utils.ExtractVariables(code, "java")
		for _, v := range vars {
			if v == "main" || v == "args" || v == "String" {
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
		reAdd := regexp.MustCompile(`\b([a-zA-Z_]\w*|\d+)\s*\+\s*([a-zA-Z_]\w*|\d+)\b`)
		code = reAdd.ReplaceAllString(code, "(($1 ^ $2) + 2 * ($1 & $2))")

		reSub := regexp.MustCompile(`\b([a-zA-Z_]\w*|\d+)\s*-\s*([a-zA-Z_]\w*|\d+)\b`)
		code = reSub.ReplaceAllString(code, "(($1 ^ ~$2) + 2 * ($1 & ~$2) + 1)")
	}

	if opts.HideAPI {
		wrapperClass := fmt.Sprintf("_0xSys%x", rand.Int31())
		wrapperMethod := fmt.Sprintf("_0xOut%x", rand.Int31())
		wrapperCode := fmt.Sprintf(`
class %s {
    public static void %s(Object o) {
        System.out.println(o);
    }
}
`, wrapperClass, wrapperMethod)
		
		rePrint := regexp.MustCompile(`System\.out\.println`)
		if rePrint.MatchString(code) {
			code = rePrint.ReplaceAllString(code, wrapperClass+"."+wrapperMethod)
			code = code + "\n" + wrapperCode
		}
	}

	if opts.AntiDebug {
		reMain := regexp.MustCompile(`(public\s+static\s+void\s+main\s*\(\s*String\[\]\s+\w+\s*\)\s*\{)`)
		antiDebugCode := `
        if (java.lang.management.ManagementFactory.getRuntimeMXBean().getInputArguments().toString().contains("-agentlib:jdwp")) {
            System.exit(1);
        }
`
		code = reMain.ReplaceAllString(code, "$1"+antiDebugCode)
	}

	if opts.AddDeadCode {
		reFunc := regexp.MustCompile(`\)\s*\{`)
		code = reFunc.ReplaceAllStringFunc(code, func(match string) string {
			pred := utils.GenerateOpaquePredicate("java", nil)
			body := utils.GenerateDeadCodeBody("java", string(source))
			body = strings.ReplaceAll(body, "\n", "\n            ")
			return fmt.Sprintf(") {\n        if(%s) {\n            %s\n        }\n", pred, body)
		})
	}

	if opts.OpaquePreds {
		reBlock := regexp.MustCompile(`(\)\s*\{|else\s*\{)`)
		code = reBlock.ReplaceAllStringFunc(code, func(match string) string {
			if rand.Float32() > 0.3 {
				return match
			}
			randVal := rand.Intn(100) + 1
			pred := fmt.Sprintf("((%d * %d + %d) %% 2 != 0)", randVal, randVal, randVal)
			randVal2 := rand.Intn(100) + 1
			pred2 := fmt.Sprintf("((%d * %d + %d) %% 2 != 0)", randVal2, randVal2, randVal2)
			body := utils.GenerateDeadCodeBody("java", "")
			return fmt.Sprintf("%s\n        if (%s) { while(%s){ %s } }\n", match, pred, pred2, body)
		})
	}

	if opts.ObfuscateStrings {
		re := regexp.MustCompile(`"([^"\\]*)"`)
		
		poly := utils.NewPolymorphicStringObfuscator()
		decryptorClass := `
class _0xDec {
` + poly.GenerateJavaDecryptor() + `
}
`
		hasStrings := false
		code = re.ReplaceAllStringFunc(code, func(match string) string {
			hasStrings = true
			content := match[1 : len(match)-1]
			encBytes := poly.Encrypt(content)
			var hexArr string
			for i, b := range encBytes {
				if i > 0 {
					hexArr += ", "
				}
				hexArr += fmt.Sprintf("(byte)0x%02x", b)
			}
			return fmt.Sprintf("_0xDec.%s(new byte[]{%s})", poly.FuncName, hexArr)
		})
		
		if hasStrings {
			code = code + "\n" + decryptorClass
		}
	}

	code = utils.MangleWhitespace(code)

	return []byte(code), nil
}
