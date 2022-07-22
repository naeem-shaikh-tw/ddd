package item

import "ddd/product"

type Item struct {
	Product  product.Product
	Quantity int
}

func (i *Item) GetProducts() []product.Product {
	var products []product.Product
	for index := 0; index < i.Quantity; index++ {
		products = append(products, i.Product)
	}
	return products
}

func NewItem(p product.Product, quantity int) Item {
	return Item{
		Product:  p,
		Quantity: quantity,
	}
}
