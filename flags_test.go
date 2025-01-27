package flags_test

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/vandi37/flags"
)

type okCase struct {
	name      string
	args      []string
	shortcuts map[rune]string
	res       map[string][]string
}

func TestOk(t *testing.T) {
	cases := []okCase{
		{
			name:      "easy",
			args:      []string{"--flag", "value"},
			shortcuts: map[rune]string{},
			res: map[string][]string{
				"flag": {"value"},
			},
		},
		{
			name:      "no value",
			args:      []string{"--flag"},
			shortcuts: map[rune]string{},
			res: map[string][]string{
				"flag": {},
			},
		},
		{
			name:      "many values",
			args:      []string{"--flag", "value", "other_value"},
			shortcuts: map[rune]string{},
			res: map[string][]string{
				"flag": {"value", "other_value"},
			},
		},
		{
			name:      "1 shortcut",
			args:      []string{"-f", "value"},
			shortcuts: map[rune]string{'f': "flag"},
			res: map[string][]string{
				"flag": {"value"},
			},
		},
		{
			name:      "multiple shortcuts",
			args:      []string{"-ft", "value", "other_value"},
			shortcuts: map[rune]string{'f': "flag", 't': "test"},
			res: map[string][]string{
				"flag": {"value"},
				"test": {"other_value"},
			},
		},
		{
			name:      "combo",
			args:      []string{"-ft", "value", "other_value", "--other_flag", "one_more_value", "-s", "value_for_shortcut"},
			shortcuts: map[rune]string{'f': "flag", 't': "test", 's': "shortcut"},
			res: map[string][]string{
				"flag":       {"value"},
				"test":       {"other_value"},
				"other_flag": {"one_more_value"},
				"shortcut":   {"value_for_shortcut"},
			},
		},
		{
			name:      "shortcut with many values",
			args:      []string{"-fts", "value_for_first_flag", "value_for_second_flag", "value_for_third_flag", "value_also_for_third_flag", "this_value_also_is_for_third_flag"},
			shortcuts: map[rune]string{'f': "flag", 't': "test", 's': "shortcut"},
			res: map[string][]string{
				"flag":     {"value_for_first_flag"},
				"test":     {"value_for_second_flag"},
				"shortcut": {"value_for_third_flag", "value_also_for_third_flag", "this_value_also_is_for_third_flag"},
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("#%d %s", i, tc.name), func(t *testing.T) {
			res, err := flags.ParseWithShortcuts(tc.args, tc.shortcuts)
			if err != nil {
				t.Fatalf("got an error: %v", err)
			}

			if !maps.EqualFunc(tc.res, res, func(v1, v2 []string) bool { return slices.Equal(v1, v2) }) {
				t.Fatalf("got different maps: expected %v, got %v", tc.res, res)
			}
		})
	}
}

type errorCase struct {
	name      string
	args      []string
	shortcuts map[rune]string
	errs      []error
	res       map[string][]string
}

func TestError(t *testing.T) {
	cases := []errorCase{
		{
			name:      "no flags",
			args:      []string{"value"},
			shortcuts: map[rune]string{},
			errs:      []error{flags.ARGUMENT_NOT_NEED()},
			res:       map[string][]string{},
		},
		{
			name:      "same flags",
			args:      []string{"--flag", "--flag"},
			shortcuts: map[rune]string{},
			errs:      []error{flags.TWICE_FLAG()},
			res:       map[string][]string{"flag": {}},
		},
		{
			name:      "hasn't got a shortcut",
			args:      []string{"-f"},
			shortcuts: map[rune]string{},
			errs:      []error{flags.WRONG_SHORTCUT()},
			res:       map[string][]string{},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("#%d %s", i, tc.name), func(t *testing.T) {
			res, err := flags.ParseWithShortcuts(tc.args, tc.shortcuts)
			for _, e := range tc.errs {
				if !errors.Is(err, e) {
					t.Fatalf("got different errors expected %v, got %v", e, err)
				}
			}

			if !maps.EqualFunc(tc.res, res, func(v1, v2 []string) bool { return slices.Equal(v1, v2) }) {
				t.Fatalf("got different maps: expected %v, got %v", tc.res, res)
			}
		})
	}
}
