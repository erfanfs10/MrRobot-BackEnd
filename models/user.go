package models

type UserGetORCreate struct {
	Name  *string `json:"name" db:"name" form:"name"`
	Email *string `json:"email" db:"email" form:"email"`
	Image *string `json:"image" db:"image"`
}
