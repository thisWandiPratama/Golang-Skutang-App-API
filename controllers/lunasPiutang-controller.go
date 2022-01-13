package controllers

import (
	"fmt"
	"golang_api_hupiutang/dto"
	"golang_api_hupiutang/helper"
	"golang_api_hupiutang/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//BookController is a ...
type LunasPiutangController interface {
	UpdateIsLunasPiutang(context *gin.Context)
}

type lunaspiutangController struct {
	lunasPiutangService service.LunasPiutangService
	jwtService          service.JWTService
}

//NewBookController create a new instances of BoookController
func NewLunasPiutangController(lunasPiutangServ service.LunasPiutangService, jwtServ service.JWTService) LunasPiutangController {
	return &lunaspiutangController{
		lunasPiutangService: lunasPiutangServ,
		jwtService:          jwtServ,
	}
}

func (c *lunaspiutangController) UpdateIsLunasPiutang(context *gin.Context) {
	var isLunasPiutangUpdateDTO dto.IsLunasPiutangUpdateDTO
	// fmt.Println("con", isLunasHutangUpdateDTO.Islunas)

	errDTO := context.ShouldBind(&isLunasPiutangUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.lunasPiutangService.LunasPiutangIsAllowedToEdit(userID, isLunasPiutangUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			isLunasPiutangUpdateDTO.UserID = id
		}
		result := c.lunasPiutangService.UpdateLunasPiutang(isLunasPiutangUpdateDTO)
		// fmt.Println("result", result)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
