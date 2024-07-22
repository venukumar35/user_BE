package response

import "time"

type UserResponse struct {
	Id       uint    `json:"id" gorm:"column:id"`
	Name     string  `json:"name" gorm:"column:name"`
	Email    string  `json:"email" gorm:"column:email"`
	Mobile   string  `json:"mobile" gorm:"column:mobile"`
	RoleId   int     `json:"roleId" gorm:"column:roleId"`
	WebToken *string `json:"webToken" gorm:"column:webToken"`
	IsActive bool    `json:"isActive" gorm:"column:isActive"`
}

type LoginResponse struct {
	Id           uint      `json:"id" gorm:"column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Email        string    `json:"email" gorm:"column:email"`
	Mobile       string    `json:"mobile" gorm:"column:mobile"`
	RoleId       int       `json:"roleId" gorm:"column:roleId"`
	WebToken     *string   `json:"webToken" gorm:"column:webToken"`
	IsActive     bool      `json:"isActive" gorm:"column:isActive"`
	RoleName     string    `json:"roleName" gorm:"column:roleName"`
	OtpCreatedAt time.Time `json:"otpCreatedAt" gorm:"column:otpCreatedAt"`
}
