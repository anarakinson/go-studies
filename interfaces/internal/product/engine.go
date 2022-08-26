package product

type Engine struct {
    price int64
    vol float64
}

func NewEngine(vol float64, price int64) *Engine {
    return &Engine{
        vol: vol,
        price: price,
    }
}

func (p *Engine) Price() int64 {
    return p.price
}

func (p *Engine) Volume() float64 {
    return p.vol
}
