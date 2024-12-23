package models

// "time"
// "github.com/google/uuid"

type Post struct {
	Userid string `gorm:"type:varchar(16);uniqueIndex;primary_key" json:"userid"`
	K      string `gorm:"type:varchar(32);uniqueIndex;primary_key" json:"k"`
	V      string `gorm:"" json:"v"`
}

type CreatePostRequest struct {
	Userid string `json:"userid" binding:"required"`
	K      string `json:"k"  binding:"required"`
	V      string `json:"v" binding:""`
}

type GetPostRequest struct {
	Userid string `json:"userid" binding:"required"`
	K      string `json:"k"  binding:"required"`
}
