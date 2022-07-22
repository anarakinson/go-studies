package main

import (
    "fmt"
    "internal/customer"
    "internal/debtor"
)

func main() {
    // cust1 := customer.Customer{
    //     Name: "Billy",
    //     Age: 20,
    //     Balance: 10000,
    //     Debt: 1000,
    // }

    cust1 := customer.NewCustomer("Bobby", 21, 20000, 200, true)
    part1 := customer.NewPartner("Billy", 23, 25000, 250)

    discount, err := cust1.CalcDiscount(1000)
    if err == nil {
        fmt.Println(discount)
    } else {
        fmt.Println(err)
    }

    fmt.Println("***")

    fmt.Println("first customer balance", cust1.Balance)
    startTransaction(cust1)
    fmt.Println("first customer balance after wr. off balance", cust1.Balance)

    fmt.Println("***")

    fmt.Println("first partner: \nbalance:", part1.Balance(), "\ndebt:", part1.Debt())
    startTransaction(part1)
    fmt.Println("first partner after wr. off balance:\nbalance:", part1.Balance(), "\ndebt:", part1.Debt())
}

func startTransaction(dbtr debtor.Debtor) {
    dbtr.WrOffDebt()
}
