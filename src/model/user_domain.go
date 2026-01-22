package model

import (
	"crypto/md5"
	"encoding/hex"
	"go_app/src/configuration/rest_err"
)

func NewUserDomain(
	email, password, name string,
	age int8,
	weight, height float64,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age, weight, height,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
	Weight   float64
	Height   float64
}

type UserDomainInterface interface {
	CreateUser() *rest_err.Resterr
	UpdateUser(string) *rest_err.Resterr
	FindUser(string) (*UserDomain, *rest_err.Resterr)
	DeleteUser(string) *rest_err.Resterr
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
