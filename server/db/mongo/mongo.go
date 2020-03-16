package mongo

import (
	"context"
	"github.com/error2215/url_shortener/server/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var client *mongo.Client

func init() {
	if os.Getenv("TESTING") != "true" {
		var err error
		link := "mongodb://127.0.0.1:"
		if os.Getenv("DOCKER_COMPOSE") == "true" {
			link = "mongodb://mongo:"
		}
		client, err = mongo.NewClient(options.Client().ApplyURI(link + config.GlobalConfig.MongoPort))
		if err != nil {
			log.WithField("method", "server.db.mongo.init").Fatal(err)
		}

		err = client.Connect(context.TODO())
		if err != nil {
			log.WithField("method", "server.db.mongo.init").Fatal(err)
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.WithField("method", "server.db.mongo.init").Fatal(err)
		}
		log.Info("Connection to MongoDB finished. Address: " + config.GlobalConfig.MongoPort)
	}
}

func GetClient() *mongo.Client {
	return client
}
