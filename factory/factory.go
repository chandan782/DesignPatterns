// The Factory pattern is used to create objects without specifying the exact class of object that will be created.
package factory

type Product interface {
	Use() string
}

type concreteProductA struct {
}

func (cpa concreteProductA) Use() string {
	return "Concrete Product A"
}

type concreteProductB struct {
}

func (cpb concreteProductB) Use() string {
	return "Concrete Product B"
}

func CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return concreteProductA{}
	case "B":
		return concreteProductB{}
	default:
		return nil
	}
}
