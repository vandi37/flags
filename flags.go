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
