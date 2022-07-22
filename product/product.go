package product

import "ddd/price"

type Product struct {
	Name  string
	Price price.Price
}

func NewProduct(name string, price price.Price) Product {
	return Product{
		Name:  name,
		Price: price,
	}
}
