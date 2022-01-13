package repository

import (
	"golang_api_hupiutang/entity"

	"gorm.io/gorm"
)

//BookRepository is a ....
type CicilanPiutangRepository interface {
	UpdateCicilanPiutang(h entity.Piutang) entity.Piutang
	FindCicilanPiutangByID(piutangID uint64) entity.Piutang
}

type cicilanPiutangConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewCicilanPiutangRepository(dbConn *gorm.DB) CicilanPiutangRepository {
	return &cicilanPiutangConnection{
		connection: dbConn,
	}
}

func (db *cicilanPiutangConnection) FindCicilanPiutangByID(piutangID uint64) entity.Piutang {
	var cicilan entity.Piutang
	db.connection.Preload("User").Find(&cicilan, piutangID)
	return cicilan
}

func (db *cicilanPiutangConnection) UpdateCicilanPiutang(h entity.Piutang) entity.Piutang {
	cicilan := h.CicilanPiutang
	db.connection.Preload("User").First(&h)
	totalCicilan := cicilan + h.CicilanPiutang
	h.CicilanPiutang = cicilan + h.CicilanPiutang
	h.NominalPiutang = h.NominalPiutang - h.CicilanPiutang
	if h.NominalPiutang == 0 || totalCicilan >= h.NominalPiutang || h.NominalPiutang >= totalCicilan {
		h.NominalPiutang = 0
		h.IsLunas = true
	}
	db.connection.Save(&h)
	return h
}
