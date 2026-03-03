package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type CommandFlags struct {
	unique *bool
	random *bool
}

func parseCommandLine() (ans CommandFlags) {
	ans.unique = flag.Bool("unique", false, "return unique lines")
	flag.Parse()
	return
}

func main() {
	flags := parseCommandLine()

	var readers []io.Reader
	if flag.NArg() == 0 {
		readers = append(readers, os.Stdin)
	} else {
		for _, filename := range flag.Args() {
			fileHandle, err := os.Open(filename)
			if err != nil {
				panic(fmt.Sprintf("Cannot open file %s", filename))
			}
			readers = append(readers, fileHandle)
		}
	}

	var lines []string
	for _, reader := range readers {
		bytes, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}

		lines = append(lines, strings.Split(string(bytes), "\n")...)
	}

	seen := make(map[string]bool)

	slices.Sort(lines)
	for _, line := range lines {
		if flags.unique != nil && *flags.unique {
			if _, ok := seen[line]; ok {
				continue
			}
		}

		seen[line] = true
		fmt.Println(line)
	}
}
