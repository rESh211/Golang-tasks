package products

//Task one.
type DeliveryMethod interface {
	CalculateShippingCost(int) int
}

// Стратегия курьерской доставки.
type CourierDelivery struct{}

func (CourierDelivery) CalculateShippingCost(numder int) int {
	return numder + 100
}

// Стратегия почтовой доставки.
type PostDelivery struct{}

func (PostDelivery) CalculateShippingCost(number int) int {
	return number + 50
}

// Стратегия самовывоза.
type PickupDelivery struct{}

func (PickupDelivery) CalculateShippingCost(number int) int {
	return number
}

// Функция для расчета стоимости доставки
func CalculateOrderShippingCost(delivery DeliveryMethod, number int) int {
	return delivery.CalculateShippingCost(number)
}
