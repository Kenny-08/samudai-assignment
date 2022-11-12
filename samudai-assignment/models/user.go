package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id	int	`json:"ID" gorm:"primary_key"`
	Name	string	`json:"name" gorm:"not null"`
	Email	string	`json:"email" gorm:"not null;unique"`
	Password	string	`json:"password" gorm:"not null"`
}

func (user *User) SaveUser() (*User, error){
	fmt.Println("Inside save user function")
	var err error
	err = DB.Create(&user).Error
	if err!=nil{
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil

}