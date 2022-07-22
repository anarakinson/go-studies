package customer

import (
    "errors"
)


const DEFAULT_DISCOUNT = 500

type Customer struct {
    Name string
    Age int
    Balance int
    debt int
    discount bool
}

func NewCustomer(
        name string,
        age int,
        balance int,
        debt int,
        discount bool,
    ) (*Customer) {
        return &Customer{
            Name: name,
            Age: age,
            Balance: balance,
            debt: debt,
            discount: discount,
        }
}

func (cust *Customer) WrOffDebt() (error) {
    if (cust.debt > cust.Balance) {
        return errors.New("Not enouth funds")
    }

    cust.Balance -= cust.debt
    cust.debt = 0
    return nil
}

func (cust *Customer) CalcDiscount(params ...interface{}) (int, error) {
    price := params[0].(int)
    if !cust.discount {
        return 0, errors.New("Discount not available")
    }
    result := price - (DEFAULT_DISCOUNT - cust.debt)
    if result < 0 {
        result = 0
    }
    return result, nil
}
