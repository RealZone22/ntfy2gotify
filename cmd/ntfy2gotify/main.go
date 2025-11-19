package ntfy2gotify

import (
	"ntfy2gotify/cmd/api"
	"ntfy2gotify/pkg/utils"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	utils.Logger.Info().Msg("Initializing...")

	go api.StartRestAPI()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	utils.Logger.Info().Msg("Shutting down...")
}
