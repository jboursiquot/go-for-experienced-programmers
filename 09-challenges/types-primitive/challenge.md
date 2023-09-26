# Challenge: Working with Primitive Types

This challenge will help reinforce your knowledge of primitive types, scoping rules, and getting more comfortable with pointers in the process.

## Acceptance Criteria

Follow each step below to produce output similar to the sample.

1. Declare a package-level variable `val` with inferred type `string` with an assigned value
2. In `main`, use short variable declaration for a local variable `val` and assign it an integer
3. Print out the type and value of the local `val`
4. Write a function to print out the type and value of the global `val`
5. Write a function to update the global `val`
6. Print out the type and value of a dereference local `val`
7. Assign a value to the local `val` directly to the underlying memory location you  - you will need to use a pointer dereference for this.
8. Print out the local `val` to show you managed to update the underlying memory location.

## Sample expected output:

```
int, local val = 42
global val =  global
global val =  updated global
*int, local &val = 0xc00001a0a8
local val =  100
```
