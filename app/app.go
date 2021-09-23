package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")
var subDirToDelete = ""
var confirmDelete = ""

func Run(dirName, subdirToRemove, confirm string) error {
	// Application runtime code goes here
	subDirToDelete = subdirToRemove
	confirmDelete = confirm
	err := filepath.Walk(dirName, visit)
	if err!= nil {
		fmt.Println(err)
	}
	return nil
}


// `visit` is called for every file or directory found
// recursively by `filepath.Walk`.
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		var message = err.Error()
		if (strings.Contains(message, "cannot find the path specified")) {
			return nil
		} else {
			return err
		}
	}

	//P is the directory path (and file)
	var tempSubdirString = strings.TrimPrefix(info.Name(), p)
	if (info.IsDir() && tempSubdirString == subDirToDelete) {
		if (confirmDelete == "y") {
			//perform the removeall to remove the directory and any files contained within
			fmt.Println("Removing directory: ", p)
			os.RemoveAll(p)
		} else
		{
			fmt.Println("Found target directory: ", p)
		}
	}

	return nil
}


func Shutdown() {
	// Shutdown contexts, listeners, and such
}

