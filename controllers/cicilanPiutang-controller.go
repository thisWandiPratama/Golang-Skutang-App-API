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
type CicilanPiutangController interface {
	UpdateCicilanPiutang(context *gin.Context)
}

type cicilanpiutangController struct {
	cicilanPiutangService service.CicilanPiutangService
	jwtService            service.JWTService
}

//NewBookController create a new instances of BoookController
func NewCicilanPiutangController(cicilanPiutangServ service.CicilanPiutangService, jwtServ service.JWTService) CicilanPiutangController {
	return &cicilanpiutangController{
		cicilanPiutangService: cicilanPiutangServ,
		jwtService:            jwtServ,
	}
}

func (c *cicilanpiutangController) UpdateCicilanPiutang(context *gin.Context) {
	var cicilanPiutangUpdateDTO dto.CicilanUpdatePiutangDTO
	errDTO := context.ShouldBind(&cicilanPiutangUpdateDTO)
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
	if c.cicilanPiutangService.CicilanPiutangAllowedToEdit(userID, cicilanPiutangUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			cicilanPiutangUpdateDTO.UserID = id
		}
		result := c.cicilanPiutangService.UpdateCicilanPiutang(cicilanPiutangUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
