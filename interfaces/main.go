package main

import(
    "fmt"
    "internal/product"
    "internal/basket"
)

func main() {
    engine := product.NewEngine(2.5, 5000)
    truck := product.NewCar("Pickup", 15, 12000)
    bike := product.NewBike("Ural", 30, 6500)

    bask := basket.NewBasket(
        "1",
        []product.Product{
            engine,
            truck,
            bike,
        },
    )

    fmt.Println(truck.Type())

    fmt.Println(bask.Total())
}
