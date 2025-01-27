package flags

import (
	"reflect"
	"slices"
	"time"
)

// it appends a new time format string to the list of accepted time formats.
//
// The provided format string will be added to the end of the existing list of
// formats.  This allows the application to recognize additional, user-defined
// or custom time formats.
func AddTimeFormat(format string) {
	timeFormats = append(timeFormats, format)
}

// it returns a copy of the current list of accepted time format strings.
//
// The returned slice is a deep copy, meaning that modifications made to the
// returned slice will not affect the original time format list maintained
// internally. This ensures that the internal state is protected from accidental
// or intentional changes from external code.
func GetTimeFormats() []string {
	return slices.Clone(timeFormats)
}

// TODO: Adding removing or other better working with time formats, cause now it is bad(
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
	if c, ok := convertString(s); ok {
		s = c
	}
	for _, format := range timeFormats {
		if t, err := time.Parse(format, s); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

func setTime(val reflect.Value, args []string, fieldName string) error {
	if len(args) != 1 {
		return TOO_MANY_ARGUMENTS(fieldName)
	}

	if t, ok := parseTime(args[0]); ok {
		val.Set(reflect.ValueOf(t))
	} else {
		return CANT_CONVERT(args[0], "time.Time")
	}

	return nil
}
