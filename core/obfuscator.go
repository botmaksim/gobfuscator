package core

import "gobfuscator/options"

// Obfuscator defines the contract for all language-specific obfuscators.
type Obfuscator interface {
	// Obfuscate takes the source code and configuration, returning the obfuscated code.
	Obfuscate(source []byte, opts options.Options) ([]byte, error)
}

var registry = make(map[string]Obfuscator)

// Register adds an obfuscator to the registry.
func Register(lang string, obf Obfuscator) {
	registry[lang] = obf
}

// Get retrieves an obfuscator from the registry.
func Get(lang string) (Obfuscator, bool) {
	obf, exists := registry[lang]
	return obf, exists
}
