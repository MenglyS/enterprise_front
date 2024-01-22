package models

type Language struct {
	LanguageId uint ` gorm:"primary_key"`
	Name       string
}
