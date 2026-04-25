package utils

import (
	"fmt"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func ExecuteMethods(obj any) {
	routeValue := reflect.ValueOf(obj)
	for i := range routeValue.NumMethod() {
		routeValue.Method(i).Call(nil)
		fmt.Println("Each Method Registered")
	}
}
