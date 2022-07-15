package main

import (
    "fmt"
    "gopackages/words"
    _ "gopackages/init"
    "github.com/fatih/color"
    "github.com/anarakinson/go_test_module" // package name - <utils> - is used below, not repository name - <go_test_module>
    utilsv2 "github.com/anarakinson/go_test_module/v2" // package from other github branch - v2
    utilsv3 "github.com/anarakinson/go_test_module/v3" // package from other github branch - v3
)

func main() {
    fmt.Println(words.Hello)
    for i := 0; i < 5; i++ {
        fmt.Println(words.Random())
    }
    color.Red("Red text")

    fmt.Println("---")

    var target = "three"
    var slc = []string{"zero", "one", "two", "three", "four"}
    fmt.Println("Searching for word:", target, "in slice:", slc)
    res := utils.Contains(slc, target)
    fmt.Println("Target is in slice:", res)

    fmt.Println("---")

    res = utilsv3.InSlice(slc, target)
    fmt.Println("Target is in slice(v3):", res)

    fmt.Println("---")

    var targetI = 3
    var slcI = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Println("Searching for word:", targetI, "in slice:", slcI)
    resI := utilsv2.InSliceInt(slcI, targetI)
    fmt.Println("Target is in slice:", resI)
}
