# Challenge: Flow Control

This challenge has us reading a local file to count the number of letters, numbers, and symbols in it. Along the way we get to make use of conditionals, loops, maps, and the standard library to help us with file-reading and working with character data.

## Acceptance Criteria

1. Use standard library `os` package to read the provided data file and panic if there's an error. BONUS: Use `os.Args` to get file name to read
2. Handle panics with a defered recovery.
3. Use a map to track count for words, letters, numbers, and symbols.
4. Use `for-range` loops
5. Use conditionals

## Sample Output

```
(map[string]int) (len=4) {
 (string) (len=7) "letters": (int) 1784,
 (string) (len=7) "symbols": (int) 84,
 (string) (len=7) "numbers": (int) 2,
 (string) (len=5) "words": (int) 427
}
```

> Note, use the `spew` library to dump out the map that's tracking your counts to get the output above.
