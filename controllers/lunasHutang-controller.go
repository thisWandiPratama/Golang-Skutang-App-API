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
type LunasHutangController interface {
	UpdateIsLunasHutang(context *gin.Context)
}

type lunashutangController struct {
	lunasHutangService service.LunasHutangService
	jwtService         service.JWTService
}

//NewBookController create a new instances of BoookController
func NewLunasHutangController(lunasHutangServ service.LunasHutangService, jwtServ service.JWTService) LunasHutangController {
	return &lunashutangController{
		lunasHutangService: lunasHutangServ,
		jwtService:         jwtServ,
	}
}

func (c *lunashutangController) UpdateIsLunasHutang(context *gin.Context) {
	var isLunasHutangUpdateDTO dto.IsLunasUpdateDTO
	// fmt.Println("con", isLunasHutangUpdateDTO.Islunas)

	errDTO := context.ShouldBind(&isLunasHutangUpdateDTO)
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
	if c.lunasHutangService.LunasIsAllowedToEdit(userID, isLunasHutangUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			isLunasHutangUpdateDTO.UserID = id
		}
		result := c.lunasHutangService.UpdateLunas(isLunasHutangUpdateDTO)
		// fmt.Println("result", result)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
