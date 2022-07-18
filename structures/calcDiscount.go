package main

func CalcDiscount(params ...interface{}) (int, error) {
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
