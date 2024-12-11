package models

import (
	"time"

	"github.com/google/uuid"
)

// ユーザのデータモデル
type User struct {
	Userid    string    `gorm:"type:varchar(16);uniqueIndex;primary_key"`
	ID        uuid.UUID `gorm:"type:uuid;uniqueIndex;default:uuid_generate_v4();"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"type:varchar(255);"`
	Provider  string    `gorm:""`
	Photo     string    `gorm:""`
	Verified  bool      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type SignUpInput struct {
	Userid          string `json:"userid" binding:"required,min=4,max=16,alphanum"`
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:""`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:""`
}

type SignInInput struct {
	Userid   string `json:"userid"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// ユーザ情報の応答
type UserResponse struct {
	Userid    string    `json:"userid,omitempty"`
	Name      string    `json:"name,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
