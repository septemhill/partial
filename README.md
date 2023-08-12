# partial

The `partial` package provides a way make functions have more flexibility.

We still need type parameters to infer `FnArgX` currently,
but after [#58650](https://github.com/golang/go/issues/58650) be implemented, 
we could make it more clear.

EDIT: After testing on Go 1.21, we still need add type inference manually
please references to https://github.com/golang/go/issues/61948#issuecomment-1675196625 

```go
func Sum(a, b, c, d int) int {
    return a + b + c + d
}

func main() {
    r1 := partial.FnArg4[int, int, int, int, int](Sum).
        P1(1).
        P1(2).
        P1(3)(4)

    // After #58650 be implemented
    // r1 := partial.FnArg4(Sum).
    //     P1(1).
    //     P1(2).
    //     P1(3)(4)

    fmt.Println(r1) // The `r1` would be Sum(1,2,3,4) => 10

    r2 := partial.FnArg4[int, int, int, int, int](Sum).
        P2(1, 2)
        P1(3)(4)

    // After #58650 be implemented
    // r2 := partial.FnArg4(Sum).
    //     P2(1, 2).
    //     P1(3)(4)

    fmt.Println(r2) // The `r2` would be Sum(1,2,3,4) => 10

    r3 := partial.FnArg4[int, int, int, int, int](Sum)(1, 2, 3, 4)

    // After #58650 be implemented
    // r3 := partial.FnArg4(Sum)(1, 2, 3, 4)

    fmt.Println(r3) // The `r3` would be Sum(1,2,3,4) => 10
}
```
