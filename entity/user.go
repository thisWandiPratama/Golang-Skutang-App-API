package entity

import "time"

//User represents users table in database
type User struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Email      string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password   string `gorm:"->;<-;not null" json:"-"`
	Nowhatsapp uint64 `gorm:"type:varchar(255)" json:"no_whatsapp"`
	Alamat     string `gorm:"type:text" json:"alamat"`
	Token      string `gorm:"-" json:"token,omitempty"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	// Books      *[]Book `json:"books,omitempty"`
}
