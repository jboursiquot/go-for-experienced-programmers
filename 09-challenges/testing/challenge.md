# Challenge: Table-Driven Tests

This challenge will have you write table-driven tests for the counters we used in the challenge on interfaces, mainly the letter counter, the number counter, and the symbol counter.

For your convenience, I've put the counters and their count methods all in the challenge's `begin` directory's `main.go` file. The `main_test.go`, also in that directory, is where you'll write your tests.

## Acceptance Criteria

1. You have `TestLetterCount`, `TestNumberCount`, and `TestSymbolCount` functions in your test file.
2. Each of the tests uses an appropriate set of test cases.
3. You make use of subtests to test each test case in your functions.

### Sample Test Output

```
=== RUN   TestLetterCount
=== RUN   TestLetterCount/#00
=== RUN   TestLetterCount/a2_32_^_&/)
=== RUN   TestLetterCount/812_%6ab//
--- PASS: TestLetterCount (0.00s)
    --- PASS: TestLetterCount/#00 (0.00s)
    --- PASS: TestLetterCount/a2_32_^_&/) (0.00s)
    --- PASS: TestLetterCount/812_%6ab// (0.00s)
=== RUN   TestNumberCount
=== RUN   TestNumberCount/#00
=== RUN   TestNumberCount/abc_1,?/
=== RUN   TestNumberCount/abc_12&8_^
--- PASS: TestNumberCount (0.00s)
    --- PASS: TestNumberCount/#00 (0.00s)
    --- PASS: TestNumberCount/abc_1,?/ (0.00s)
    --- PASS: TestNumberCount/abc_12&8_^ (0.00s)
=== RUN   TestSymbolCount
=== RUN   TestSymbolCount/#00
=== RUN   TestSymbolCount/abc_1,?/
=== RUN   TestSymbolCount/abc_12&8_^_
--- PASS: TestSymbolCount (0.00s)
    --- PASS: TestSymbolCount/#00 (0.00s)
    --- PASS: TestSymbolCount/abc_1,?/ (0.00s)
    --- PASS: TestSymbolCount/abc_12&8_^_ (0.00s)
PASS
coverage: 100.0% of statements
```
