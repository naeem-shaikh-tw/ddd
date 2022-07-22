package competitor_based_pricer

import (
	"ddd/domain/e-commerce/price"
)

func GetDiscountedPrice(discount float64, competitorPrice price.Price) price.Price {
	return price.NewPrice(competitorPrice.Amount/100*discount, competitorPrice.Currency)
}
