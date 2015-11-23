package main

import (
    "fmt"
    "os"
    "sort"
    "strings"
)

func main() {
    if (len(os.Args) < 3) {
        fmt.Println("Usage: ex3_12 string1 string2")
    }

    var isAnagram = isAnagram(os.Args[1], os.Args[2]);

    if (isAnagram) {
        fmt.Println("Yes, they are anagrams.")
    } else {
        fmt.Println("No, they are not anagrams.")
    }
}

func isAnagram(a string, b string) bool {
    var strArrayA, strArrayB []string

    for _,v := range a {
       strArrayA = append(strArrayA, string(v))
    }

    for _,v := range b {
        strArrayB = append(strArrayB, string(v))
    }

    sort.Strings(strArrayA)
    sort.Strings(strArrayB)

    a = strings.Join(strArrayA, "")
    b = strings.Join(strArrayB, "")

    return a == b
}
