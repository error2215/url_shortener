package cache

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/error2215/url_shortener/server/config"
)

var globalCache *cache

func GetCache() *cache {
	return globalCache
}

type cache struct {
	mutex             sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]Item
}

type Item struct {
	Value      string
	Created    time.Time
	Expiration int64
}

func init() {

	items := make(map[string]Item)

	cache := cache{
		items:             items,
		defaultExpiration: time.Second * time.Duration(config.GlobalConfig.DefaultCacheExpirationTime),
		cleanupInterval:   time.Second * time.Duration(config.GlobalConfig.CleanupInterval),
	}
	log.Info("Cache client started: defaultExpiration: ", cache.defaultExpiration, "; cleanupInterval: ", cache.cleanupInterval)
	if cache.cleanupInterval > 0 {
		cache.startGC()
	}

	globalCache = &cache
}

func (c *cache) Set(key string, value string, duration time.Duration) {

	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.mutex.Lock()

	defer c.mutex.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}

}

func (c *cache) Get(key string) (string, bool) {

	c.mutex.RLock()

	defer c.mutex.RUnlock()

	item, found := c.items[key]

	if !found {
		return "", false
	}

	if item.Expiration > 0 {

		if time.Now().UnixNano() > item.Expiration {
			return "", false
		}

	}

	return item.Value, true
}

func (c *cache) startGC() {
	go c.gC()
}

func (c *cache) gC() {

	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)

		}

	}

}

func (c *cache) expiredKeys() (keys []string) {

	c.mutex.RLock()

	defer c.mutex.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

func (c *cache) clearItems(keys []string) {

	c.mutex.Lock()

	defer c.mutex.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
