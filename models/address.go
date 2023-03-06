package models

import "time"

type Addresses struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	Name      string       `json:"name" gorm:"varchar(255)"`
	Phone     string       `json:"phone" gorm:"varchar(255)"`
	Address   string       `json:"address" gorm:"type: text"`
	PostCode  string       `json:"post_code" gorm:"varchar(255)"`
	UserID    int          `json:"-" gorm:"int"`
	User      UserResponse `json:"-"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}

func (Addresses) TableName() string {
	return "addresses"
}

type AddressesResponse struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	PostCode string `json:"post_code"`
	UpdateAt string `json:"update_at"`
	UserID   int    `json:"-"`
}

func (AddressesResponse) TableName() string {
	return "addresses"
}
