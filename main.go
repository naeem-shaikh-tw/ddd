package ddd

import (
	"ddd/apple_pencil"
	"ddd/cart"
)

func main() {
	c := cart.NewCart()
	c.Add(apple_pencil.NewApplePencil())
}
