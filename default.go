package flags

import (
	"strconv"
	"strings"
	"time"
)

var quotationMarks = []string{`"`, "'", "`"}

func convertString(s string) (string, bool) {
	for _, mark := range quotationMarks {
		if strings.HasPrefix(s, mark) && strings.HasSuffix(s, mark) {
			return strings.TrimPrefix(strings.TrimSuffix(s, mark), mark), true
		}
	}

	return "", false
}

func convertInt(s string, size int) (int64, bool) {
	if n, err := strconv.ParseInt(s, 10, size); err == nil {
		return n, true
	}
	for base := 2; base <= 16; base++ {
		if n, err := strconv.ParseInt(s, base, size); err == nil {
			return n, true
		}
	}

	if n, err := time.ParseDuration(s); err == nil {
		return int64(n), true
	}
	return 0, false
}

func convertUint(s string, size int) (uint64, bool) {
	if n, err := strconv.ParseUint(s, 10, size); err == nil {
		return n, true
	}
	for base := 2; base <= 16; base++ {
		if n, err := strconv.ParseUint(s, base, size); err == nil {
			return n, true
		}
	}
	return 0, false
}

func defaultConvert(s string) any {
	if s, ok := convertString(s); ok {
		return s
	}

	if n, ok := convertInt(s, 64); ok {
		return n
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	if b, err := strconv.ParseBool(s); err == nil {
		return b
	}

	if t, ok := parseTime(s); ok {
		return t
	}

	if c, err := strconv.ParseComplex(s, 128); err == nil {
		return c
	}

	return nil
}


