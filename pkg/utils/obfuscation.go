package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomName generates a random obfuscated name.
func RandomName() string {
	return fmt.Sprintf("_0x%x", rand.Int31())
}

// ExtractVariables uses simple regex to find variable names in C++/Java style code.
func ExtractVariables(code, lang string) []string {
	var vars []string
	var re *regexp.Regexp

	switch lang {
	case "java", "cpp", "c++":
		re = regexp.MustCompile(`\b(int|String|char|float|double|auto)\s+([a-zA-Z_]\w*)\s*=`)
	case "js", "javascript":
		re = regexp.MustCompile(`\b(var|let|const)\s+([a-zA-Z_]\w*)\s*=`)
	case "python":
		re = regexp.MustCompile(`(?m)^\s*([a-zA-Z_]\w*)\s*=`)
	case "ruby":
		re = regexp.MustCompile(`(?m)^\s*([a-zA-Z_]\w*)\s*=`)
	default:
		return vars
	}

	matches := re.FindAllStringSubmatch(code, -1)
	for _, m := range matches {
		if len(m) > 2 {
			vars = append(vars, m[2])
		} else if len(m) > 1 {
			vars = append(vars, m[1])
		}
	}
	return vars
}

// GenerateOpaquePredicate generates an if condition that always evaluates to false.
func GenerateOpaquePredicate(lang string, vars []string) string {
	templates := []string{}
	varBase := ""
	if len(vars) > 0 {
		varBase = vars[rand.Intn(len(vars))]
	}

	switch lang {
	case "java":
		templates = append(templates, "(System.currentTimeMillis() < 0L)", "(Math.sin(0) > 1)", "(\"a\".length() == -1)")
		if varBase != "" {
			templates = append(templates, fmt.Sprintf("(String.valueOf(%s).length() == -1)", varBase))
		}
	case "cpp", "c++":
		templates = append(templates, "(sizeof(char) == 0)", "((new int[1]) == nullptr)")
		if varBase != "" {
			templates = append(templates, fmt.Sprintf("(sizeof(%s) == 0)", varBase))
		}
	case "js", "javascript":
		templates = append(templates, "(Math.random() > 2)", "(Date.now() === -1)")
		if varBase != "" {
			templates = append(templates, fmt.Sprintf("(%s === undefined && %s === null && false)", varBase, varBase))
		}
	case "python":
		templates = append(templates, "(len(str(id(0))) == -1)", "(max(1, 2) == 0)")
		if varBase != "" {
			templates = append(templates, fmt.Sprintf("(id(%s) == 0)", varBase))
		}
	case "ruby":
		templates = append(templates, "(rand > 2)", "(Time.now.to_i == -1)")
		if varBase != "" {
			templates = append(templates, fmt.Sprintf("(%s.nil? && false)", varBase))
		}
	case "go", "golang":
		templates = append(templates, "(len(make([]int, 0)) < 0)", "(1 == 0 && true)")
		if varBase != "" {
			templates = append(templates, fmt.Sprintf("(%s == nil && false)", varBase))
		}
	default:
		templates = append(templates, "(false)")
	}

	return templates[rand.Intn(len(templates))]
}

// GenerateRandomMathExpression creates a deeply nested math expression
func GenerateRandomMathExpression(depth int) string {
	if depth <= 0 {
		return fmt.Sprintf("%d", rand.Intn(1000))
	}
	ops := []string{"+", "-", "^", "*", "<<"}
	op := ops[rand.Intn(len(ops))]
	
	if op == "<<" {
		return fmt.Sprintf("(%s %s %d)", GenerateRandomMathExpression(depth-1), op, rand.Intn(3)+1)
	}
	return fmt.Sprintf("(%s %s %s)", GenerateRandomMathExpression(depth-1), op, GenerateRandomMathExpression(depth-1))
}

// GenerateDeadCodeBody generates garbage math and optionally safe snippets from original code.
func GenerateDeadCodeBody(lang, originalCode string) string {
	r1, r2 := rand.Intn(1000), rand.Intn(100)
	mathExpr := GenerateRandomMathExpression(rand.Intn(3)+2)
	
	switch lang {
	case "cpp", "c++":
		id := rand.Int31()
		return fmt.Sprintf("int _0x%x = %s;\nif(_0x%x > 0) { _0x%x = %d; }", id, mathExpr, id, id, r1)
	case "java":
		id := rand.Int31()
		return fmt.Sprintf("long _0x%x = %s;\nif(_0x%x > 0) { _0x%x = %d; }", id, mathExpr, id, id, r2)
	case "go", "golang":
		return fmt.Sprintf("_0x%x := %s\nif _0x%x > 0 {\n\t_ = %d\n}", rand.Int31(), mathExpr, rand.Int31(), r1)
	case "js", "javascript":
		return fmt.Sprintf("var _0x%x = %s;", rand.Int31(), mathExpr)
	case "ruby":
		return fmt.Sprintf("_0x%x = %s", rand.Int31(), mathExpr)
	case "python", "py":
		return fmt.Sprintf("_0x%x = %s", rand.Int31(), mathExpr)
	default:
		return fmt.Sprintf("var _0x%x = %d;", rand.Int31(), r1)
	}
}

