package dto

import "time"

//BookUpdateDTO is a model that client use when updating a book
type HutangUpdateDTO struct {
	ID            uint64 `json:"id" form:"id" binding:"required"`
	Name          string `json:"name" form:"name" binding:"required"`
	Alamat        string `json:"alamat" form:"alamat" binding:"required"`
	Nowhatsapp    uint64 `json:"no_whatsapp" form:"no_whatsapp" binding:"required"`
	NominalHutang uint64 `json:"nominal_hutang" form:"nominal_hutang" binding:"required"`
	CicilanHutang uint64 `json:"cicilan_hutang" form:"cicilan_hutang" binding:"required"`
	DeadlineLunas string `json:"deadline_lunas" form:"deadline_lunas" binding:"required"`
	UserID        uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//BookCreateDTO is is a model that clinet use when create a new book
type HutangCreateDTO struct {
	Name          string `json:"name" form:"name" binding:"required"`
	Alamat        string `json:"alamat" form:"alamat" binding:"required"`
	Nowhatsapp    uint64 `json:"no_whatsapp" form:"no_whatsapp" binding:"required"`
	NominalHutang uint64 `json:"nominal_hutang" form:"nominal_hutang" binding:"required"`
	CicilanHutang uint64 `json:"cicilan_hutang" form:"cicilan_hutang" binding:"required"`
	DeadlineLunas string `json:"deadline_lunas" form:"deadline_lunas" binding:"required"`
	UserID        uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type IsLunasUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	Islunas   bool   `json:"is_lunas" form:"is_lunas"`
	UserID    uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CicilanUpdateDTO struct {
	ID            uint64 `json:"id" form:"id" binding:"required"`
	CicilanHutang uint64 `json:"cicilan_hutang" form:"cicilan_hutang" binding:"required"`
	UserID        uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
