package cart

import (
	"ddd/item"
)

type Cart struct {
	Items []item.Item
}

func (c Cart) Add(i item.Item, q int) {
	for quantity := 0; quantity < q; quantity++ {
		c.Items = append(c.Items, quantity)
	}
}

func NewCart() Cart {
	return Cart{}
}
