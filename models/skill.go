package models

type Skill struct {
	SkillId uint `gorm:"primary_key"`
	Name    string
}
