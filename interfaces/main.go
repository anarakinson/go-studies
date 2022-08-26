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

    // fmt.Println(truck.Type())
    fmt.Print("truck type: ")
    product.PrintType(truck)
    fmt.Print("engine type: ")
    product.PrintType(engine)

    fmt.Println("Total price:", bask.Total())

    fmt.Println("Total car price:", bask.CarTotalPrice())
}
