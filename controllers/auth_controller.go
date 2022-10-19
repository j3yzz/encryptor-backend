package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/j3yzz/encryptor/common"
	"github.com/j3yzz/encryptor/database"
	"github.com/j3yzz/encryptor/services/actions"
	"github.com/j3yzz/encryptor/services/models"
	"github.com/j3yzz/encryptor/validators"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(requestUser)
	fmt.Println(requestUser)
	responseStatus, token := actions.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func Register(context *gin.Context) {
	var user models.User

	userValidator := validators.NewUserValidator()
	if err := userValidator.Bind(context); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	if err := context.ShouldBindJSON(&user); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-2",
				"message": err.Error()})
		return
	}

	if err := user.HashedPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"id": user.Id, "email": user.Email, "username": user.Username})
}
