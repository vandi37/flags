package flags

import "os"

func Load(args []string, v any, shortcuts map[rune]string) error {
	if f, err := Parse(args, shortcuts); err != nil {
		return err
	} else {
		return Insert(f, v)
	}
}

func Args(v any, shortcuts map[rune]string) error {
	return Load(os.Args, v, shortcuts)
}
