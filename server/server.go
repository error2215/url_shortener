package server

import (
	"context"
	"github.com/error2215/url_shortener/server/api"
	"github.com/error2215/url_shortener/server/api/rest"
	"github.com/error2215/url_shortener/server/config"
	"github.com/error2215/url_shortener/server/db/mongo"
	log "github.com/sirupsen/logrus"
	"sync"
)

func Start() {
	defer func() {
		err := mongo.GetClient().Disconnect(context.TODO())

		if err != nil {
			log.Fatal(err)
		}
		log.Info("Connection to MongoDB closed.")
	}()
	apiPort := config.GlobalConfig.ApiPort
	log.WithFields(log.Fields{
		"apiPort": apiPort,
	}).Info("Launching API server")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		g := &rest.Server{}
		start(g)
	}()

	wg.Wait()
}

func start(api api.API) {
	api.Start()
}
