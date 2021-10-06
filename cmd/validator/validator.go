// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/openplaybooks/libcacao/objects/playbook"
	"github.com/pborman/getopt"
)

// These global variables hold build information. The Build variable will be
// populated by the Makefile and uses the Git Head hash as its identifier.
// These variables are used in the console output for --version and --help.
var (
	Version = "0.1.1"
	Build   string
)

// These global variables are for dealing with command line options
var (
	bOptDebug = getopt.BoolLong("debug", 0, "Debug")
	bOptHelp  = getopt.BoolLong("help", 0, "Help")
	bOptVer   = getopt.BoolLong("version", 0, "Version")
)

func main() {
	processCommandLineFlags()

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The validator needs JSON data to be passed to it via a pipe.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []byte

	for {
		input, err := reader.ReadByte()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	p, _ := playbook.Decode(output)

	valid, count, details := p.Valid()
	fmt.Println("Object valid: ", valid)
	fmt.Println("Error Count: ", count)

	// If the debug flag is set then print out detail about what is good and bad
	// otherwise just print out what is wrong.
	fmt.Println("\nDetails:")
	if *bOptDebug {
		for _, result := range details {
			fmt.Println(result)
		}
	} else {
		for _, result := range details {
			if result[0:2] != "++" {
				fmt.Println(result)
			}
		}
	}
}

// --------------------------------------------------
// Private functions
// --------------------------------------------------

// processCommandLineFlags - This function will process the command line flags
// and will print the version or help information as needed.
func processCommandLineFlags() {
	getopt.HelpColumn = 35
	getopt.DisplayWidth = 120
	getopt.SetParameters("")
	getopt.Parse()

	// Lets check to see if the version command line flag was given. If it is
	// lets print out the version infomration and exit.
	if *bOptVer {
		printOutputHeader()
		os.Exit(0)
	}

	// Lets check to see if the help command line flag was given. If it is lets
	// print out the help information and exit.
	if *bOptHelp {
		printOutputHeader()
		getopt.Usage()
		os.Exit(0)
	}
}

// printOutputHeader - This function will print a header for all console output
func printOutputHeader() {
	fmt.Println("")
	fmt.Println("CACAO Validator")
	fmt.Println("Copyright, Bret Jordan")
	fmt.Println("Version:", Version)
	if Build != "" {
		fmt.Println("Build:", Build)
	}
	fmt.Println("")
}
