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
type CicilanHutangController interface {
	UpdateCicilanHutang(context *gin.Context)
}

type cicilanhutangController struct {
	cicilanHutangService service.CicilanHutangService
	jwtService           service.JWTService
}

//NewBookController create a new instances of BoookController
func NewCicilanHutangController(cicilanHutangServ service.CicilanHutangService, jwtServ service.JWTService) CicilanHutangController {
	return &cicilanhutangController{
		cicilanHutangService: cicilanHutangServ,
		jwtService:           jwtServ,
	}
}

func (c *cicilanhutangController) UpdateCicilanHutang(context *gin.Context) {
	var cicilanHutangUpdateDTO dto.CicilanUpdateDTO
	errDTO := context.ShouldBind(&cicilanHutangUpdateDTO)
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
	if c.cicilanHutangService.CicilanAllowedToEdit(userID, cicilanHutangUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			cicilanHutangUpdateDTO.UserID = id
		}
		result := c.cicilanHutangService.UpdateCicilan(cicilanHutangUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
