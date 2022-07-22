package customer

import (
	"ddd/domain/banking/account"
	"ddd/domain/banking/address"
)

var id = 0

type Customer struct {
	Id       int
	Accounts []*account.Account
	Address  address.Address
}

func (c *Customer) UpdateAddress(address address.Address) {
	c.Address = address
	for _, a := range c.Accounts {
		a.UpdateAddress(address)
	}
}

func (c *Customer) AddAccount(acc *account.Account) {
	c.Accounts = append(c.Accounts, acc)
}
func NewCustomer(address address.Address) Customer {
	id++
	return Customer{
		Id:      id,
		Address: address,
	}
}
