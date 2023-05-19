package main

import (
	"fmt"
	"payme-go/internal/config"
)

func main() {
	cnf, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}
	fmt.Println(cnf.Server.Port)
}
