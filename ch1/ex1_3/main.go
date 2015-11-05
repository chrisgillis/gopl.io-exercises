// Measure difference between echo functions
// Note this is not a good evaluation of performance as it only
// does one run of each
package main

import (
	"fmt"
	"os"
	"time"
        "strings"
)

func main() {
    t := time.Now()
    echo1()
    fmt.Println(fmt.Sprintf("echo1 took %.8fs", time.Since(t).Seconds()));

    t = time.Now()
    echo2()
    fmt.Println(fmt.Sprintf("echo2 took %.8fs", time.Since(t).Seconds()))
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
    fmt.Println(strings.Join(os.Args[1:], " "));
}
