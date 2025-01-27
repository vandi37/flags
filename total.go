package flags

import (
	"os"
)

// it parses command-line arguments and loads the results into a struct.
//
// It combines the functionality of [Parse] and [Insert] to directly
// populate a struct with values extracted from command-line flags.
//
// The provided `args` slice represents the command-line arguments to be
// parsed. The `v` parameter should be a pointer to the struct that
// will be populated.
//
// This function does not support short flag shortcuts. For that, use
// [LoadWithShortcuts].
//
// It returns an error if there is an issue during argument parsing or
// struct population.
//
// Full flag forming rules amd flag values parsing rules are in readme
func Load(args []string, v any) error {
	return LoadWithShortcuts(args, v, map[rune]string{})
}

// it parses command-line arguments, including short flag
// shortcuts, and loads the results into a struct.
//
// It combines the functionality of [ParseWithShortcuts] and [Insert]
// to populate a struct with values extracted from command-line flags,
// including handling short flag shortcuts.
//
// The `args` slice contains the command-line arguments to be parsed. The
// `v` parameter is a pointer to the struct that will be populated. The
// `shortcuts` map defines the short flag to full flag mappings.
//
// It returns an error if there is an issue during argument parsing or
// struct population.
//
// Full flag forming rules amd flag values parsing rules are in readme
func LoadWithShortcuts(args []string, v any, shortcuts map[rune]string) error {
	if f, err := ParseWithShortcuts(args, shortcuts); err != nil {
		return err
	} else {
		return Insert(f, v)
	}
}

// ir parses command-line arguments (from [os.Args]) and loads the results
// into a struct.
//
// It is similar to [Load] but uses the command-line arguments provided
// to the current program (i.e., `os.Args[1:]`) instead of taking the
// arguments as an explicit parameter. It does not support short flag shortcuts.
//
// The `v` parameter should be a pointer to the struct to be populated.
//
// It returns an error if there is an issue during argument parsing or
// struct population.
//
// Full flag forming rules amd flag values parsing rules are in readme
func Args(v any) error {
	return ArgsWithShortcuts(v, map[rune]string{})
}

// it parses command-line arguments (from `os.Args`),
// including short flag shortcuts, and loads the results into a struct.
//
// It is similar to [LoadWithShortcuts] but it uses the command-line
// arguments passed to the current program (i.e., `os.Args[1:]`) instead
// of taking the arguments as an explicit parameter.
//
// The `v` parameter is a pointer to the struct to be populated. The
// `shortcuts` map defines the short flag to full flag mappings.
//
// It returns an error if there is an issue during argument parsing or
// struct population.
//
// Full flag forming rules amd flag values parsing rules are in readme
func ArgsWithShortcuts(v any, shortcuts map[rune]string) error {
	return LoadWithShortcuts(os.Args[1:], v, shortcuts)
}
