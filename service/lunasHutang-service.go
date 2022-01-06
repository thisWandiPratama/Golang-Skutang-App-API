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
type LunasHutangService interface {
	UpdateLunas(b dto.IsLunasUpdateDTO) entity.Hutang
	LunasIsAllowedToEdit(userID string, hutangID uint64) bool
}

type lunashutangService struct {
	lunashutangRepository repository.LunasHutangRepository
}

//NewBookService .....
func NewLunasHutangService(lunashutangRepo repository.LunasHutangRepository) LunasHutangService {
	return &lunashutangService{
		lunashutangRepository: lunashutangRepo,
	}
}

func (service *lunashutangService) UpdateLunas(b dto.IsLunasUpdateDTO) entity.Hutang {
	lunashutang := entity.Hutang{}
	err := smapping.FillStruct(&lunashutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.lunashutangRepository.UpdateIsLunasHutang(lunashutang)
	return res
}

func (service *lunashutangService) LunasIsAllowedToEdit(userID string, hutangID uint64) bool {
	b := service.lunashutangRepository.FindLunasHutangByID(hutangID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
