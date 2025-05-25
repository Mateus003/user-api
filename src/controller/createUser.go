package controller

import (
	"net/http"
	"github.com/Mateus003/user-api/src/configuration/validation"
	"github.com/Mateus003/user-api/src/model/requests"
	"github.com/Mateus003/user-api/src/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var userRequest requests.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		restError := validation.ValidateUserErrors(err)
		c.JSON(restError.Code, restError)

		return
	}


	response := response.UserResponse{
		ID: "test",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	c.JSON(http.StatusOK, response)


}