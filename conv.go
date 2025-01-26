package flags

import (
	"regexp"
	"strings"
)

func camelToSnake(s string) string {
	re := regexp.MustCompile("([A-Z])")
	snake := re.ReplaceAllString(s, "_$1")
	return strings.ToLower(strings.TrimPrefix(snake, "_"))
}
