package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/j3yzz/encryptor/common"
	"github.com/j3yzz/encryptor/services/models"
)

type UserValidator struct {
	User struct {
		Id       int    `json:"id"`
		Username string `json:"username" gorm:"unique"`
		Name     string `json:"name"`
		Email    string `json:"email" gorm:"unique"`
		Password string `json:"password"`
	} `json:"user"`
	userModel models.User `json:"-"`
}

func (self *UserValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.userModel.Username = self.User.Username
	self.userModel.Email = self.User.Email
	self.userModel.Name = self.User.Name
	if self.User.Password != common.NBRandomPassword {
		self.userModel.HashedPassword(self.User.Password)
	}
	return nil
}

func NewUserValidator() UserValidator {
	userValidator := UserValidator{}
	return userValidator
}
