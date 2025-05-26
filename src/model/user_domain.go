package model

import (
	"github.com/Mateus003/user-api/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomain(email string, password string, name string, age int8) *UserDomain {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}
func (ud *UserDomain) EncryptPassword() *rest_err.Rest_Err {
	hashed, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		return rest_err.NewInternalServerError("Failed to encrypt password")
	}
	ud.Password = string(hashed)
	return nil
}

type UserDomainInterface interface {
	CreateUser() *rest_err.Rest_Err
	UpdateUser(string,UserDomain) *rest_err.Rest_Err
	FindUser(string) (*UserDomain, *rest_err.Rest_Err)
	DeleteUser(string) *rest_err.Rest_Err
}