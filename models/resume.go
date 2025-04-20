package models

type Resume struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FileName string `json:"file_name"`
	FileURL  string `json:"file_url"`
	
}