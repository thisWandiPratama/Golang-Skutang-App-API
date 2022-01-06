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
type CicilanHutangService interface {
	UpdateCicilan(b dto.CicilanUpdateDTO) entity.Hutang
	CicilanAllowedToEdit(userID string, hutangID uint64) bool
}

type cicilanhutangService struct {
	cicilanhutangRepository repository.CicilanHutangRepository
}

//NewBookService .....
func NewCicilanHutangService(cicilanhutangRepo repository.CicilanHutangRepository) CicilanHutangService {
	return &cicilanhutangService{
		cicilanhutangRepository: cicilanhutangRepo,
	}
}

func (service *cicilanhutangService) UpdateCicilan(b dto.CicilanUpdateDTO) entity.Hutang {
	cicilanhutang := entity.Hutang{}
	err := smapping.FillStruct(&cicilanhutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	cicilanhutang.CicilanHutang = b.CicilanHutang
	res := service.cicilanhutangRepository.UpdateCicilanHutang(cicilanhutang)
	// fmt.Println(res)
	return res
}

func (service *cicilanhutangService) CicilanAllowedToEdit(userID string, hutangID uint64) bool {
	b := service.cicilanhutangRepository.FindCicilanHutangByID(hutangID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
