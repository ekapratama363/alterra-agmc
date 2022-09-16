package models

type Books struct {
	Id   string `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name"`
}
