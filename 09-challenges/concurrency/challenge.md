# Challenge: Goroutine Synchronization

This challenge will reinforce our understanding of goroutine synchronization by expanding on a previous challenge.

We'll use a modified version of our solution to the interfaces challenge and introduce goroutines to have the counters do their work in a separate thread.

Open up `challenges/concurrency/begin/main.go`.

Things should look very familiar as most of the solution is unchanged. What we do have is a package-global `random` variable that we use an `init` function to initialized. `init()` functions in packages are always executed before any other function. This will ensure `random` is initialized at the start of the program so that we get non-deterministic numbers when we need to get randome numbers.

The next set of changes are to simulate delays in the counter types' count functions. Using the `random` variable from earlier, we're artificially creating a delay with the counting so we can get a better visual representation for what we want to do next.

That's where your task comes in. Your mission, should you choose to accept it 😜, is to perform the counting and assignment of the result of said counting in its own goroutine. This means in our `doAnalysis` function, as you loop through the counters, you must launch each one in its own goroutine.

Nothing else in the program needs to change but you'll need to answer the following acceptance criteria are met.

## Acceptance Criteria

1. You launch a goroutine for each counter that needs to perform a count.
2. You prevent goroutines from creating a data race by attempting to write to the `analysis` map at the same time.
3. You ensure that the `doAnalysis` function waits for every goroutine it launches.

## Sample Output

```
letters working...
symbols working...
numbers working...
(map[string]int) (len=4) {
 (string) (len=5) "words": (int) 427,
 (string) (len=7) "letters": (int) 1784,
 (string) (len=7) "symbols": (int) 510,
 (string) (len=7) "numbers": (int) 2
}
```