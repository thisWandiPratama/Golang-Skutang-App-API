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
type PiutangService interface {
	Insert(b dto.PiutangCreateDTO) entity.Piutang
	Update(b dto.PiutangUpdateDTO) entity.Piutang
	Delete(b entity.Piutang)
	All(userID string) []entity.Piutang
	FindByID(piutangID uint64) entity.Piutang
	IsAllowedToEdit(userID string, piutangID uint64) bool
}

type piutangService struct {
	piutangRepository repository.PiutangRepository
}

//NewBookService .....
func NewPiutangService(piutangRepo repository.PiutangRepository) PiutangService {
	return &piutangService{
		piutangRepository: piutangRepo,
	}
}

func (service *piutangService) Insert(b dto.PiutangCreateDTO) entity.Piutang {
	piutang := entity.Piutang{}
	err := smapping.FillStruct(&piutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.piutangRepository.InsertPiutang(piutang)
	return res
}

func (service *piutangService) Update(b dto.PiutangUpdateDTO) entity.Piutang {
	piutang := entity.Piutang{}
	err := smapping.FillStruct(&piutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.piutangRepository.UpdatePiutang(piutang)
	return res
}

func (service *piutangService) Delete(b entity.Piutang) {
	service.piutangRepository.DeletePiutang(b)
}

func (service *piutangService) All(userID string) []entity.Piutang {
	return service.piutangRepository.AllPiutang(userID)
}

func (service *piutangService) FindByID(piutangID uint64) entity.Piutang {
	return service.piutangRepository.FindPiutangByID(piutangID)
}

func (service *piutangService) IsAllowedToEdit(userID string, piutangID uint64) bool {
	b := service.piutangRepository.FindPiutangByID(piutangID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
