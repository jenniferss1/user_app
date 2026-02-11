package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetWeight() float64
	GetHeight() float64

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
	weight, height float64,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age, weight, height,
	}
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) GetWeight() float64 {
	return ud.weight
}

func (ud *userDomain) GetHeight() float64 {
	return ud.height
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
	weight   float64
	height   float64
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
