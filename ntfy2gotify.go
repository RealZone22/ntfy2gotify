package main

import (
	"log"
	"math/rand"
	"ntfy2gotify/cmd/ntfy2gotify"
	"ntfy2gotify/pkg/utils"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	err := utils.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	utils.InitLogger()

	utils.Logger.Debug().Msg("Pre-Initialization finished")

	ntfy2gotify.Init()
}
