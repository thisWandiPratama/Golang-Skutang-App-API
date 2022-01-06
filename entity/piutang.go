package entity

import "time"

//Book struct represents books table in database
type Piutang struct {
	ID             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name           string `gorm:"type:varchar(255)" json:"title"`
	Alamat         string `gorm:"type:text" json:"alamat"`
	Nowhatsapp     uint64 `gorm:"type:varchar(255)" json:"no_whatsapp"`
	NominalPiutang uint64 `gorm:"type:varchar(255)" json:"nominal_piutang"`
	CicilanPiutang uint64 `gorm:"type:varchar(255)" json:"cicilan_piutang"`
	DeadlineLunas  string `gorm:"type:varchar(255)" json:"deadline_luas"`
	IsLunas        bool   `gorm:"type:bool;default:0;not null" json:"is_lunas"`
	UserID         uint64 `gorm:"not null" json:"-"`
	User           User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
