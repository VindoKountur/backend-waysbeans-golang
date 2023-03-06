package models

import "time"

type Profile struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	Phone     string       `json:"phone" gorm:"varchar(255)"`
	Photo     string       `json:"photo" gorm:"varchar(255)"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// Associated with (User)
type ProfileResponse struct {
	Phone  string `json:"phone"`
	Gender string `json:"gender"`
	Photo  string `json:"photo"`
	UserID int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
