package debtor

type Debtor interface {
    WrOffDebt() error
}

type Partner interface {
    WrOffDebt() error
}
