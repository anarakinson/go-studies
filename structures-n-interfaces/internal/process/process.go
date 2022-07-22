package process

type Debtor interface {
    WrOffDebt() error
}

type Partner interface {
    WrOffDebt() error
}

type Discounter interface {
    CalcDiscount(params ...interface{}) (int, error)
}