type StringOperation struct {
	OpType string
	Key    byte
}

type PolymorphicStringObfuscator struct {
	FuncName   string
	Operations []StringOperation
}

func NewPolymorphicStringObfuscator() *PolymorphicStringObfuscator {
	ops := []string{"XOR", "ADD", "SUB", "NOT"}
	numOps := rand.Intn(3) + 2
	
	so := &PolymorphicStringObfuscator{
		FuncName: fmt.Sprintf("_0xStr%x", rand.Int31()),
	}
	
	for i := 0; i < numOps; i++ {
		so.Operations = append(so.Operations, StringOperation{
			OpType: ops[rand.Intn(len(ops))],
			Key:    byte(rand.Intn(250) + 1),
		})
	}
	return so
}

func (so *PolymorphicStringObfuscator) Encrypt(s string) []byte {
	bytes := []byte(s)
	for i := range bytes {
		val := bytes[i]
		for _, op := range so.Operations {
			switch op.OpType {
			case "XOR": val = val ^ op.Key
			case "ADD": val = val + op.Key
			case "SUB": val = val - op.Key
			case "NOT": val = ^val
			}
		}
		bytes[i] = val
	}
	return bytes
}

func (so *PolymorphicStringObfuscator) GenerateCppDecryptor() string {
	var revOps strings.Builder
	for i := len(so.Operations) - 1; i >= 0; i-- {
		op := so.Operations[i]
		switch op.OpType {
		case "XOR": revOps.WriteString(fmt.Sprintf("val = val ^ %d;\n", op.Key))
		case "ADD": revOps.WriteString(fmt.Sprintf("val = val - %d;\n", op.Key))
		case "SUB": revOps.WriteString(fmt.Sprintf("val = val + %d;\n", op.Key))
		case "NOT": revOps.WriteString("val = ~val;\n")
		}
	}
	
	return fmt.Sprintf(`
inline const char* %s(const char* data, int len) {
    char* res = (char*)malloc(len + 1);
    for(int i = 0; i < len; ++i) {
        unsigned char val = data[i];
        %s
        res[i] = (char)val;
    }
    res[len] = '\0';
    return res;
}
`, so.FuncName, revOps.String())
}

func (so *PolymorphicStringObfuscator) GenerateJavaDecryptor() string {
	var revOps strings.Builder
	for i := len(so.Operations) - 1; i >= 0; i-- {
		op := so.Operations[i]
		switch op.OpType {
		case "XOR": revOps.WriteString(fmt.Sprintf("val = (byte)(val ^ %d);\n", op.Key))
		case "ADD": revOps.WriteString(fmt.Sprintf("val = (byte)(val - %d);\n", op.Key))
		case "SUB": revOps.WriteString(fmt.Sprintf("val = (byte)(val + %d);\n", op.Key))
		case "NOT": revOps.WriteString("val = (byte)(~val);\n")
		}
	}
	
	return fmt.Sprintf(`
    public static String %s(byte[] data) {
        byte[] res = new byte[data.length];
        for(int i = 0; i < data.length; ++i) {
            byte val = data[i];
            %s
            res[i] = val;
        }
        return new String(res);
    }
`, so.FuncName, revOps.String())
}

// MangleWhitespace randomly replaces spaces with tabs/newlines outside of strings
func MangleWhitespace(code string) string {
	var out strings.Builder
	
	lines := strings.Split(code, "\n")
	inPreprocessor := false
	
	for i, line := range lines {
		trimLine := strings.TrimSpace(line)
		
		if strings.HasPrefix(trimLine, "#") {
			inPreprocessor = true
		}
		
		canAddNewlines := !inPreprocessor && !strings.Contains(line, "//") && !strings.Contains(line, "/*")
		
		if inPreprocessor || strings.HasPrefix(trimLine, "//") {
			out.WriteString(line)
		} else {
			inString := false
			var stringChar rune
			for _, r := range line {
				if (r == '"' || r == '\'' || r == '`') && !inString {
					inString = true
					stringChar = r
					out.WriteRune(r)
					continue
				}
				if r == stringChar && inString {
					inString = false
					out.WriteRune(r)
					continue
				}

				if !inString && r == ' ' {
					roll := rand.Intn(100)
					if roll < 15 {
						out.WriteString("\t")
					} else if roll < 30 && canAddNewlines {
						out.WriteString("\n\t")
					} else if roll < 45 {
						out.WriteString("  ")
					} else {
						out.WriteRune(' ')
					}
				} else {
					out.WriteRune(r)
				}
			}
		}
		
		if i < len(lines)-1 {
			out.WriteString("\n")
		}
		
		if inPreprocessor && !strings.HasSuffix(trimLine, "\\") {
			inPreprocessor = false
		}
	}

	return out.String()
}
