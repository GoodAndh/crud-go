package invalid

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func SetPassword(password string) []byte {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("a")

		return nil
	}
	return hashpassword
}

// input for hashed password ,password for client
func CheckPassword(input string, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(input), []byte(password))
	if err != nil {
		fmt.Println("b")
		return err
	}
	return nil
}
