package main

import (
	"bitbucket.org/ayopop/ct-logger/constant"
	"bitbucket.org/ayopop/ct-logger/logger"
)

func main() {
	log := logger.New("ct", constant.EnvDev, "debug")
	log.Debug("Mashok")
	iniTest(log)
}

// iniTest ...
func iniTest(log logger.ILogger) {
	log.Debug("Masuk")
}
