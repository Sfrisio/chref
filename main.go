package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"syscall"
)

const version = "v0.1"

func main() {
	// RFILE: is the reference file, from which the user:group and permissions will be extracted
	// DFILE: is the destination file, to which the user:group membership and the permissions extracted from the reference file will be applied
	referencedObj := flag.String("reference", "", "referenced file or directory")
	destinationObj := flag.String("destination", "", "destination file or directory")
	var statRef, statDef syscall.Stat_t

	var refPermissions fs.FileMode
	var defIsDir fs.FileInfo
	var usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s --reference RFILE --destination DFILE\nRelease %s\n", flag.CommandLine.Name(), version)
		//flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()
	// Checking starts from os.Args[1:] because the first value of this slice (os.Args[0]) is always the path to the program
	if len(os.Args[1:]) == 0 {
		usage()
	}
	// Flags return value is stored in pointer, so we have to use a pointer to evaluate the values ​​of the two strings
	if *referencedObj == "" || *destinationObj == "" {
		usage()
		os.Exit(1)
	}
	// Check if RFILE exists
	if _, errRefNotExists := os.Stat(*referencedObj); errors.Is(errRefNotExists, os.ErrNotExist) {
		fmt.Printf("%s: %s\n", flag.CommandLine.Name(), errRefNotExists)
		os.Exit(2)
	} else {
		// If RFILE exists, save its information in a struct (syscall.Stat_t)
		if errStatRef := syscall.Stat(*referencedObj, &statRef); errStatRef != nil {
			fmt.Printf("%s: %s\n", flag.CommandLine.Name(), errStatRef)
			os.Exit(3)
		} else {
			refPermissions = os.FileMode(statRef.Mode & 0777)
		}

	}
	// Check if DFILE exists
	if _, errDestNotExists := os.Stat(*destinationObj); errors.Is(errDestNotExists, os.ErrNotExist) {
		fmt.Printf("%s: %s\n", flag.CommandLine.Name(), errDestNotExists)
		os.Exit(2)
	} else {
		// If DFILE exists, save its information in a struct (syscall.Stat_t)
		if errStatDef := syscall.Stat(*destinationObj, &statDef); errStatDef != nil {
			fmt.Printf("%s: %s\n", flag.CommandLine.Name(), errStatDef)
			os.Exit(3)
		} else {
			// Check if DFILE is a file or directory
			defIsDir, _ = os.Stat(*destinationObj)
			// TO-DO error management
			// TO-DO move applyChown & applyChmod here

			// This is only for a temporary debug meanwhile implementing the recursive feature
			if 1 == 0 {
				fmt.Printf("Name: %s IsDir?: %t\n", defIsDir.Name(), defIsDir.IsDir())
			}
		}

	}

	if errChown := applyChown(*destinationObj, int32(statRef.Uid), int32(statRef.Gid)); errChown != nil {
		fmt.Printf("%s: %s\n", flag.CommandLine.Name(), errChown)
		os.Exit(4)
	}
	if errChmod := applyChmod(*destinationObj, refPermissions); errChmod != nil {
		fmt.Printf("%s: %s\n", flag.CommandLine.Name(), errChmod)
		os.Exit(4)
	}
}

// TO-DO for a recursive purpose we need to check id it's a file or directory
// we can use object like refPermissions.IsDir() but only for destination scope.
// Also can be a good idea to have a generic funcion to check if file or directory exist without going into detail whether source or destination

func applyChown(destFile string, refUID, refGID int32) error {
	return syscall.Chown(destFile, int(refUID), int(refGID))
}

func applyChmod(destFile string, permissions fs.FileMode) error {
	return syscall.Chmod(destFile, uint32(permissions))
}
