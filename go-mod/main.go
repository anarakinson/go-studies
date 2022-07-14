package main

import (
    "fmt"
    "gopackages/words"
    _ "gopackages/init"
    "github.com/fatih/color"
)

func main() {
    fmt.Println(words.Hello)
    for i := 0; i < 5; i++ {
        fmt.Println(words.Random())
    }
    color.Red("Red text")
}
