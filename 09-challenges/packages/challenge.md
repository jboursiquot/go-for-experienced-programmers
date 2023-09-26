# Challenge: Working with Third-Party Packages

The first challenge is about learning to use third party packages and writing your own function.

You are to use the `github.com/jboursiquot/go-proverbs` package already part of the project's dependencies to retrieve and print a random proverb from it.

## Acceptance Criteria

1. You have a `main.go` file that imports `github.com/jboursiquot/go-proverbs`
2. You have a function (i.e. getRandomProverb) that returns a string and inside of it is uses the imported proverbs package to retrieve and return a random proverb's saying.
3. You make use of the standard library's `fmt` package to emit strings to `STDOUT`
4. You run `go mod tidy` to ensure you have the package in your local module cache.
5. You run your program with `go run` and get a random proverb with every run.
