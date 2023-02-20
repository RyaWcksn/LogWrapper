package main

import (
	"bitbucket.org/ayopop/ct-logger/constant"
	"bitbucket.org/ayopop/ct-logger/logger"
)

func main() {
	log := logger.New("ct-logger", constant.EnvDev, "debug")
	log.Debug("Message")
}
