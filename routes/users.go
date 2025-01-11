package routes

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"eventBooking.com/event-booking/models"
	"eventBooking.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	body, _ := io.ReadAll(context.Request.Body)
	fmt.Println("Raw Request Body:", string(body))
	context.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset the body for further use

	if err != nil {
		fmt.Println("Error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		fmt.Println("Error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Login successfull", "token": token})
}
