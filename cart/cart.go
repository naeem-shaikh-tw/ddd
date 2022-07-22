package cart

import "ddd/item"

type Cart struct {
	Items []item.Item
}

func (c Cart) Add(i item.Item, q int) {
	for quantity := 0; quantity < q; quantity++ {
		c.Items = append(c.Items, quantity)
	}
}

func (c Cart) Remove(i item.Item) {
	for index := 0; index < len(c.Items); index++ {
		if c.Items[index] == i {
			c.Items = append(c.Items[:index], c.Items[index+1:]...)
		}
	}

}
func NewCart() Cart {
	return Cart{}
}
