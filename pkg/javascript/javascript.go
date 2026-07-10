package javascript

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"gobfuscator/core"
	"gobfuscator/options"
	"gobfuscator/pkg/utils"
)

func init() {
	core.Register("javascript", &JavaScriptObfuscator{})
	core.Register("js", &JavaScriptObfuscator{})
	rand.Seed(time.Now().UnixNano())
}

type JavaScriptObfuscator struct{}

func (j *JavaScriptObfuscator) Obfuscate(source []byte, opts options.Options) ([]byte, error) {
	code := string(source)

	if opts.RenameVariables {
		vars := utils.ExtractVariables(code, "js")
		for _, v := range vars {
			if v == "main" || v == "console" || v == "Math" || v == "Date" || v == "Buffer" {
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
		wrapperFunc := fmt.Sprintf("_0xLog%x", rand.Int31())
		wrapperCode := fmt.Sprintf(`
function %s(msg) {
    console.log(msg);
}
`, wrapperFunc)
		
		rePrint := regexp.MustCompile(`console\.log`)
		if rePrint.MatchString(code) {
			code = rePrint.ReplaceAllString(code, wrapperFunc)
			code = wrapperCode + "\n" + code
		}
	}

	if opts.AntiDebug {
		antiDebugCode := `
(function() {
    var _0xStart = Date.now();
    debugger;
    if (Date.now() - _0xStart > 100) {
        process.exit(1);
    }
})();
`
		code = antiDebugCode + "\n" + code
	}

	if opts.AddDeadCode {
		reFunc := regexp.MustCompile(`(function\s*[\w]*\s*\([^)]*\)\s*\{|=>\s*\{)`)
		code = reFunc.ReplaceAllStringFunc(code, func(match string) string {
			pred := utils.GenerateOpaquePredicate("js", nil)
			body := utils.GenerateDeadCodeBody("js", string(source))
			return fmt.Sprintf("%s\n        if(%s) {\n            %s\n        }\n", match, pred, body)
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
			return fmt.Sprintf("%s\n        if (%s) { throw\"\"; }\n", match, pred)
		})
	}

	code = utils.MangleWhitespace(code)

	if opts.ObfuscateStrings {
		encoded := base64.StdEncoding.EncodeToString([]byte(code))
		decoderSnippet := fmt.Sprintf(`(function(p){var s=typeof Buffer!=='undefined'?Buffer.from(p,'base64').toString('utf8'):atob(p);eval(s);})("%s");`, encoded)
		code = utils.MangleWhitespace(decoderSnippet)
	}

	return []byte(code), nil
}
