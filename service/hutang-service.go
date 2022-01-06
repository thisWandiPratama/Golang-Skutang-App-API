package service

import (
	"fmt"
	"golang_api_hupiutang/dto"
	"golang_api_hupiutang/entity"
	"golang_api_hupiutang/repository"
	"log"

	"github.com/mashingan/smapping"
)

//BookService is a ....
type HutangService interface {
	Insert(b dto.HutangCreateDTO) entity.Hutang
	Update(b dto.HutangUpdateDTO) entity.Hutang
	Delete(b entity.Hutang)
	All(userID string) []entity.Hutang
	FindByID(hutangID uint64) entity.Hutang
	IsAllowedToEdit(userID string, hutangID uint64) bool
}

type hutangService struct {
	hutangRepository repository.HutangRepository
}

//NewBookService .....
func NewHutangService(hutangRepo repository.HutangRepository) HutangService {
	return &hutangService{
		hutangRepository: hutangRepo,
	}
}

func (service *hutangService) Insert(b dto.HutangCreateDTO) entity.Hutang {
	hutang := entity.Hutang{}
	err := smapping.FillStruct(&hutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.hutangRepository.InsertHutang(hutang)
	return res
}

func (service *hutangService) Update(b dto.HutangUpdateDTO) entity.Hutang {
	hutang := entity.Hutang{}
	err := smapping.FillStruct(&hutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.hutangRepository.UpdateHutang(hutang)
	return res
}

func (service *hutangService) Delete(b entity.Hutang) {
	service.hutangRepository.DeleteHutang(b)
}

func (service *hutangService) All(userID string) []entity.Hutang {
	return service.hutangRepository.AllHutang(userID)
}

func (service *hutangService) FindByID(hutangID uint64) entity.Hutang {
	return service.hutangRepository.FindHutangByID(hutangID)
}

func (service *hutangService) IsAllowedToEdit(userID string, hutangID uint64) bool {
	b := service.hutangRepository.FindHutangByID(hutangID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
