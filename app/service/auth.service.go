package service

import (
	"demo/pkg/util"
	"fmt"
)

func RegisterUser() {

}

func Login(email, password string) error {
	user, err := GetUserByEmail(email)
	if err != nil {
		return err
	}
print("Hello ")
	if password := util.CheckPasswordHash(password, user.Password); password == true {
		print("COrrect passwprd")
	} else {
		print("Wrong password")
		return fmt.Errorf("invalid email or password")
	}


	return err
}
