package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"


	"github.com/juju/loggo"
)

var logger *loggo.Logger

func main() {
	loggo.ConfigureLoggers("<root>=TRACE")
	newLogger :=  loggo.GetLogger("vesta")
	logger = &newLogger

	//config := CollectConfig()

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	nch := make(chan os.Signal)
	signal.Notify(nch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-nch)

	logger.Infof("Done!")
}

