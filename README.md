# Go Partial

A small, generic-based Go library for partial function application.

This library provides a way to "fix" a number of arguments to a function, producing another function of smaller arity. This is achieved using Go 1.18+ generics.

## Features

- Supports functions with up to 5 arguments.
- Type-safe partial application.
- Simple and intuitive API.

## Usage

To use the library, you first need to wrap your function using the appropriate `FnArgX` wrapper, where `X` is the number of arguments your function takes.

Once wrapped, you can use the `P` methods (`P1`, `P2`, etc.) to apply arguments partially.

- `P1(arg1)`: Applies the first argument.
- `P2(arg1, arg2)`: Applies the first two arguments.
- ...and so on.

Each `P` method returns a new function that takes the remaining arguments.

### Example

Let's say you have a function that adds three integers:

```go
package main

import (
	"fmt"
	"your_module/partial" // Make sure to use your actual module path
)

func add(a, b, c int) int {
	return a + b + c
}

func main() {
	// 1. Wrap the original function to get a partial-ready function.
	partialAdd := partial.FnArg3(add)

	// 2. Apply the first argument.
	// P1 returns a function that takes the remaining 2 arguments (b, c).
	add10 := partialAdd.P1(10)
	result1 := add10(20, 5) // 10 + 20 + 5 = 35
	fmt.Println("Result 1:", result1)

	// 3. Apply the first two arguments.
	// P2 returns a function that takes the remaining 1 argument (c).
	add10and20 := partialAdd.P2(10, 20)
	result2 := add10and20(5) // 10 + 20 + 5 = 35
	fmt.Println("Result 2:", result2)

	// 4. Apply all arguments.
	// P3 returns a function that takes no arguments.
	add10and20and5 := partialAdd.P3(10, 20, 5)
	result3 := add10and20and5() // 10 + 20 + 5 = 35
	fmt.Println("Result 3:", result3)

	// 5. You can also chain partial applications directly.
	// Each `P` method returns a new partial function object, allowing for fluent chaining.
	result4 := partialAdd.P1(10).P1(20)(5) // 10 + 20 + 5 = 35
	fmt.Println("Result 4 (chained):", result4)
}
```

## API

The library provides wrappers and methods for functions with 0 to 5 arguments.

- `FnArg0`, `FnArg1`, `FnArg2`, `FnArg3`, `FnArg4`, `FnArg5`
- Methods `P1`, `P2`, `P3`, `P4`, `P5` for partial application on the corresponding function types.