package main

import (
	"ddd/domain/e-commerce/cart"
	"ddd/domain/e-commerce/domain_service/competitor_based_pricer"
	"ddd/domain/e-commerce/item"
	"ddd/domain/e-commerce/order"
	"ddd/domain/e-commerce/price"
	"ddd/domain/e-commerce/product"
	"fmt"
)

func main() {
	priceMap := map[string]float64{"Apple Pencil": 5, "Sony wireless headphone": 30}
	var discount float64 = 10

	c := cart.NewCart()

	// Add Apple Pencil
	applePencilDiscountedPrice := competitor_based_pricer.GetDiscountedPrice(discount, price.NewPrice(priceMap["Apple Pencil"], "USD"))
	applePencil := product.NewProduct("Apple Pencil", applePencilDiscountedPrice)
	applePencilItem := item.NewItem(applePencil, 2)
	c.Add(applePencilItem)

	// Add Sony Wireless Headphone
	sonyWirelessHeadphoneDiscountedPrice := competitor_based_pricer.GetDiscountedPrice(discount, price.NewPrice(priceMap["Sony wireless headphone"], "USD"))

	sonyWirelessHeadphone := product.NewProduct("Sony wireless headphone", sonyWirelessHeadphoneDiscountedPrice)
	c.Add(item.NewItem(sonyWirelessHeadphone, 2))

	fmt.Println("Cart: ", c.GetItems())

	// Remove Apple Pencil
	//c.Remove(applePencilItem)

	removedItems := c.GetRemovedItems()
	fmt.Println("Removed Items:", removedItems)

	c1 := cart.NewCart()
	c2 := cart.NewCart()
	c1.Add(applePencilItem)
	c2.Add(applePencilItem)
	fmt.Println(c1.Equals(c1)) // true
	fmt.Println(c1.Equals(c2)) // false

	products := c.Checkout()
	o := order.NewOrder(products)
	fmt.Println("Order: ", o.GetOrder())
}
