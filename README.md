fix orthography in this readme

# FLags

Flags are an alternative for [flag standard package](https://pkg.go.dev/flag) and other libs.

## Flag format

Here are some rules for flags

### Normal flags

1. A flag **always** starts with "--"

Example:

```string
--flag
```

2. A flag can have values after

Example:

```string
--flag value
```

Flags can have multiple value

```string
--flag value1 value2 value3 ...
```

### Shortcuts

Flags can have shortcuts (as chars)

1. Shortcuts **always** starts with "-"

Example shortcut 'f' to flag "flag":

```string
-f
```

2. You can use multiple shortcuts

Example shortcut 'f' to flag "flag" and 'o' to flag "other_flag":

```string
-fo
```

3. Singular shortcuts have the same value rules like normal flags

Example shortcut 'f' to flag "flag":

```string
-f value
```

```string
-f value1 value2 value3 ...
```

4. With multiple shortcuts each shortcut will take only one value

Example shortcut 'f' to flag "flag" and 'o' to flag "other_flag":

```string
-fo value_for_flag value_for_other_flag 
```

But the last shortcut will use all value that are left

```string
-fo value_for_flag value_for_other_flag also_value_for_other_flag ...
```

## Converting to types rules

You can insert flags into a structure. here are some rules.

Convert order is:

### integers (int, int8...int64, uint, uint8...uint64, uintptr, time.Duration, unsafe.Pointer)

1. Converting to base 10

2. Converting form base 2 to base 16

3. for integers (int, int8...int64) converting to time.Duration

### boolean

1. Without args is true

2. Converting to bool (strconv.ParseBool)

### float (float32, float64)

Converting to float (strconv.ParseFloat)

### complex (complex64, complex128)

Converting to complex (strconv.ParseComplex)

### string

If there are brackets (", ', `) it trims them and gets the string

> [!WARNING]
>
> it won't convert the string without brackets

### array

It goes through the array and fills it with values (multiple values) for each value in the array it is using the same converting rule

> [!WARNING]
>
> for arrays and slices it won't work if there are 2d, 3d ... arrays/slices

### slice

It starts with the last elements in the slice and appends all values (multiple values) for each value it is using the same converting rule

### time

It will parse using all time formats. You can add your format

### interface

If it is an empty interface(`interface{}`) it will use default converting

1. string
2. int
3. float
4. bool
5. time
6. complex

### pointer

it will convert using the same convert rules into the type the pointer is linking to

### struct

it will do the same converting with the same flags for this struct

> [!WARNING]
>
> channels, maps, functions are not supported

## LICENSE

[MIT](license)