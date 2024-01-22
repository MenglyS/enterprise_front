package models

type Position struct {
	PositionId uint `gorm:"primary_key"`
	Name       string
}
