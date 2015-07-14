package main

import (
	"fmt"
	"go/build"
	"os"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	// Loader is a tool for opening Go files, it loads from a Config type
	conf := loader.Config{
		Build: &build.Default,
	}

	file, err := conf.ParseFile("looper.go", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Creation of a single file main pacakge
	conf.CreateFromFiles("looper", file)
	conf.Import("runtime")

	p, err := conf.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Finally, create SSA representation from the package we've loaded
	program := ssautil.CreateProgram(p, ssa.SanityCheckFunctions)
	looperPkg := program.Package(p.Created[0].Pkg)

	fmt.Println("RIGHT IN THE SINGLE STATIC ASSIGNMENT FORM:")
	looperPkg.WriteTo(os.Stdout)

	fmt.Println("LOOK AT THIS HERE LOOPER FUNC:")
	looperFunc := looperPkg.Func("Looper")
	looperFunc.WriteTo(os.Stdout)
}
