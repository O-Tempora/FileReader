package main

import (
	"freader/cmd"
	"freader/logger"
	"os"
)

func main() {
	logfile, err := os.OpenFile(
		"logs/logs.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0766,
	)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	logger.InitLogger(logfile)
	//defer logger.Logger.Info().Msg("Application stopped")

	cmd.Execute()
}
