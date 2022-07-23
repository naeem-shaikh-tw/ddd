package product

import (
	"ddd/domain/e-commerce/price"
)

type Product struct {
	Name          string
	Price         price.Price
	WeightInGrams float64
}

func NewProduct(name string, price price.Price, weightInGrams float64) Product {
	return Product{
		Name:          name,
		Price:         price,
		WeightInGrams: weightInGrams,
	}
}
