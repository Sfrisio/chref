package main

import (
	"bufio"
	b "chref/build"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

var statRef, statDef syscall.Stat_t

var refPermissions fs.FileMode
var defIsDir fs.FileInfo

func main() {
	// RFILE: is the reference file, from which the user:group and permissions will be extracted
	// DFILE: is the destination file, to which the user:group membership and the permissions extracted from the reference file will be applied
	referencedObj := flag.String("reference", "", "referenced file or directory")
	destinationObj := flag.String("destination", "", "destination file or directory")
	implicitRecursive := flag.Bool("R", false, "implicit recursive if destination is a directory")
	var usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s --reference RFILE --destination DFILE [-R]\nchref Release: %s\nBuild Time: %s\nBuild User: %s\n", flag.CommandLine.Name(), b.Version, b.BuildTime, b.BuildUser)
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

			// Its a dir?
			if defIsDir.IsDir() {
				var answer bool
				// If true ask if apply to directory only or apply recursive
				if !*implicitRecursive {
					answer = askForYesOrNo("It seems a directory, do you want to change recursiverly?")
				} else {
					answer = true
				}
				if answer {
					// Use the filepath's Walk method and its WalkFunction to recursively walk through directories
					if errFilepathWalk := filepath.Walk(*destinationObj, printFullFilePath); errFilepathWalk != nil {
						fmt.Printf("%s: %s\n", os.Args[0], errFilepathWalk)
					}
				}
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

// TO-DO can be a good idea to have a generic funcion to check if file or directory exist without going into detail whether source or destination

func applyChown(destFile string, refUID, refGID int32) error {
	return syscall.Chown(destFile, int(refUID), int(refGID))
}

func applyChmod(destFile string, permissions fs.FileMode) error {
	return syscall.Chmod(destFile, uint32(permissions))
}

// askForYesOrNo as per name this function receive a "question" as an input string
// and retunr a boolean based on your answer y/yes(true) n/no(false)
// Suggestion: can be an idea to use a flag for an implicit yes?
func askForYesOrNo(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", question)

		useranswer, errReadString := reader.ReadString('\n')
		if errReadString != nil {
			log.Fatal(errReadString)
		}
		//lower and trim answer
		useranswer = strings.ToLower(strings.TrimSpace(useranswer))

		if useranswer == "y" || useranswer == "yes" {
			return true
		} else if useranswer == "n" || useranswer == "no" {
			return false
		}
	}
}

// printFullFilePath this function was passed as filepath.WalkFunc and is based on documentation examples
func printFullFilePath(dfilepath string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	// Here we can do a chown and chmod recursively :P instead to simply print the full file path
	if errChown := applyChown(dfilepath, int32(statRef.Uid), int32(statRef.Gid)); errChown != nil {
		fmt.Printf("%s: %s\n", flag.CommandLine.Name(), errChown)
		os.Exit(4)
	}
	if errChmod := applyChmod(dfilepath, refPermissions); errChmod != nil {
		fmt.Printf("%s: %s\n", flag.CommandLine.Name(), errChmod)
		os.Exit(4)
	}
	return nil
}
