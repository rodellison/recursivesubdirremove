package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/rodellison/recursivesubdirremove/app"
	"os"
)
// options is used to process command line arguments. These are minimal, in lieu of using
// environment variables to control the server.
type options struct {
	Debug bool `long:"debug" description:"Enable debug logging"`
}

func main() {

	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	var opts options
	args, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	// You can get individual args with normal indexing.
	argDir := args[0]
	argSubDirName := args[1]
	argConfirm := "n"

	if (len(args) == 3) {
		argConfirm = args[2]
	}

	fmt.Println(argDir)
	fmt.Println(argSubDirName)

	// Run the App
	err = app.Run(argDir, argSubDirName, argConfirm)
	if err != nil {
		// do stuff
		os.Exit(1)
	}



}


