package main

import (
	"github.com/magafagafa/go-app/route"

	config "github.com/magafagafa/go-app/conf"

	"github.com/sirupsen/logrus"
)

func main() {
	router := route.Init()
	router.StartTLS(":" + config.Config.Port, config.Config.Sslpem, config.Config.Sslkey)

	// router.Start(":" + config.Config.Port)

}
func init() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
