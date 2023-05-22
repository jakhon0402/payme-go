package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"payme-go/internal/config"
	"payme-go/pkg/logging"
)

func main() {
	cnf, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}
	fmt.Println(cnf.Server.Port)

	log := logging.NewLogger()
	log.Info("Salom")
	log.Warning("Qanday uzi kurinmaysan!")
	log.WithFields(logrus.Fields{"Hola": "Bir bima"}).Info("fdsfdsdsg")
	log.WithFields(logrus.Fields{"Hola": "Bir bima"}).Error("Bu error")

}
