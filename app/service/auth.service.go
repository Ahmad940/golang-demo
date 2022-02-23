package service

import (
	"demo/pkg/util"
	"demo/postgres"
	"fmt"
)

func RegisterUser(user postgres.InsertUserParams) error {
	InsertUser(user)

	return fmt.Errorf("")
}

func Login(email, password string) (string, error) {
	user, _ := GetUserByEmail(email)
	if password := util.CheckPasswordHash(password, user.Password); password == true {
		println("Correct password")
		token, err := util.GenerateToken(user.ID.String())
		if err != nil {
			return "", err
		}
		fmt.Printf("Token %v", token)
		return token, err
	} else {
		println("Wrong password")
		return "", fmt.Errorf("invalid email or password")
	}
}
