package basket

import (
    // "fmt"
    "internal/product"
)

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

func (b *Basket) CarTotalPrice() int64 {
    var result int64
    for _, p := range b.products {
        if !b.isCar(p) {
            continue
        }
        result += p.Price()
    }
    return result
}

func (b *Basket) isCar(p interface{}) bool {
    vehicle, ok := p.(product.Vehicle)
    // fmt.Println(vehicle, ok)
    if !ok {
        return false
    }
    if vehicle.Type() != product.CarVehicleType {
        return false
    }
    return true
}
