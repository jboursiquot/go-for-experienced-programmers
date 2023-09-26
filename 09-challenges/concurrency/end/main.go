package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type counter interface {
	name() string
	count(input string) int
}

type letterCounter struct{ identifier string }

func (l letterCounter) name() string {
	return l.identifier
}

func (l letterCounter) count(input string) int {
	result := 0
	fmt.Println(l.name(), "working...")
	time.Sleep(time.Duration(random.Intn(5)) * time.Second)
	for _, char := range input {
		if unicode.IsLetter(char) {
			result++
		}
	}
	return result
}

type numberCounter struct{ designation string }

func (n numberCounter) name() string {
	return n.designation
}

func (n numberCounter) count(input string) int {
	result := 0
	fmt.Println(n.name(), "working...")
	time.Sleep(time.Duration(random.Intn(5)) * time.Second)
	for _, char := range input {
		if unicode.IsNumber(char) {
			result++
		}
	}
	return result
}

type symbolCounter struct{ label string }

func (s symbolCounter) name() string {
	return s.label
}

func (s symbolCounter) count(input string) int {
	result := 0
	fmt.Println(s.name(), "working...")
	time.Sleep(time.Duration(random.Intn(5)) * time.Second)
	for _, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			result++
		}
	}
	return result
}

func doAnalysis(data string, counters ...counter) map[string]int {
	// initialize a map to store the counts
	analysis := make(map[string]int)

	// capture the length of the words in the data
	analysis["words"] = len(strings.Fields(data))

	var wg sync.WaitGroup
	wg.Add(len(counters))

	var mu sync.Mutex

	// loop over the counters and use their name as the key
	for _, c := range counters {
		go func(c counter) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			analysis[c.name()] = c.count(data)
		}(c)
	}

	wg.Wait()

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

	// call doAnalysis and pass in the data and the counters
	analysis := doAnalysis(data,
		letterCounter{identifier: "letters"},
		numberCounter{designation: "numbers"},
		symbolCounter{label: "symbols"},
	)

	// dump the map to the console using the spew package
	spew.Dump(analysis)
}
