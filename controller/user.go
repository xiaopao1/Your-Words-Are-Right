package controller

import (
	"Your_Words_Are_Right/model/database"
	"Your_Words_Are_Right/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var usermodel database.User

	username := c.Query("username")
	password := c.Query("password")

	if !usermodel.CheckUsernameIsExist(username) {
		c.JSON(http.StatusOK, response.UserLoginResponse{
			Response: response.Response{
				StatusCode: 1, StatusMsg: "username has exist",
			},
			UserId: 0,
			Token:  "",
		})
	} else {
		token, userId := usermodel.CreateUser(username, password)
		c.JSON(http.StatusOK, response.UserLoginResponse{
			Response: response.Response{
				0, "success",
			},
			UserId: userId,
			Token:  token,
		})
	}
}
func Login(c *gin.Context) {

}
