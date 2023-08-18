package main

import (
	"designpatterns/factory"
	"fmt"
)

func main() {
	fmt.Println(factory.CreateProduct("A").Use())
	fmt.Println(factory.CreateProduct("B").Use())
}
