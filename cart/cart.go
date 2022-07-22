package cart

import (
	"ddd/item"
)

type Cart struct {
	Items        []item.Item
	RemovedItems []item.Item
}

func (c *Cart) Add(i item.Item) {
	c.Items = append(c.Items, i)
}

func (c *Cart) Remove(i item.Item) {
	for index := 0; index < len(c.Items); index++ {
		if c.Items[index] == i {
			c.Items = append(c.Items[:index], c.Items[index+1:]...)
			c.RemovedItems = append(c.RemovedItems, i)
		}
	}
}

func (c *Cart) GetRemovedItems() []item.Item {
	return c.RemovedItems
}

func (c *Cart) GetItems() []item.Item {
	return c.Items
}
func NewCart() Cart {
	return Cart{}
}
