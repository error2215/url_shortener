package url

import (
	"context"
	"github.com/error2215/url_shortener/server/cache"
	db "github.com/error2215/url_shortener/server/db/mongo/url"
	"github.com/sirupsen/logrus"
)

func FindShortenedUrl(ctx context.Context, url string) (shortened string, err error) {
	shortened, err = db.FindShortenedUrl(ctx, url)
	if err != nil {
		return "", err
	}
	return shortened, nil
}

func CreateShortenedUrl(ctx context.Context, url string) (shortened string, err error) {
	shortened, err = db.CreateShortenedUrl(ctx, url)
	if err != nil {
		return "", err
	}
	return shortened, nil
}

func FindRealUrl(ctx context.Context, shortened string) (realUrl string, err error) {
	if realUrl, found := cache.GetCache().Get(shortened); found == true {
		logrus.Info("Found url in cache")
		return realUrl, nil
	}
	realUrl, err = db.FindRealUrl(ctx, shortened)
	if err != nil {
		return "", err
	}
	if realUrl == "" {
		return "", nil
	}
	logrus.Info("Found url in database")
	cache.GetCache().Set(shortened, realUrl, 0)
	return realUrl, nil
}
