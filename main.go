package ddd

import (
	"ddd/apple_pencil"
	"ddd/cart"
	"ddd/sony_wireless_headphone"
)

func main() {
	c := cart.NewCart()
	applePencil := apple_pencil.NewApplePencil()
	c.Add(applePencil, 1)
	c.Add(sony_wireless_headphone.NewSonyWirelessHeadphone(), 1)

	c.Remove(applePencil)
}
