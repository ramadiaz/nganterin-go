package helpers

import (
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) (dto.User, *exceptions.Exception){
	var result dto.User
	user_data, _ := c.Get("user")

	result, ok := user_data.(dto.User)
	if !ok {
		return result, exceptions.NewException(http.StatusUnauthorized, exceptions.ErrInvalidTokenStructure)
	}
	

	return result, nil
}