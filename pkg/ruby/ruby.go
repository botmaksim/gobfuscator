package ruby

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"math/rand"
	"regexp"

	"gobfuscator/core"
	"gobfuscator/options"
	"gobfuscator/pkg/utils"
)

func init() {
	core.Register("ruby", &RubyObfuscator{})
	core.Register("rb", &RubyObfuscator{})
}

type RubyObfuscator struct{}

func (r *RubyObfuscator) Obfuscate(source []byte, opts options.Options) ([]byte, error) {
	code := string(source)

	if opts.RenameVariables {
		vars := utils.ExtractVariables(code, "ruby")
		for _, v := range vars {
			if v == "main" || v == "require" {
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
		wrapperFunc := utils.RandomName()
		wrapperCode := fmt.Sprintf(`
def %s(*args)
  puts(*args)
end
`, wrapperFunc)
		
		rePuts := regexp.MustCompile(`\bputs\b`)
		if rePuts.MatchString(code) {
			code = rePuts.ReplaceAllString(code, wrapperFunc)
			code = wrapperCode + "\n" + code
		}
	}

	if opts.AntiDebug {
		antiDebugCode := `
if defined?(Byebug) || defined?(Debugger)
  exit(1)
end
`
		code = antiDebugCode + "\n" + code
	}

	if opts.AddDeadCode {
		reFunc := regexp.MustCompile(`(?m)^(def\s+[a-zA-Z0-9_]+\b.*)$`)
		code = reFunc.ReplaceAllStringFunc(code, func(match string) string {
			pred := utils.GenerateOpaquePredicate("ruby", nil)
			body := utils.GenerateDeadCodeBody("ruby", string(source))
			return fmt.Sprintf("%s\n  if %s\n    %s\n  end\n", match, pred, body)
		})
	}

	if opts.OpaquePreds {
		reFunc := regexp.MustCompile(`(?m)^(def\s+[a-zA-Z0-9_]+\b.*)$`)
		code = reFunc.ReplaceAllStringFunc(code, func(match string) string {
			if rand.Float32() > 0.3 {
				return match
			}
			randVal := rand.Intn(100) + 1
			pred := fmt.Sprintf("((%d * %d + %d) %% 2 != 0)", randVal, randVal, randVal)
			return fmt.Sprintf("%s\n  if %s\n    exit(1)\n  end\n", match, pred)
		})
	}

	if opts.ObfuscateStrings {
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		w.Write([]byte(code))
		w.Close()
		encoded := base64.StdEncoding.EncodeToString(b.Bytes())
		
		decoderSnippet := fmt.Sprintf("require 'zlib'\nrequire 'base64'\neval(Zlib::Inflate.inflate(Base64.decode64(\"%s\")))", encoded)
		
		code = decoderSnippet
	}

	return []byte(code), nil
}
