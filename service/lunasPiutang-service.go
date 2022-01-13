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
type LunasPiutangService interface {
	UpdateLunasPiutang(b dto.IsLunasPiutangUpdateDTO) entity.Piutang
	LunasPiutangIsAllowedToEdit(userID string, piutangID uint64) bool
}

type lunaspiutangService struct {
	lunaspiutangRepository repository.LunasPiutangRepository
}

//NewBookService .....
func NewLunasPiutangService(lunaspiutangRepo repository.LunasPiutangRepository) LunasPiutangService {
	return &lunaspiutangService{
		lunaspiutangRepository: lunaspiutangRepo,
	}
}

func (service *lunaspiutangService) UpdateLunasPiutang(b dto.IsLunasPiutangUpdateDTO) entity.Piutang {
	lunaspiutang := entity.Piutang{}
	err := smapping.FillStruct(&lunaspiutang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.lunaspiutangRepository.UpdateIsLunasPiutang(lunaspiutang)
	return res
}

func (service *lunaspiutangService) LunasPiutangIsAllowedToEdit(userID string, putangID uint64) bool {
	b := service.lunaspiutangRepository.FindLunasPiutangByID(putangID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
