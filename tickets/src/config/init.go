package config

import (
	"log"
	"os"
)

type Config struct {
	MongoUri string
	DatabaseName string
}

func Init() *Config {
	conf := &Config{}

	conf.MongoUri = os.Getenv("MONGO_URI")
	if conf.MongoUri == "" {
		log.Fatal("MONGO_URI is not defined")
	}

	conf.DatabaseName = os.Getenv("DATABASE_NAME")
	if conf.DatabaseName == "" {
		log.Fatal("DATABASE_NAME is not defined")
	}

	return conf
}