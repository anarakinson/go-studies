package main

import (
    "fmt"
    "internal/customer"
    "errors"
)

const DEFAULT_DISCOUNT = 500

func main() {
    cust0 := customer.Customer{
        Name: "Billy",
        Age: 20,
        Balance: 10000,
        Debt: 1000,
    }
    fmt.Println("first customer", cust0)

    cust1 := customer.NewCustomer("Bobby", 21, 20000, 200, true)

    cust1.CalcDiscount = func(params ...interface{}) (int, error) {
        cust := params[0].(*customer.Customer)
        price := params[1].(int)
        if !cust.Discount {
            return 0, errors.New("Discount not available")
        }
        result := price - (DEFAULT_DISCOUNT - cust.Debt)
        if result < 0 {
            result = 0
        }
        return result, nil
    }

    discount, err := cust1.CalcDiscount(cust1, 1000)
    fmt.Println("first customer", cust1)
    if err == nil {
        fmt.Println(discount)
    } else {
        fmt.Println(err)
    }
}
