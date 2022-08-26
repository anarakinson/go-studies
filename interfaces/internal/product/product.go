package product

import (
    "fmt"
)

type Product interface {
    Price() int64
}

type Vehicle interface {
    Type() string
    Model() string
}

func PrintType(p interface{}) {
    switch t := p.(type){
    default:
        fmt.Println("unexpected type", t)
    case Engine:
        fmt.Println("engine", t)
    case *Car:
        fmt.Println("pointer to car:", *t)
    case *Bike:
        fmt.Println("pointer to bike", *t)
    case *Engine:
        fmt.Println("pointer to engine", *t)
    }
}
