package models

import "time"

type User struct {
	UserId     uint `gorm:"primary_key"`
	Username   string
	Email      string
	RoleId     int
	Dob        string
	PositionId int
	Phone      string
	Salary     float32
	Profile    *string
	Password   string
	CreatedBy  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     int
}
