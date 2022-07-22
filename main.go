package main

import (
	"ddd/cart"
	"ddd/item"
	"ddd/price"
	"ddd/product"
	"fmt"
)

func main() {
	c := cart.NewCart()

	// Add Apple Pencil
	applePencil := product.NewProduct("Apple Pencil", price.NewPrice(20, "USD"))
	applePencilItem := item.NewItem(applePencil, 2)
	c.Add(applePencilItem)

	// Add Sony Wireless Headphone
	sonyWirelessHeadphone := product.NewProduct("Sony wireless headphone", price.NewPrice(20, "USD"))
	c.Add(item.NewItem(sonyWirelessHeadphone, 1))

	fmt.Println("Cart: ", c.GetItems())

	// Remove Apple Pencil
	c.Remove(applePencilItem)

	removedItems := c.GetRemovedItems()
	fmt.Println("Removed Items:", removedItems)

	c1 := cart.NewCart()
	c2 := cart.NewCart()
	c1.Add(applePencilItem)
	c2.Add(applePencilItem)
	fmt.Println(c1.Equals(c1)) // true
	fmt.Println(c1.Equals(c2)) // false
}
