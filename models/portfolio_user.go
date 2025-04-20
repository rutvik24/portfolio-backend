package models

type PortfolioUser struct {
	ID                uint              `gorm:"primaryKey" json:"id"`
	Email             string            `json:"email"`
	Name              string            `json:"name"`
	Locality          string            `json:"locality"`
	Availability      string            `json:"availability"`
	SocialLinks       map[string]string `gorm:"type:json" json:"social_links"`
	Profession        string            `json:"profession"`
	Skills            map[string][]string `gorm:"type:json" json:"skills"`
	SkillsHighlighted []string          `gorm:"type:json" json:"skills_highlighted"`
	Passion           string            `json:"passion"`
	Description       string            `gorm:"type:text" json:"description"`
}