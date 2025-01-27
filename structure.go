package flags

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func Insert(flags map[string][]string, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return TYPE_ERROR()
	}

	rv = rv.Elem()
	rt := rv.Type()

	return insert(flags, rv, rt)
}

func insert(flags map[string][]string, v reflect.Value, t reflect.Type) error {
	if v.Kind() != reflect.Struct {
		return IS_NOT_A_STRUCT()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		fieldName := fieldType.Tag.Get("flag")
		if fieldName == "" {
			fieldName = camelToSnake(fieldType.Name)
		}

		if field.Kind() == reflect.Struct {
			if err := insert(flags, field, fieldType.Type); err != nil {
				return err
			}
			continue
		}

		args, ok := flags[fieldName]
		if !ok || !field.CanSet() || args == nil {
			continue
		}

		if err := setValue(args, field, fieldName); err != nil {
			return err
		}

	}
	return nil
}

func setValue(args []string, field reflect.Value, fieldName string) error {
	switch field.Kind() {
	case reflect.Bool:
		if len(args) <= 0 {
			field.SetBool(true)
			return nil
		}
		if len(args) != 1 {
			return TOO_MANY_ARGUMENTS(fieldName)
		}

		if b, err := strconv.ParseBool(args[0]); err != nil {
			return CANT_CONVERT(args[0], "bool")
		} else {
			field.SetBool(b)
		}
	case reflect.Int:
		if err := setInt(field, fieldName, args, 0); err != nil {
			return err
		}
	case reflect.Int8:
		if err := setInt(field, fieldName, args, 8); err != nil {
			return err
		}
	case reflect.Int16:
		if err := setInt(field, fieldName, args, 16); err != nil {
			return err
		}
	case reflect.Int32:
		if err := setInt(field, fieldName, args, 32); err != nil {
			return err
		}
	case reflect.Int64:
		if err := setInt(field, fieldName, args, 64); err != nil {
			return err
		}
	case reflect.Uint:
		if err := setUint(field, fieldName, args, 0); err != nil {
			return err
		}
	case reflect.Uint8:
		if err := setUint(field, fieldName, args, 8); err != nil {
			return err
		}
	case reflect.Uint16:
		if err := setUint(field, fieldName, args, 16); err != nil {
			return err
		}
	case reflect.Uint32:
		if err := setUint(field, fieldName, args, 32); err != nil {
			return err
		}
	case reflect.Uint64:
		if err := setUint(field, fieldName, args, 64); err != nil {
			return err
		}
	case reflect.Uintptr:
		if err := setUint(field, fieldName, args, 64); err != nil {
			return err
		}
	case reflect.Float32:
		if len(args) != 1 {
			return TOO_MANY_ARGUMENTS(fieldName)
		}

		if f, err := strconv.ParseFloat(args[0], 32); err != nil {
			return CANT_CONVERT(args[0], "float32")
		} else {
			field.SetFloat(f)
		}
	case reflect.Float64:
		if len(args) != 1 {
			return TOO_MANY_ARGUMENTS(fieldName)
		}

		if f, err := strconv.ParseFloat(args[0], 64); err != nil {
			return CANT_CONVERT(args[0], "float64")
		} else {
			field.SetFloat(f)
		}
	case reflect.Complex64:
		if len(args) != 1 {
			return TOO_MANY_ARGUMENTS(fieldName)
		}

		if c, err := strconv.ParseComplex(args[0], 64); err != nil {
			return CANT_CONVERT(args[0], "complex64")
		} else {
			field.SetComplex(c)
		}
	case reflect.Complex128:
		if len(args) != 1 {
			return TOO_MANY_ARGUMENTS(fieldName)
		}

		if c, err := strconv.ParseComplex(args[0], 128); err != nil {
			return CANT_CONVERT(args[0], "complex128")
		} else {
			field.SetComplex(c)
		}
	case reflect.UnsafePointer:
		if len(args) != 1 {
			return TOO_MANY_ARGUMENTS(fieldName)
		}

		if n, ok := convertUint(args[0], 64); !ok {
			return CANT_CONVERT(args[0], "unsafe.Pointer")
		} else {
			field.SetPointer(unsafe.Pointer(uintptr(n)))
		}
	case reflect.String:
		if len(args) != 1 {
			return TOO_MANY_ARGUMENTS(fieldName)
		}

		if s, ok := convertString(args[0]); !ok {
			return CANT_CONVERT(args[0], "string")
		} else {
			field.SetString(s)
		}
	case reflect.Array:
		if field.Len() != len(args) {
			return TOO_MANY_ARGUMENTS(fieldName)
		}

		for i := 0; i < field.Len(); i++ {
			if err := setValue([]string{args[i]}, field.Index(i), strconv.Itoa(i)); err != nil {
				return err
			}
		}

	case reflect.Slice:
		filedLen := field.Len()
		if len(args) > field.Cap()-filedLen {
			newSlice := reflect.MakeSlice(field.Type(), filedLen+len(args), filedLen+len(args))
			reflect.Copy(newSlice, field)
			field.Set(newSlice)
		}

		for i := 0; i < len(args); i++ {
			if err := setValue([]string{args[i]}, field.Index(filedLen+i), strconv.Itoa(filedLen+i)); err != nil {
				return err
			}
		}
	case reflect.Interface:
		if field.NumMethod() != 0 {
			return UNSUPPORTABLE_TYPE(field.Kind().String())
		}

		if len(args) == 0 {
			field.Set(reflect.ValueOf(true))
		}
		var vals = []any{}
		for _, arg := range args {
			var conv = defaultConvert(arg)
			if conv == nil {
				return CANT_DEFAULT_CONVERT(arg)
			}
			vals = append(vals, conv)
		}
		var res any
		if len(vals) == 1 {
			res = vals[0]
		} else {
			res = vals
		}
		field.Set(reflect.ValueOf(res))
	case reflect.Pointer:
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		return setValue(args, field.Elem(), fieldName)
	default:
		return UNSUPPORTABLE_TYPE(field.Kind().String())

	}
	return nil
}

func setInt(val reflect.Value, fieldName string, args []string, size int) error {
	if len(args) != 1 {
		return TOO_MANY_ARGUMENTS(fieldName)
	}

	if n, ok := convertInt(args[0], size); !ok {
		return CANT_CONVERT(args[0], fmt.Sprint("int", size))
	} else {
		val.SetInt(n)
	}

	return nil
}

func setUint(val reflect.Value, fieldName string, args []string, size int) error {
	if len(args) != 1 {
		return TOO_MANY_ARGUMENTS(fieldName)
	}

	if n, ok := convertUint(args[0], size); !ok {
		return CANT_CONVERT(args[0], fmt.Sprint("uint", size))
	} else {
		val.SetUint(n)
	}

	return nil
}
