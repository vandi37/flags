package flags

import "os"

func Load(args []string, v any) error {
	return LoadWithShortcuts(args, v, map[rune]string{})
}

func LoadWithShortcuts(args []string, v any, shortcuts map[rune]string) error {
	if f, err := ParseWithShortcuts(args, shortcuts); err != nil {
		return err
	} else {
		return Insert(f, v)
	}
}

func Args(v any) error {
	return ArgsWithShortcuts(v, map[rune]string{})
}

func ArgsWithShortcuts(v any, shortcuts map[rune]string) error {
	return LoadWithShortcuts(os.Args[1:], v, shortcuts)
}
