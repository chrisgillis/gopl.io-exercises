package main

import "fmt"

// Package level declaration, not exported as does not begin with upper case
const boilingF = 212.0

func main() {
    // Local declaration
    var f = boilingF
    var c = fToC(f) 
    fmt.Printf("boiling point = %gdegF or %gdegC\n", f, c)

    p := &f; // p, of type pointer-to-float64, points to f
    fmt.Println(p) // Print the address p points to
    fmt.Println(*p) // Print the value p points to
    fmt.Println(&p) // Print the address of p
}

func fToC(f float64) float64 {
    return (f - 32) * (5/9)
}

