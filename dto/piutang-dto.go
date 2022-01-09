package dto

import "time"

//BookUpdateDTO is a model that client use when updating a book
type PiutangUpdateDTO struct {
	ID             uint64 `json:"id" form:"id" binding:"required"`
	Name           string `json:"name" form:"name" binding:"required"`
	Alamat         string `json:"alamat" form:"alamat" binding:"required"`
	Nowhatsapp     uint64 `json:"no_whatsapp" form:"no_whatsapp" binding:"required"`
	NominalPiutang uint64 `json:"nominal_piutang" form:"nominal_piutang" binding:"required"`
	CicilanPiutang uint64 `json:"cicilan_piutang" form:"cicilan_piutang" binding:"required"`
	DeadlineLunas  string `json:"deadline_lunas" form:"deadline_lunas" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

//BookCreateDTO is is a model that clinet use when create a new book
type PiutangCreateDTO struct {
	Name           string `json:"name" form:"name" binding:"required"`
	Alamat         string `json:"alamat" form:"alamat" binding:"required"`
	Nowhatsapp     uint64 `json:"no_whatsapp" form:"no_whatsapp" binding:"required"`
	NominalPiutang uint64 `json:"nominal_piutang" form:"nominal_piutang" binding:"required"`
	CicilanPiutang uint64 `json:"cicilan_piutang" form:"cicilan_piutang" binding:"required"`
	DeadlineLunas  string `json:"deadline_lunas" form:"deadline_lunas" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type IsLunasPiutangUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	Islunas   bool   `json:"is_lunas" form:"is_lunas"`
	UserID    uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CicilanUpdatePiutangDTO struct {
	ID             uint64 `json:"id" form:"id" binding:"required"`
	CicilanPiutang uint64 `json:"cicilan_piutang" form:"cicilan_piutang" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
