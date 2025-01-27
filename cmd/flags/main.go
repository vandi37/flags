package main

import (
	"fmt"
	"os"

	"github.com/vandi37/flags"
)

func main() {
	res, err := flags.Parse(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, "got error:", err.Error())
	}

	fmt.Println("your flags:", string_flags(res))
}

func string_flags(flags map[string][]string) string {
	var res string
	for flag, vals := range flags {
		res += "\n" + flag

		for _, val := range vals {
			res += "\n- " + val
		}
	}

	return res
}
