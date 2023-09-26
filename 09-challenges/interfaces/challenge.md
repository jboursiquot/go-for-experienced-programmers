# Challenge: Working with Interfaces

This challenge will have us implement a few interfaces to get more comfortable with them. 

We won't start from scratch though as we'll pull parts of the previous challenge's solution to give us a headstart and keep our focus on working with interfaces.

Open up the starting file located at `challenges/interfaces/begin/main.go` and look at the `doAnalysis` function signature:

`func doAnalysis(data string, counters ...counter) map[string]int`

This is a variadic function. It takes in some `string` data and a list of counters that must implement the `counter` interface. Inside the function, the `name()` method and `count()` methods of each counter is used to capture a key and value pair respectively.

The `letterCounter`, `numberCounter`, and `symbolCounter` types already predeclared here will need to support this.

You task is to ensure that these custom types implement the `counter` interface and do the character counting that is relevant to each. 

The `main` function already has some boilerplate for the file reading. You'll need to call on the `doAnalysis` function with the data and counters and dump the result as before.

The result of the analysis should be the same as when we implemented the solution for the previous challenge.

## Acceptance Criteria

1. Implement the `counter` interface for `letterCounter`, `numberCounter`, and `symbolCounter`.
2. Call on `doAnalysis` with data and counters to perform the analysis.
3. Dump the output and ensure it is the same as the sample output below.

## Sample Output

```
(map[string]int) (len=4) {
 (string) (len=5) "words": (int) 427,
 (string) (len=7) "letters": (int) 1784,
 (string) (len=7) "numbers": (int) 2,
 (string) (len=7) "symbols": (int) 510
}
```

Note that the maps have unordered keys by design. As long as your numbers match, you're good!