package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Folders map[string]string `json:"folders"`
}

var AppConfig Config

func LoadConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	return decoder.Decode(&AppConfig)
}
