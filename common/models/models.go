package models

// Константы для статусов заказа
const (
	OrderStatusPending   = "pending"
	OrderStatusCompleted = "completed"
	OrderStatusCancelled = "cancelled"
)

// Константы для статусов доставки
const (
	DeliveryStatusScheduled = "scheduled"
	DeliveryStatusInTransit = "in_transit"
	DeliveryStatusDelivered = "delivered"
)
