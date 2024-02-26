package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Realm                string `json:"Realm"`
	IdentityKey          string `json:"IdentityKey"`
	PublicAccess         string `json:"PublicAccess"`
	IsAuthenticatedFully string `json:"IsAuthenticatedFully"`
	TokenLookup          string `json:"TokenLookup"`
	TokenHeadName        string `json:"TokenHeadName"`
	DefaultPort          string `json:"DefaultPort"`
	DatabaseDSN          string `json:"DatabaseDSN"`
	JwtTestMail          string `json:"JwtTestMail"`
	JwtTestPass          string `json:"JwtTestPass"`
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
