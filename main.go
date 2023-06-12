package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const hyphenSign = "-"

var help *bool = flag.Bool("h", false, "help")

func main() {

	flag.Parse()

	if *help {
		exitWithUsage(0)
	}

	var (
		in        string
		out       string
		inputData []byte
	)

	if len(os.Args) == 1 {
		printErr(fmt.Errorf("no args supplied"))
	}

	if len(os.Args) > 1 {
		in = os.Args[1]
	}

	if len(os.Args) > 2 {
		out = os.Args[2]
	}

	if in == hyphenSign {
		// read from stdin
		bs, err := io.ReadAll(os.Stdin)
		if err != nil {
			printErr(err)
		}
		inputData = bs
	} else {
		// read from file
		bs, err := os.ReadFile(in)
		if err != nil {
			printErr(err)
		}
		inputData = bs
	}

	outputData, err := Uncompress(inputData)
	if err != nil {
		printErr(err)
	}

	if out == hyphenSign || out == "" {
		// write to stdout
		fmt.Fprint(os.Stdout, string(outputData))
	} else {
		// write to file
		err := os.WriteFile(out, outputData, 0644)
		if err != nil {
			printErr(err)
		}
	}
}

func printErr(err error) {
	fmt.Fprintf(os.Stderr, "\nERR: %s\n\n", err)
	exitWithUsage(-1)
}

const usage string = `Usage: go-dejsonlz4 [-h] IN_FILE [OUT_FILE]
Example: go-dejsonlz4 ~/.mozilla/firefox/aks8v8c0.default-release/bookmarkbackups/bookmarks-2023-03-01_1011_OItiw5WByHsdl6u-lQ08mQ==.jsonlz4
Decompress Firefox bookmark files with .jsonlz4 extension from IN_FILE to OUT_FILE:
	* -h, display help message and exit,
	* IN_FILE='-', uncompress from standard input,
   	* OUT_FILE='-' or missing, uncompress to standard output.
`

func exitWithUsage(code int) {
	var w io.Writer = os.Stdout
	if code != 0 {
		w = os.Stderr
	}
	fmt.Fprint(w, usage)
	os.Exit(code)
}
