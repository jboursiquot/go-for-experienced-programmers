package main

import (
	"fmt"

	// import the proverbs package
	"github.com/jboursiquot/go-proverbs"
)

// getRandomProverb returns a random proverb from the proverbs package
func getRandomProverb() string {
	return proverbs.Random().Saying
}

func main() {
	// print the result of calling your getRandomProverb function
	fmt.Println(getRandomProverb())
}
