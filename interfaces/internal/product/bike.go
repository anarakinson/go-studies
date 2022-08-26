package product

const bikeVehicleType = "bike"

type Bike struct {
    price int64
    age int
    model string
}

func NewBike(model string, age int, price int64) *Bike {
    return &Bike{
        price: price,
        age: age,
        model: model,
    }
}

func (b *Bike) Type() string {
    return bikeVehicleType
}

func (b *Bike) Model() string {
    return b.model
}

func (b *Bike) Price() int64 {
    return b.price
}

func (b *Bike) Age() int {
    return b.age
}
