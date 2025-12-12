package models

type CartItem struct {
	ProductDetail
	Quantity *int `json:"quantity"`
}

type OrderCreate struct {
	UserID        *int    `json:"user_id"`
	AddressID     *int    `json:"address_id"`
	TotalProducts *int    `json:"total_products"`
	ShippingPrice *int    `json:"shipping_price"`
	Cart          *[]CartItem `json:"cart"`
}
