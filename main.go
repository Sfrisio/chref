package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"syscall"
)

func main() {
	// RFILE: is the reference file, from which the user:group and permissions will be extracted
	// DFILE: is the destination file, to which the user:group membership and the permissions extracted from the reference file will be applied
	referencedObj := flag.String("reference", "", "referenced file or directory")
	destinationObj := flag.String("destination", "", "destination file or directory")
	var statRef syscall.Stat_t
	var refPermissions fs.FileMode
	var usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s --reference RFILE --destination DFILE\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()
	// Checking starts from os.Args[1:] because the first value of this slice (os.Args[0]) is always the path to the program
	if len(os.Args[1:]) == 0 {
		usage()
	}
	// Flags return value is stored in pointer, so we have to use a pointer to evaluate the values ​​of the two strings
	if *referencedObj == "" || *destinationObj == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// Check if RFILE exists
	if _, errRefNotExists := os.Stat(*referencedObj); errors.Is(errRefNotExists, os.ErrNotExist) {
		fmt.Printf("%s: %s\n", os.Args[0], errRefNotExists)
		os.Exit(2)
	} else {
		// If RFILE exists, save its information in a struct (syscall.Stat_t)
		if errStatRef := syscall.Stat(*referencedObj, &statRef); errStatRef != nil {
			fmt.Printf("%s: %s\n", os.Args[0], errStatRef)
			os.Exit(3)
		} else {
			refPermissions = os.FileMode(statRef.Mode & 0777)
		}
	}
	// Check if DFILE exists
	if _, errDestNotExists := os.Stat(*destinationObj); errors.Is(errDestNotExists, os.ErrNotExist) {
		fmt.Printf("%s: %s\n", os.Args[0], errDestNotExists)
		os.Exit(2)
	}
	if errChown := applyChown(*destinationObj, int32(statRef.Uid), int32(statRef.Gid)); errChown != nil {
		fmt.Printf("%s: %s\n", os.Args[0], errChown)
		os.Exit(4)
	}
	if errChmod := applyChmod(*destinationObj, refPermissions); errChmod != nil {
		fmt.Printf("%s: %s\n", os.Args[0], errChmod)
	}
}

func applyChown(destFile string, refUID, refGID int32) error {
	return syscall.Chown(destFile, int(refUID), int(refGID))
}

func applyChmod(destFile string, permissions fs.FileMode) error {
	return syscall.Chmod(destFile, uint32(permissions))
}
