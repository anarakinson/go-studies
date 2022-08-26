package product

const carVehicleType = "car"

type Car struct {
    price int64
    age int
    model string
}

func NewCar(model string, age int, price int64) *Car {
    return &Car{
        price: price,
        age: age,
        model: model,
    }
}

func (c *Car) Type() string {
    return carVehicleType
}

func (c *Car) Model() string {
    return c.model
}

func (c *Car) Price() int64 {
    return c.price
}

func (c *Car) Age() int {
    return c.age
}
