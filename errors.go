package flags

var (
	// need a string
	TWICE_FLAG = Err("twice flag", "flag '%s' is used twice in the arguments")
	// need a char
	WRONG_SHORTCUT = Err("wrong shortcut", "shortcut '%c' does not exist")
	// need a string
	ARGUMENT_NOT_NEED = Err("argument not need", "argument '%s' isn't need")
	TYPE_ERROR        = Err("type error", "type of value is not a pointer or is nil")
	IS_NOT_A_STRUCT   = Err("is not a struct", "the value is not a structure")
	// need a value and a string
	CANT_CONVERT = Err("can't convert", "cant convert value '%v' to type %s")
	// need a string
	TOO_MANY_ARGUMENTS = Err("too many arguments", "flag %s has to many arguments")
	// need a string
	UNSUPPORTABLE_TYPE = Err("unsupportable type", "type %s isn't supported")
	// need a value
	CANT_DEFAULT_CONVERT = Err("cant default convert", "cant do default converting for val %v")
)
