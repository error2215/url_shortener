package config

import (
	"github.com/error2215/go-convert"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ApiPort                    string
	MongoPort                  string
	DefaultCacheExpirationTime int32 //seconds
	CleanupInterval            int32 //seconds
}

var GlobalConfig Config

func init() {
	if os.Getenv("TESTING") != "true" {
		if err := godotenv.Load(); err != nil {
			log.Warn("Error loading .env file")
		}
		GlobalConfig = Config{
			ApiPort:                    os.Getenv("API_PORT"),
			MongoPort:                  os.Getenv("MONGO_PORT"),
			DefaultCacheExpirationTime: convert.Int32(os.Getenv("DEFAULT_CACHE_EXPIRATION_TIME")),
			CleanupInterval:            convert.Int32(os.Getenv("CLEANUP_INTERVAL")),
		}
	}
}
