package main

import (
	"ddd/apple_pencil"
	"ddd/cart"
	"ddd/sony_wireless_headphone"
)

func main() {
	c := cart.NewCart()
	c.Add(apple_pencil.NewApplePencil(), 2)
	c.Add(sony_wireless_headphone.NewSonyWirelessHeadphone(), 1)
}
