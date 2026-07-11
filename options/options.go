package options

// Options defines the configuration for the obfuscation process.
type Options struct {
	// RenameVariables enables renaming of variables and functions.
	RenameVariables bool
	
	// ObfuscateStrings enables string obfuscation/encryption.
	ObfuscateStrings bool

	// ObfuscateNumbers enables replacing numbers with math expressions.
	ObfuscateNumbers bool

	// AddDeadCode enables injection of opaque predicates and dead code.
	AddDeadCode bool

	// FlattenFlow enables Control Flow Flattening (CFF) for supported languages.
	FlattenFlow bool

	// SplitStrings splits strings into randomized fragments.
	SplitStrings bool

	// HideAPI hides standard library and native API calls behind indirect variables.
	HideAPI bool

	// SemanticJunk injects structurally valid but unused dummy classes, functions, and structs.
	SemanticJunk bool

	// AntiDebug injects debugger detection code (e.g. timing checks or trace hooks).
	AntiDebug bool

	// MBA replaces simple arithmetic with complex Mixed Boolean-Arithmetic equivalents.
	MBA bool

	// OpaquePreds injects control flow branches that are always statically true/false but hard to analyze.
	OpaquePreds bool

	// RubyNoEval disables wrapping the final Ruby script in a hex-encoded eval block.
	RubyNoEval bool

	// JSNoEval disables wrapping the final JS script in a base64-encoded eval block.
	JSNoEval bool

	// GlobalNameMap stores mapping of original names to obfuscated names across the project.
	GlobalNameMap map[string]string
}

// DefaultOptions returns a basic set of options enabled by default.
func DefaultOptions() Options {
	return Options{
		RenameVariables:  true,
		ObfuscateStrings: true,
		ObfuscateNumbers: true,
		AddDeadCode:      true,
		FlattenFlow:      true,
		SplitStrings:     true,
		HideAPI:          true,
		SemanticJunk:     true,
		AntiDebug:        true,
		MBA:              true,
		OpaquePreds:      true,
		GlobalNameMap:    make(map[string]string),
	}
}
