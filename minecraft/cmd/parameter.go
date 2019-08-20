package cmd

import "reflect"

// Parameter is an interface for a generic parameters. Users may have types as command parameters that
// implement this parameter.
type Parameter interface {
	// Parse takes an arbitrary amount of arguments from the command Line passed and parses it, so that it can
	// store it to value v. If the arguments cannot be parsed from the Line, an error should be returned.
	Parse(line *Line, v reflect.Value) error
	// Type returns the type of the parameter. It will show up in the usage of the command, and, if one of the
	// known type names, will also show up client-side.
	Type() string
}

// Enum is an interface for enum-type parameters. Users may have types as command parameters that implement
// this parameter in order to allow a specific set of options only.
type Enum interface {
	// Type returns the type of the enum. This type shows up client-side in the command usage, in the spot
	// where parameter types otherwise are.
	Type() string
	// Options should return a list of options that show up on the client side. The command will ensure that
	// the argument passed to the enum parameter will be equal to one of these options.
	Options() []string
	// SetOption sets the option that was selected in the argument. It is guaranteed to be one of the options
	// obtained using Enum.Options().
	SetOption(option string, v reflect.Value)
}