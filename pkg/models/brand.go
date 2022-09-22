package models

type Brand struct {
	Id           int64  `json:"id" gorm:"primaryKey"`
	Name_brand   string `json:"name_brand"`
	Date_created string `json:"created_at"`
}
