package repository

import (
	"golang_api_hupiutang/entity"

	"gorm.io/gorm"
)

//BookRepository is a ....
type PiutangRepository interface {
	InsertPiutang(h entity.Piutang) entity.Piutang
	UpdatePiutang(h entity.Piutang) entity.Piutang
	DeletePiutang(h entity.Piutang)
	AllPiutang(userID string) []entity.Piutang
	FindPiutangByID(piutangID uint64) entity.Piutang
}

type piutangConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewPiutangRepository(dbConn *gorm.DB) PiutangRepository {
	return &piutangConnection{
		connection: dbConn,
	}
}

func (db *piutangConnection) InsertPiutang(h entity.Piutang) entity.Piutang {
	db.connection.Save(&h)
	db.connection.Preload("User").Find(&h)
	return h
}

func (db *piutangConnection) UpdatePiutang(h entity.Piutang) entity.Piutang {
	getDulu := h
	db.connection.First(&getDulu)
	// fmt.Println(h.CreatedAt)
	h.CreatedAt = getDulu.UpdatedAt
	db.connection.Save(&h)
	// fmt.Println(h.Name)
	db.connection.Preload("User").Find(&h)
	return h
}

func (db *piutangConnection) DeletePiutang(h entity.Piutang) {
	db.connection.Delete(&h)
}

func (db *piutangConnection) FindPiutangByID(piutangID uint64) entity.Piutang {
	var hutang entity.Piutang
	db.connection.Preload("User").Find(&hutang, piutangID)
	return hutang
}

func (db *piutangConnection) AllPiutang(userID string) []entity.Piutang {
	var piutangs []entity.Piutang
	db.connection.Preload("User").Where("user_id", userID).Find(&piutangs)
	return piutangs
}
