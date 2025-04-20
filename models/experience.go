package models

type Experience struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Company     string `json:"company"`
	Role        string `json:"role"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Present     bool   `json:"present" gorm:"default:false"`
}