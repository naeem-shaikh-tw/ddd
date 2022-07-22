package main

import (
	"ddd/domain/banking/account"
	"ddd/domain/banking/address"
	"ddd/domain/banking/customer"
	"fmt"
)

func main() {
	addr := address.NewAddress("Pune")
	accAddr := address.NewAddress("Mumbai")

	customer := customer.NewCustomer(addr)
	acc := account.NewAccount(accAddr)
	customer.AddAccount(&acc)
	fmt.Println(customer.Address, customer.Accounts[0].Address)

	newAddr := address.NewAddress("Delhi")
	customer.UpdateAddress(newAddr)

	fmt.Println(customer.Address, customer.Accounts[0].Address)
}
