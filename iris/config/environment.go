package config

import (
	"os"
	"strings"

	"github.com/gofor-little/env"
	"github.com/spf13/viper"
)

func IsModeDev() bool {
	command := os.Args[0]
	if strings.Contains(command, "exe") { // go run server.go
		return true
	} else {
		return false
	}

}

func LoadEnv() {
	const (
		EnvDevFile  = ".env.development"
		EnvProdFile = ".env.production"
	)
	// Load the .env.dev file
	if IsModeDev() {
		err := env.Load(EnvDevFile)
		if err != nil {
			panic(err)
		}
	} else {
		err := env.Load(EnvProdFile)
		if err != nil {
			panic(err)
		}
	}
}

func GetEnv(key string) string {
	return env.Get(key, "")
}

func GetVipper(key string) string {
	return viper.GetString(key)
}
