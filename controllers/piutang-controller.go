package controllers

import (
	"fmt"
	"golang_api_hupiutang/dto"
	"golang_api_hupiutang/entity"
	"golang_api_hupiutang/helper"
	"golang_api_hupiutang/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//BookController is a ...
type PiutangController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type piutangController struct {
	piutangService service.PiutangService
	jwtService     service.JWTService
}

//NewBookController create a new instances of BoookController
func NewPiutangController(piutangServ service.PiutangService, jwtServ service.JWTService) PiutangController {
	return &piutangController{
		piutangService: piutangServ,
		jwtService:     jwtServ,
	}
}

func (c *piutangController) All(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	fmt.Println(id)
	var piutangs []entity.Piutang = c.piutangService.All(id)
	res := helper.BuildResponse(true, "OK", piutangs)
	context.JSON(http.StatusOK, res)

}

func (c *piutangController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var piutang entity.Piutang = c.piutangService.FindByID(id)
	if (piutang == entity.Piutang{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", piutang)
		context.JSON(http.StatusOK, res)
	}
}

func (c *piutangController) Insert(context *gin.Context) {
	var piutangCreateDTO dto.PiutangCreateDTO
	errDTO := context.ShouldBind(&piutangCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			piutangCreateDTO.UserID = convertedUserID
		}
		result := c.piutangService.Insert(piutangCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *piutangController) Update(context *gin.Context) {
	var piutangUpdateDTO dto.PiutangUpdateDTO
	errDTO := context.ShouldBind(&piutangUpdateDTO)
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
	if c.piutangService.IsAllowedToEdit(userID, piutangUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			piutangUpdateDTO.UserID = id
		}
		result := c.piutangService.Update(piutangUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *piutangController) Delete(context *gin.Context) {
	var piutang entity.Piutang
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	piutang.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.piutangService.IsAllowedToEdit(userID, piutang.ID) {
		c.piutangService.Delete(piutang)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *piutangController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
