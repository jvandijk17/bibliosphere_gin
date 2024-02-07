package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Realm                string `json:"realm"`
	IdentityKey          string `json:"identityKey"`
	PublicAccess         string `json:"publicAccess"`
	IsAuthenticatedFully string `json:"isAuthenticatedFully"`
	TokenLookup          string `json:"tokenLookup"`
	TokenHeadName        string `json:"tokenHeadName"`
	DefaultPort          string `json:"defaultPort"`
	DatabaseDSN          string `json:"databaseDSN"`
}

var AppConfig Config

func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	file, err := os.Open(fmt.Sprintf("config/config.%s.json", env))
	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		log.Fatalf("Cannot decode config JSON: %v", err)
	}

}
