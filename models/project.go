package models

type Project struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	RepoURL     string `json:"repo_url"`
	DemoURL     string `json:"demo_url"`
	TechStack   string `json:"tech_stack"`
}