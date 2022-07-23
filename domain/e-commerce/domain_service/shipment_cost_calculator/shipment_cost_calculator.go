package shipment_cost_calculator

import "ddd/domain/e-commerce/order"

func CalculateShipmentCost(o order.Order) float64 {
	var shipmentCost float64 = 0
	for _, product := range o.Products {
		shipmentCost += product.WeightInGrams * 0.1
	}
	return shipmentCost
}
