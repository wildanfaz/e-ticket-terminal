package pkg

import "golang.org/x/crypto/bcrypt"

type Login struct {
	Token string `json:"token"`
}

func NewLogin(token string) Login {
	return Login{
		Token: token,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
