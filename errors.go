package flags

var (
	// need a string
	TWICE_FLAG = err("twice flag", "flag '%s' is used twice in the arguments")
	// need a char
	WRONG_SHORTCUT = err("wrong shortcut", "shortcut '%c' does not exist")
	// need a string
	ARGUMENT_NOT_NEED = err("argument not need", "argument '%s' isn't need")
	TYPE_ERROR        = err("type error", "type of value is not a pointer or is nil")
	IS_NOT_A_STRUCT   = err("is not a struct", "the value is not a structure")
	// need a value and a string
	CANT_CONVERT = err("can't convert", "cant convert value '%v' to type %s")
	// need a string
	TOO_MANY_ARGUMENTS = err("too many arguments", "flag %s has to many arguments")
	// need a string
	UNSUPPORTABLE_TYPE = err("unsupportable type", "type %s isn't supported")
	// need a value
	CANT_DEFAULT_CONVERT = err("cant default convert", "cant do default converting for val %v")
)
