// Flags are an alternative to the [flag standard package](https://pkg.go.dev/flag).
//
// # Flag Format
//
// # Here are some rules for flags
//
// - Normal Flags
//
// 1. A flag **always** starts with "--"
//
// Example:
// `--flag`
//
// 2. A flag may have values
//
// Example:
// `--flag value`
//
// Flags can have multiple values.
//
// `--flag value1 value2 value3 ...`
//
// - Shortcut Flags
//
// Shortcut always starts with.
//
// 1. Shortcuts **always** starts with "-"
//
//	Example, shortcut 'f' for flag "flag":
//
// `-f`
//
// 2. You can have multiple shortcuts.
//
// Example shortcut 'f' for flag "flag" and 'o' to flag "other_flag":
// `-fo`
//
// 3. Singular shortcut has the same rules as normal flags.
//
// Example shortcut 'f' for flag "flag":
// `-f value`
//
// `-f value1 value2 value3 ...`
//
// 4. Each shortcut takes only one value.
//
// Example shortcut 'f' to flag "flag" and 'o' to flag "other_flag":
// `-fo value_for_flag value_for_other_flag`
//
// With multiple shortcuts, the last one will use any remaining values.
//
// `-fo value_for_flag value_for_other_flag also_value_for_other_flag ...`
//
// # Converting Flags to Types
//
// Flags may be inserted into a structure, here are the rules:
//
// - Integers (int, int8...int64, uint, uint8...uint64, uintptr, time.Duration, unsafe.Pointer)
//
// 1. Convert to base 10.
//
// 2. Convert from base 2 (binary) to base 16 (hex).
//
// 3. For integers (int, int8...int64) convert to time.Duration
//
// - Boolean
//
// 1. Without arguments, it's true.
//
// 2. Convert to bool using strconv.ParseBool
//
// - float (float32, float64)
//
// # Convert to float using strconv.ParseFloat
//
// - Complex (complex64, complex128)
//
// # Convert to complex using strconv.ParseComplex
//
// - String
//
// If there are brackets (", ', `) , trim them and get the string
//
// !! it won't convert the string without brackets
//
// - Array
//
// It goes through the array and fills it with multiple values for each value. It uses the same conversion rule for all values in the array.
//
// !! for arrays and slices it won't work if there are 2d, 3d ... arrays/slices
//
// - Slice
//
// It starts with the last element in the slice, and appends all the values (multiple values), for each value, using the same conversion rule.
//
// - Time
//
// Parse using all time formats. You can specify your own formats.
//
// - interface
//
// If it's an empty interface (interface{}), use default conversion
//
// 1. string
// 2. int
// 3. float
// 4. bool
// 5. time
// 6. complex
//
// - Pointer
//
// # Convert using same conversion rules to type pointed to by pointer
//
// - Struct
//
// Do same conversion with same flags on this struct.
//
// !!  channels, maps, functions are not supported
package flags

import (
	"strings"
)

// it parses a slice of strings into a map of flag names to their values.
//
// It processes command-line arguments, recognizing flags (e.g., "--flag")
// and their associated values. Multiple values for the same flag are allowed.
//
// It returns a map where keys are flag names (without the leading "--") and
// values are slices of strings representing the values for that flag.
//
// This function does not support short flag shortcuts. For that, use
// [ParseWithShortcuts].
//
// Full flag forming rules are in readme
func Parse(args []string) (map[string][]string, error) {
	return ParseWithShortcuts(args, map[rune]string{})
}

// it parses a slice of strings into a map of flag names to
// their values, allowing for short flag shortcuts.
//
// It works similarly to [Parse], but also recognizes short flag shortcuts
// (e.g., "-f" which could be mapped to "--flag", also -fo could me mapped to "--flag" and "--other_flag"). The `shortcuts` map defines
// these mappings, where the key is the shortcut rune (e.g., 'f') and the
// value is the full flag name (e.g., "flag").
//
// It returns a map where keys are full flag names (without the leading "--")
// and values are slices of strings representing the values for that flag.
//
// Full flag forming rules are in readme
func ParseWithShortcuts(args []string, shortcuts map[rune]string) (map[string][]string, error) {
	res := make(map[string][]string)
	errs := []error{}

	currentFlags := []string{}
	for _, el := range args {
		if strings.HasPrefix(el, "--") {
			el = strings.TrimPrefix(el, "--")
			if _, ok := res[el]; ok {
				errs = append(errs, TWICE_FLAG(el))
				continue
			}

			res[el] = []string{}
			currentFlags = []string{el}
		} else if strings.HasPrefix(el, "-") {
			el = strings.TrimPrefix(el, "-")
			currentFlags = []string{}

			for _, s := range el {
				var ok bool
				var fl string
				if fl, ok = shortcuts[s]; !ok {
					errs = append(errs, WRONG_SHORTCUT(s))
					continue
				}

				res[fl] = []string{}
				currentFlags = append(currentFlags, fl)
			}
		} else {
			if len(currentFlags) <= 0 {
				errs = append(errs, ARGUMENT_NOT_NEED(el))
				continue
			}

			if len(currentFlags) == 1 {
				res[currentFlags[0]] = append(res[currentFlags[0]], el)
				continue
			}

			if len(currentFlags) > 1 {
				res[currentFlags[0]] = append(res[currentFlags[0]], el)
				currentFlags = currentFlags[1:]
			}

		}
	}

	var err error
	if len(errs) > 0 {
		err = mega("got some errors", errs)
	}

	return res, err
}
