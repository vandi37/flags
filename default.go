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

var timeFormats = []string{
	time.Layout,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	time.DateTime,
	time.DateOnly,
	time.TimeOnly,
}

func parseTime(s string) (time.Time, bool) {
	for _, format := range timeFormats {
		if t, err := time.Parse(format, s); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}
