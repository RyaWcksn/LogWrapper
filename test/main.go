package main

import (
	"bitbucket.org/ayopop/ct-logger/constant"
	"bitbucket.org/ayopop/ct-logger/logger"
)

func main() {
	log := logger.New("ct-log", constant.EnvDev, "debug")
	log.Debug("Deep down")
}
