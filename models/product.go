package models

type Product struct {
	ID               *int         `json:"id" db:"product_id"`
	Title            *string      `json:"title" db:"product_title"`
	TitleFarsi       *string      `json:"title_farsi" db:"product_title_farsi"`
	Description      *string      `json:"description" db:"product_description"`
	ColorCode        *string      `json:"color_code" db:"product_color_code"`
	Used             *bool        `json:"used" db:"product_used"`
	Status           *string      `json:"status" db:"product_status"`
	ListPrice        *string      `json:"list_price" db:"product_list_price"`
	Tax              *string      `json:"tax" db:"product_tax"`
	Discount         *string      `json:"discount" db:"product_discount"`
	NetPrice         *string      `json:"net_price" db:"product_net_price"`
	Stock            *int         `json:"stock" db:"product_stock"`
	Variant          *string      `json:"variant" db:"product_variant"`
	VariantFarsi     *string      `json:"variant_farsi" db:"product_variant_farsi"`
	PrimaryImage     *string      `json:"primary_image" db:"primary_image"`
	Brand            *string      `json:"brand" db:"brand"`
	BrandFarsi       *string      `json:"brand_farsi" db:"brand_farsi"`
	Category         *string      `json:"category" db:"category"`
	CategoryFarsi    *string      `json:"category_farsi" db:"category_farsi"`
	ProductType      *string      `json:"product_type" db:"product_type"`
	ProductTypeFarsi *string      `json:"product_type_farsi" db:"product_type_farsi"`
	Attributes       []Attributes `json:"attributes"`
}

type ProductIDs struct {
	ProductIDs []int `json:"id" db:"id"`
}

type ProductDetail struct {
	Product
	Point               *string         `json:"point" db:"point"`
	Rates               []Rates         `json:"rates"`
	Images              []ProductImages `json:"produc_images"`
	Attributes          []Attributes    `json:"attributes"`
	Posts               []Post          `json:"posts"`
	WishListsProductIDs []int           `json:"wishlist_product_ids"`
}

type ProductID struct {
	ProductID int `json:"product_id" db:"product_id"`
}
