package order

import "ddd/product"

type Order struct {
	Products []product.Product
}

func (o *Order) GetOrder() []product.Product {
	return o.Products
}

func NewOrder(product []product.Product) Order {
	return Order{
		Products: product,
	}
}
