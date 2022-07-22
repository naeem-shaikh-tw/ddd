package cart

import "ddd/item"

type Cart struct {
	Items []item.Item
}

func (c Cart) Add(i item.Item) {
	c.Items = append(c.Items, i)
}

func NewCart() Cart {
	return Cart{}
}
