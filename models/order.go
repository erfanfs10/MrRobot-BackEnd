package models

import (
	"encoding/json"
	"time"
)

type CartItem struct {
	ProductDetail
	Quantity *int `json:"quantity"`
}

type OrderCreate struct {
	UserID        *int        `json:"user_id"`
	AddressID     *int        `json:"address_id"`
	TotalProducts *int        `json:"total_products"`
	ShippingPrice *int        `json:"shipping_price"`
	Cart          *[]CartItem `json:"cart"`
}

type Order struct {
	OrderID        *int             `json:"id" db:"order_id"`
	Amount         *string          `json:"amount" db:"order_amount"`
	ShippingPrice  *string          `json:"shipping_price" db:"order_shipping_price"`
	TotalAmount    *string          `json:"total_amount" db:"order_total_amount"`
	Status         *string          `json:"status" db:"order_status"`
	TrackingNumber *string          `json:"tracking_number" db:"order_tracking_number"`
	AddressTitle   *string          `json:"address_title" db:"address_title"`
	AddressAddress *string          `json:"address_address" db:"address_address"`
	OrderItems     *json.RawMessage `json:"order_items" db:"order_items"`
	Created        *time.Time       `json:"created" db:"order_created"`
}
