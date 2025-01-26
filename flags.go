package flags

import (
	"strings"
)

func Parse(args []string, shortcuts map[rune]string) (map[string][]string, error) {
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
		err = Mega("got some errors", errs)
	}

	return res, err
}
