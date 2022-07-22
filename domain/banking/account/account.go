package account

import "ddd/domain/banking/address"

type Account struct {
	Id      int
	Address address.Address
}

var id = 0

func (a *Account) UpdateAddress(address address.Address) {
	a.Address = address
}
func NewAccount(address address.Address) Account {
	id++
	return Account{
		Id:      id,
		Address: address,
	}
}
