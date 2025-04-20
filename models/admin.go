package models

type Role string

const (
	RoleAdmin          Role = "admin"
	RoleSuperAdmin     Role = "super_admin"
	RoleSuperSuperAdmin Role = "super_super_admin"
)

type Admin struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"unique" json:"username"`
	Password     string `json:"password"`
	Role         Role   `json:"role" gorm:"default:admin"`
	SessionToken string `gorm:"unique" json:"session_token"`
}