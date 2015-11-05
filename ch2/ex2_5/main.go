package main

import (
    "fmt"
)
var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func main() {
    for k,v := range pc {
        fmt.Printf("%v: %v\n",k,v)
    }

    fmt.Println("PopCount()")
    fmt.Println(PopCount(5))
    fmt.Println(PopCount(255))

    fmt.Println("PopCountLoop()") 
    fmt.Println(PopCountLoop(5))
    fmt.Println(PopCountLoop(255))

    fmt.Println("PopCount64()")
    fmt.Println(PopCount64(5))
    fmt.Println(PopCount64(255))
}

func PopCount(x uint64) int {
    return int(pc[byte(x >> (0*8))] +
               pc[byte(x >> (1*8))] +
               pc[byte(x >> (2*8))] +
               pc[byte(x >> (3*8))] +
               pc[byte(x >> (4*8))] +
               pc[byte(x >> (5*8))] +
               pc[byte(x >> (6*8))] +
               pc[byte(x >> (7*8))])
}

func PopCountLoop(x uint64) int {
    setBits := 0
    var i uint
    for i = 0; i < 8; i++ {
        setBits += int(pc[byte(x >> (i*8))])
    }
    return setBits
}

func PopCount64(x uint64) int {
    setBits := 0
    for ; x != 0; {
        x = x & (x - 1)
        setBits++
    }
    return setBits
}
