package main

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var redisStorage *redis.Storage
var sessStore *session.Store

func setSessionStorage() {
	redisStorage = redis.New(redis.Config{
		URL: os.Getenv("REDIS_URL"),
	})
	sessStore = session.New(session.Config{
		Storage:      redisStorage,
		CookieDomain: "http://localhost:3000",
	})
}

func getSessionStore() *session.Store {
	if sessStore == nil {
		setSessionStorage()
	}
	return sessStore
}
