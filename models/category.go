package models

type Category struct {
	CategoryId uint ` gorm:"primary_key"`
	Name       string
}

func (Category) TableName() string {
	return "categories"
}
