package product

type Product interface {
    Price() int64
}

type Vehicle interface {
    Type() string
    Model() string
}
