package utils

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog"
)

type ConfigStruct struct {
	LogLevel zerolog.Level `json:"log_level"` // 0: Debug, 1: Info, 2: Warn, 3: Error, 4: Fatal, 5: Panic, 6: NoLevel, 7: Disabled
	Api      struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"api"`
	Subscriptions map[string]string `json:"subscriptions"` // e.g. {"ntfy_topic": "gotify_app"}
}

var Config *ConfigStruct

func InitConfig() error {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		if err = createConfig(); err != nil {
			return err
		}
	}

	if err := readConfig(); err != nil {
		return err
	}

	return nil
}

func readConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	Config = &ConfigStruct{}
	err = json.NewDecoder(file).Decode(Config)
	if err != nil {
		return err
	}

	return nil
}

func createConfig() error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	data, err := json.MarshalIndent(&ConfigStruct{}, "", "   ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
