package models

import "time"

type Education struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	FieldOfStudy string `json:"field_of_study"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Grade       string `json:"grade"`
}