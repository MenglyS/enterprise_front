package models

import "time"

type Job struct {
	JobId             uint `gorm:"primary_key"`
	Title             string
	CategoryIds       string
	Description       string
	Contact           string
	ExpiryDate        string
	AnnouncementImage *string
	CreatedBy         int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Status            int
}
