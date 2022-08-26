package basket

import "internal/product"

type Basket struct {
    orderID string
    products []product.Product
}

func NewBasket(orderID string, products []product.Product) *Basket {
    return &Basket{
        orderID: orderID,
        products: products,
    }
}

func (b *Basket) Total() int64 {
    var result int64
    for _, p := range b.products {
        result += p.Price()
    }

    return result
}
