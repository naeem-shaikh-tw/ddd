package order

import (
	"ddd/domain/e-commerce/product"
)

type Order struct {
	Products []product.Product
}

func (o *Order) GetOrder() []product.Product {
	return o.Products
}

func (o *Order) Cost() float64 {
	var totalOrderCost float64 = 0
	for _, p := range o.Products {
		totalOrderCost += p.Price.Amount
	}
	return totalOrderCost
}

func NewOrder(product []product.Product) Order {
	return Order{
		Products: product,
	}
}
