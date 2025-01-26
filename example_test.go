package flags_test

import (
	"fmt"
	"strings"

	"github.com/vandi37/flags"
)

type ExampleCfg struct {
	Port int    `flag:"port"`
	Host string `flag:"host"`
}

func ExampleInsert() {
	f, err := flags.Parse(strings.Fields("--port 3700 --host 'localhost'"), map[rune]string{})
	if err != nil {
		panic(err)
	}

	cfg := new(ExampleCfg)

	err = flags.Insert(f, cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", *cfg)
	// Output: {Port:3700 Host:localhost}
}
