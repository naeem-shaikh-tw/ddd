package cart

import (
	"ddd/item"
)

var id int = 0

type Cart struct {
	Id           int
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

func (c *Cart) Equals(c2 Cart) bool {
	return c.Id == c2.Id
}
func NewCart() Cart {
	id += 1
	return Cart{
		Id: id,
	}
}
