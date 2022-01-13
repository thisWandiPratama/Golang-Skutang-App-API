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
type CicilanPiutangService interface {
	UpdateCicilanPiutang(b dto.CicilanUpdatePiutangDTO) entity.Piutang
	CicilanPiutangAllowedToEdit(userID string, piutangID uint64) bool
}

type cicilanpiutangService struct {
	cicilanpiutangRepository repository.CicilanPiutangRepository
}

//NewBookService .....
func NewCicilanPiutangService(cicilanpiutangRepo repository.CicilanPiutangRepository) CicilanPiutangService {
	return &cicilanpiutangService{
		cicilanpiutangRepository: cicilanpiutangRepo,
	}
}

func (service *cicilanpiutangService) UpdateCicilanPiutang(b dto.CicilanUpdatePiutangDTO) entity.Piutang {
	cicilanpiutang := entity.Piutang{}
	err := smapping.FillStruct(&cicilanpiutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	cicilanpiutang.CicilanPiutang = b.CicilanPiutang
	res := service.cicilanpiutangRepository.UpdateCicilanPiutang(cicilanpiutang)
	// fmt.Println(res)
	return res
}

func (service *cicilanpiutangService) CicilanPiutangAllowedToEdit(userID string, putangID uint64) bool {
	b := service.cicilanpiutangRepository.FindCicilanPiutangByID(putangID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
