// DO NOT EDIT THIS FILE.
//go:generate gb help documentation >alldocs.go

/*
gb, a project based build tool for the Go programming language.

Usage:

        gb command [arguments]

The commands are:

        build       build a package
        doc         show documentation for a package or symbol
        env         print project environment variables


Build a package

Usage:

        gb build [build flags] [packages]

Build compiles the packages named by the import paths, along with their dependencies.

The build flags are

	-f
		ignore cached packages if present, new packages built will overwrite any cached packages.
		This effectively disables incremental compilation.
	-F
		do not cache packages, cached packages will still be used for incremental compilation.
		-f -F is advised to disable the package caching system.
	-q
		decreases verbosity, effectively raising the output level to ERROR.
		In a successful build, no output will be displayed.
	-R
		sets the base of the project root search path from the current working directory to the value supplied.
		Effectively gb changes working directory to this path before searching for the project root.
	-v
		increases verbosity, effectively lowering the output level from INFO to DEBUG.
	-ldflags 'flag list'
		arguments to pass on each linker invocation.

The list flags accept a space-separated list of strings. To embed spaces in an element in the list, surround it with either single or double quotes.

For more about specifying packages, see 'gb help packages'. For more about where packages and binaries are installed, run 'gb help project'.


Show documentation for a package or symbol

Usage:

        gb doc <pkg> <sym>[.<method>]




Print project environment variables

Usage:

        gb env

Env prints project environment variables


Generate Go files by processing source




Run a plugin




Test a package




*/
package main
