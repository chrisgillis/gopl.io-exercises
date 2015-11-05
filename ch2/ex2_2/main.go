package main

import (
    "fmt"
    "flag"
    "os"
    "knd/ch2/tempconv"
    "strconv"
)

var convertTo = flag.String("conv", "", "c,f,feet,meters,lb,kg")

func main() {
    flag.Parse()

    if *convertTo == "" {
        fmt.Println("conversion type is required")
        os.Exit(1)
    }

    for _,v := range flag.Args() {
        convertedV,_ := strconv.ParseFloat(v,64)

        switch *convertTo {
            case "c":
                fmt.Println(tempconv.CToF(tempconv.Celsius(convertedV)))
            case "f":
                fmt.Println(tempconv.FToC(tempconv.Fahrenheit(convertedV)))
            case "feet":
                fmt.Println(convertedV/3.28084)
            case "meters":
                fmt.Println(convertedV*3.28084)
            case "lb":
                fmt.Println(convertedV*0.453592)
            case "kg":
                fmt.Println(convertedV/0.453592)
            default:
                fmt.Println("Invalid conversion type")
                os.Exit(1)
        }
    }
}
