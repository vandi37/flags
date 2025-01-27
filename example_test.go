package flags_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/vandi37/flags"
)

type Cfg struct {
	Port int    `flag:"port"`
	Host string `flag:"host"`
}

func ExampleArgs() {
	os.Args = []string{os.Args[0]}
	os.Args = append(os.Args, strings.Fields("--port 3700 --host 'localhost'")...)

	cfg := new(Cfg)

	err := flags.Args(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", *cfg)
	// Output: {Port:3700 Host:localhost}

}

func ExampleArgsWithShortcuts() {
	os.Args = []string{os.Args[0]}
	os.Args = append(os.Args, strings.Fields("-ph 3700 'localhost'")...)

	cfg := new(Cfg)

	err := flags.ArgsWithShortcuts(cfg, map[rune]string{'p': "port", 'h': "host"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", *cfg)
	// Output: {Port:3700 Host:localhost}

}
func ExampleLoad() {
	cfg := new(Cfg)

	err := flags.Load(strings.Fields("--port 3700 --host 'localhost'"), cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", *cfg)
	// Output: {Port:3700 Host:localhost}
}

func ExampleLoadWithShortcuts() {
	cfg := new(Cfg)

	err := flags.LoadWithShortcuts(strings.Fields("-ph 3700 'localhost'"), cfg, map[rune]string{'p': "port", 'h': "host"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", *cfg)
	// Output: {Port:3700 Host:localhost}
}

func ExampleParse() {
	f, err := flags.Parse(strings.Fields("--port 3700 --host 'localhost'"))
	if err != nil {
		panic(err)
	}

	for flag, vals := range f {
		fmt.Println(flag)

		for _, val := range vals {
			fmt.Println("-", val)
		}
	}

	// Output:
	// port
	// - 3700
	// host
	// - 'localhost'
}

func ExampleParseWithShortcuts() {
	f, err := flags.ParseWithShortcuts(strings.Fields("-ph 3700 'localhost'"), map[rune]string{'p': "port", 'h': "host"})
	if err != nil {
		panic(err)
	}

	for flag, vals := range f {
		fmt.Println(flag)

		for _, val := range vals {
			fmt.Println("-", val)
		}
	}

	// Output:
	// port
	// - 3700
	// host
	// - 'localhost'
}

func ExampleInsert() {
	cfg := new(Cfg)

	err := flags.Insert(map[string][]string{"port": {"3700"}, "host": {"'localhost'"}}, cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", *cfg)
	// Output: {Port:3700 Host:localhost}
}
