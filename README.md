Please convert this readme to go doc comment format
# FLags

Flags are an alternative to the [flag standard package](https://pkg.go.dev/flag).

## Flag Format

Here are some rules for flags

### Normal Flags

1. A flag **always** starts with "--"

Example:

```string
--flag
```

2. A flag may have values

Example:

```string
--flag value
```

Flags can have multiple values. 

```string
--flag value1 value2 value3 ...
```

### Shortcut Flags

Shortcut always starts with.

1. Shortcuts **always** starts with "-"

 Example, shortcut 'f' for flag "flag": 

```string
-f
```

2. You can have multiple shortcuts. 

Example shortcut 'f' for flag "flag" and 'o' to flag "other_flag":

```string
-fo
```

3. Singular shortcut has the same rules as normal flags.

Example shortcut 'f' for flag "flag":

```string
-f value
```

```string
-f value1 value2 value3 ...
```

4. Each shortcut takes only one value. 

Example shortcut 'f' to flag "flag" and 'o' to flag "other_flag":

```string
-fo value_for_flag value_for_other_flag 
```

With multiple shortcuts, the last one will use any remaining values.

```string
-fo value_for_flag value_for_other_flag also_value_for_other_flag ...
```

## Converting Flags to Types

Flags may be inserted into a structure, here are the rules:

### Integers (int, int8...int64, uint, uint8...uint64, uintptr, time.Duration, unsafe.Pointer)

1. Convert to base 10.

2. Convert from base 2 (binary) to base 16 (hex).

3. For integers (int, int8...int64) convert to time.Duration

### Boolean

1. Without arguments, it's true.

2. Convert to bool using strconv.ParseBool

### float (float32, float64)

Convert to float using strconv.ParseFloat

### Complex (complex64, complex128)

Convert to complex using strconv.ParseComplex

### String

If there are brackets (", ', `) , trim them and get the string

> [!WARNING]
>
> it won't convert the string without brackets

### array

It goes through the array and fills it with multiple values for each value. It uses the same conversion rule for all values in the array.

> [!WARNING]
>
> for arrays and slices it won't work if there are 2d, 3d ... arrays/slices

### slice

It starts with the last element in the slice, and appends all the values (multiple values), for each value, using the same conversion rule.

### time

Parse using all time formats. You can specify your own formats. 

### interface

If it's an empty interface (interface{}), use default conversion

1. string
2. int
3. float
4. bool
5. time
6. complex

### pointer

Convert using same conversion rules to type pointed to by pointer

### struct

Do same conversion with same flags on this struct.

> [!WARNING]
>
> channels, maps, functions are not supported

## LICENSE

[MIT](license)