package main

import (
	"fmt"

	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/utils"
)

func main() {
	password := "Password@123"

	hash, err := utils.HashPassword(password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hashed Password:")
	fmt.Println(hash)

	fmt.Println("\nPassword Match:")
	fmt.Println(utils.CheckPassword(password, hash))
}