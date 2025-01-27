package flags_test

import (
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
	"unsafe"

	"github.com/vandi37/flags"
)

type withSlice struct {
	Val   string `flag:"val"`
	Val2  int    `flag:"val2"`
	Slice []int  `flag:"slice"`
}

func TestWithSlice(t *testing.T) {
	f, err := flags.Parse(strings.Fields("--val 'be' --val2 37 --slice 37 42 418"), map[rune]string{})
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}

	val := new(withSlice)

	err = flags.Insert(f, val)
	if err != nil {
		t.Fatalf("got an error: %v", err)

	}

	need := withSlice{"be", 37, []int{37, 42, 418}}

	if val.Val != need.Val || need.Val2 != val.Val2 || !slices.Equal(need.Slice, val.Slice) {
		t.Fatalf("got structure %+v, expected %+v", val, need)
	}
}

type allTypes struct {
	Bool          bool
	Int           int
	Int8          int8
	Int16         int16
	Int32         int32
	Int64         int64
	Uint          uint
	Uint8         uint8
	Uint16        uint16
	Uint32        uint32
	Uint64        uint64
	Uintptr       uintptr
	Float32       float32
	Float64       float64
	Complex64     complex64
	Complex128    complex128
	Array         [5]int      // Array of 5 integers
	Interface     interface{} // A generic interface
	Slice         []string    // A slice of strings
	String        string
	Pointer       *int
	UnsafePointer unsafe.Pointer
}

func TestAll(t *testing.T) {
	var n = 42
	ptr := unsafe.Pointer(&n)
	f, err := flags.Parse(strings.Fields("--pointer 37 --bool --int 10 --int8 8 --int16 16 --int32 32 --int64 64 --uint 10 --uint8 8 --uint16 16 --uint32 32 --uint64 64 --uintptr 10 --float32 32.320000 --float64 64.640000 --complex64 (3+4i) --complex128 (5+6i) --array 1 2 3 4 5 --interface 'hello' --slice 'a' 'b' 'c' --string 'test' --unsafe_pointer "+strconv.Itoa(int(uintptr(ptr)))), map[rune]string{})
	if err != nil {
		t.Fatalf("got an error: %v", err)

	}

	at := new(allTypes)

	err = flags.Insert(f, at)
	if err != nil {
		t.Fatalf("got an error: %v", err)

	}

	var i int = 37
	other := allTypes{
		Bool:          true,
		Int:           10,
		Int8:          8,
		Int16:         16,
		Int32:         32,
		Int64:         64,
		Uint:          10,
		Uint8:         8,
		Uint16:        16,
		Uint32:        32,
		Uint64:        64,
		Uintptr:       10,
		Float32:       32.32,
		Float64:       64.64,
		Complex64:     complex(3, 4),
		Complex128:    complex(5, 6),
		Array:         [5]int{1, 2, 3, 4, 5},
		Interface:     "hello",
		Slice:         []string{"a", "b", "c"},
		String:        "test",
		Pointer:       &i,
		UnsafePointer: ptr,
	}

	if at.Bool != other.Bool ||
		at.Int != other.Int ||
		at.Int8 != other.Int8 ||
		at.Int16 != other.Int16 ||
		at.Int32 != other.Int32 ||
		at.Int64 != other.Int64 ||
		at.Uint != other.Uint ||
		at.Uint8 != other.Uint8 ||
		at.Uint16 != other.Uint16 ||
		at.Uint32 != other.Uint32 ||
		at.Uint64 != other.Uint64 ||
		at.Uintptr != other.Uintptr ||
		at.Float32 != other.Float32 ||
		at.Float64 != other.Float64 ||
		at.Complex64 != other.Complex64 ||
		at.Complex128 != other.Complex128 ||
		at.Array != other.Array ||
		!reflect.DeepEqual(at.Interface, other.Interface) ||
		!slices.Equal(at.Slice, other.Slice) ||
		at.String != other.String ||
		at.UnsafePointer != other.UnsafePointer ||
		*at.Pointer != *other.Pointer ||
		*(*int)(at.UnsafePointer) != n {
		t.Fatalf("got structure %+v, expected %+v", at, other)
	}

}

type base struct {
	Num int
}

type inside struct {
	Base base
	Str  string
}

func TestInside(t *testing.T) {
	val := new(inside)
	err := flags.Load(strings.Fields("--num 37 --str 'be'"), val, map[rune]string{})
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}

	need := inside{
		Str:  "be",
		Base: base{Num: 37},
	}
	if need.Base.Num != val.Base.Num || need.Str != val.Str {
		t.Fatalf("got structure %+v, expected %+v", val, need)
	}
}
