package main

import (
    "fmt"
    "internal/customer"
)

func main() {
    // cust1 := customer.Customer{
    //     Name: "Billy",
    //     Age: 20,
    //     Balance: 10000,
    //     Debt: 1000,
    // }

    cust1 := customer.NewCustomer("Bobby", 21, 20000, 200, true)

    discount, err := cust1.CalcDiscount(1000)

    cust1.WrOffDebt()
    fmt.Println("first customer balance after wr. off balance", cust1.Balance)
    
    fmt.Println("first customer balance", cust1.Balance)
    if err == nil {
        fmt.Println(discount)
    } else {
        fmt.Println(err)
    }
}
