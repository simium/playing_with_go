package main

import "fmt"

func fib(f int) int {
    if f == 0 || f == 1 {
        return f
    } else {
        return fib(f-1) + fib(f-2)
    }
}

func main() {
    x := fib(40)
    fmt.Println(x)
}
