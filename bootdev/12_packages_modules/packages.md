# Notes

Typically the package name (ex. `package packagesmodules`) aligns with the
directory name.

The name isn't strict, but all `.go` files in the folder must have the same
starting line. Otherwise, there will be an error?

The `package main` is compiled as an executable.
All other package names (not `main`) are library packages (no entry point)

You can actually look at the go source code and locate the `math/rand` package
by following the directory.

# MODULES

All source files within the same folder have access to each other (no import needed).
