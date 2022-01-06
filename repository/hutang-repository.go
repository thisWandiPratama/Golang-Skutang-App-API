package repository

import (
	"fmt"
	"golang_api_hupiutang/entity"

	"gorm.io/gorm"
)

//BookRepository is a ....
type HutangRepository interface {
	InsertHutang(h entity.Hutang) entity.Hutang
	UpdateHutang(h entity.Hutang) entity.Hutang
	DeleteHutang(h entity.Hutang)
	AllHutang(userID string) []entity.Hutang
	FindHutangByID(hutangID uint64) entity.Hutang
}

type hutangConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewHutangRepository(dbConn *gorm.DB) HutangRepository {
	return &hutangConnection{
		connection: dbConn,
	}
}

func (db *hutangConnection) InsertHutang(h entity.Hutang) entity.Hutang {
	db.connection.Save(&h)
	db.connection.Preload("User").Find(&h)
	return h
}

func (db *hutangConnection) UpdateHutang(h entity.Hutang) entity.Hutang {
	db.connection.Save(&h)
	fmt.Println(h.Name)
	db.connection.Preload("User").Find(&h)
	return h
}

func (db *hutangConnection) DeleteHutang(h entity.Hutang) {
	db.connection.Delete(&h)
}

func (db *hutangConnection) FindHutangByID(hutangID uint64) entity.Hutang {
	var hutang entity.Hutang
	db.connection.Preload("User").Find(&hutang, hutangID)
	return hutang
}

func (db *hutangConnection) AllHutang(userID string) []entity.Hutang {
	var hutangs []entity.Hutang
	db.connection.Preload("User").Where("user_id", userID).Find(&hutangs)
	return hutangs

}
