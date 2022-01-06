package repository

import (
	"golang_api_hupiutang/entity"

	"gorm.io/gorm"
)

//BookRepository is a ....
type CicilanHutangRepository interface {
	UpdateCicilanHutang(h entity.Hutang) entity.Hutang
	FindCicilanHutangByID(hutangID uint64) entity.Hutang
}

type cicilanHutangConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewCicilanHutangRepository(dbConn *gorm.DB) CicilanHutangRepository {
	return &cicilanHutangConnection{
		connection: dbConn,
	}
}

func (db *cicilanHutangConnection) FindCicilanHutangByID(hutangID uint64) entity.Hutang {
	var cicilan entity.Hutang
	db.connection.Preload("User").Find(&cicilan, hutangID)
	return cicilan
}

func (db *cicilanHutangConnection) UpdateCicilanHutang(h entity.Hutang) entity.Hutang {
	cicilan := h.CicilanHutang
	db.connection.Preload("User").First(&h)
	totalCicilan := cicilan + h.CicilanHutang
	h.CicilanHutang = cicilan + h.CicilanHutang
	h.NominalHutang = h.NominalHutang - h.CicilanHutang
	if h.NominalHutang == 0 || totalCicilan >= h.NominalHutang || h.NominalHutang >= totalCicilan {
		h.NominalHutang = 0
		h.IsLunas = true
	}
	db.connection.Save(&h)
	return h
}
