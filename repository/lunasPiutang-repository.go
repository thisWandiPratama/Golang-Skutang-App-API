package repository

import (
	"golang_api_hupiutang/entity"

	"gorm.io/gorm"
)

//BookRepository is a ....
type LunasPiutangRepository interface {
	UpdateIsLunasPiutang(h entity.Piutang) entity.Piutang
	FindLunasPiutangByID(piutangID uint64) entity.Piutang
}

type lunasPiutangConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewLunasPiutangRepository(dbConn *gorm.DB) LunasPiutangRepository {
	return &lunasPiutangConnection{
		connection: dbConn,
	}
}

func (db *lunasPiutangConnection) FindLunasPiutangByID(piutangID uint64) entity.Piutang {
	var piutang entity.Piutang
	db.connection.Preload("User").Find(&piutang, piutangID)
	return piutang
}

func (db *lunasPiutangConnection) UpdateIsLunasPiutang(h entity.Piutang) entity.Piutang {
	db.connection.Preload("User").First(&h)
	getDulu := h
	h.CreatedAt = getDulu.UpdatedAt
	h.IsLunas = !h.IsLunas
	db.connection.Save(&h)
	return h
}
