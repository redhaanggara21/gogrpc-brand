package models

type Brand struct {
	Id           int64  `json:"id" gorm:"primaryKey"`
	Brand_name   string `json:"brand_name"`
	Date_created string `json:"date_created"`
}
