package repository

import (
	"golang_api_hupiutang/entity"

	"gorm.io/gorm"
)

//BookRepository is a ....
type LunasHutangRepository interface {
	UpdateIsLunasHutang(h entity.Hutang) entity.Hutang
	FindLunasHutangByID(hutangID uint64) entity.Hutang
}

type lunasHutangConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewLunasHutangRepository(dbConn *gorm.DB) LunasHutangRepository {
	return &lunasHutangConnection{
		connection: dbConn,
	}
}

func (db *lunasHutangConnection) FindLunasHutangByID(hutangID uint64) entity.Hutang {
	var hutang entity.Hutang
	db.connection.Preload("User").Find(&hutang, hutangID)
	return hutang
}

func (db *lunasHutangConnection) UpdateIsLunasHutang(h entity.Hutang) entity.Hutang {
	db.connection.Preload("User").First(&h)
	getDulu := h
	h.CreatedAt = getDulu.UpdatedAt
	h.IsLunas = !h.IsLunas
	db.connection.Save(&h)
	return h
}
