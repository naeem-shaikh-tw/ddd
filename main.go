package main

import (
	"ddd/cart"
	"ddd/item"
	"ddd/product"
	"fmt"
)

func main() {
	c := cart.NewCart()

	// Add Apple Pencil
	applePencil := product.NewProduct("Apple Pencil")
	applePencilItem := item.NewItem(applePencil, 2)
	c.Add(applePencilItem)

	// Add Sony Wireless Headphone
	sonyWirelessHeadphone := product.NewProduct("Sony wireless headphone")
	c.Add(item.NewItem(sonyWirelessHeadphone, 1))

	fmt.Println("Cart: ", c.GetItems())

	// Remove Apple Pencil
	c.Remove(applePencilItem)
	
	removedItems := c.GetRemovedItems()
	fmt.Println("Removed Items:", removedItems)
}
