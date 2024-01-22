package models

import "time"

type Applicant struct {
	ApplicantId   uint `gorm:"primary_key"`
	Name          string
	Email         string
	Phone         string
	ScheduledDate time.Time
	LanguageIds   *string
	SkillIds      *string
	Address       *string
	Experience    *string
	Education     *string
	Summary       *string
	JobId         int
	CvFile        *string
	ApprovedBy    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        int
}
