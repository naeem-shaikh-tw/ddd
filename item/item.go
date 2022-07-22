package item

import "ddd/product"

type Item struct {
	Product  product.Product
	Quantity int
}

func NewItem(p product.Product, quantity int) Item {
	return Item{
		Product:  p,
		Quantity: quantity,
	}
}
