package customer

// import (
//     "errors"
// )

type Partner struct {
    Name string
    Age int
    balance int
    debt int
}

func NewPartner(
        name string,
        age int,
        balance int,
        debt int,
    ) (*Partner) {
    return &Partner {
        Name: name,
        Age: age,
        balance: balance,
        debt: debt,
    }
}

func (partner *Partner) WrOffDebt() (error) {
    partner.debt = 0
    return nil
}

func (partner *Partner) Balance() (int) {
    return partner.balance
}

func (partner *Partner) Debt() (int) {
    return partner.debt
}
