package actions

import (
	"encoding/json"
	"github.com/j3yzz/encryptor/api/parameters"
	"github.com/j3yzz/encryptor/core/authentication"
	"github.com/j3yzz/encryptor/services/models"
	"net/http"
)

func Login(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthentication()

	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.Id)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthentication{Token: token})
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}
