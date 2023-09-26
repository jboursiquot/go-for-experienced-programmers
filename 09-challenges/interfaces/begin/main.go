// challenges/interfaces/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type counter interface {
	name() string
	count(input string) int
}

type letterCounter struct{ identifier string }

type numberCounter struct{ designation string }

type symbolCounter struct{ label string }

func doAnalysis(data string, counters ...counter) map[string]int {
	// initialize a map to store the counts
	analysis := make(map[string]int)

	// capture the length of the words in the data
	analysis["words"] = len(strings.Fields(data))

	// loop over the counters and use their name as the key
	for _, c := range counters {
		analysis[c.name()] = c.count(data)
	}

	// return the map
	return analysis
}

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	// use os package to read the file specified as a command line argument
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	data := string(bs)
	spew.Dump(data)

	// call doAnalysis and pass in the data and the counters

	// dump the map to the console using the spew package
	// spew.Dump(analysis)
}
