package price

type Price struct {
	Amount   float64
	Currency string
}

func NewPrice(amount float64, currency string) Price {
	return Price{
		Amount:   amount,
		Currency: currency,
	}
}
