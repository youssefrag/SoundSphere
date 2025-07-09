package cache

import (
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (

  Ctx = context.Background()
  Client *redis.Client
)

func Init() {
  host := os.Getenv("REDIS_HOST")   // e.g. "redis"
  port := os.Getenv("REDIS_PORT")   // e.g. "6379"
  pass := os.Getenv("REDIS_PASSWORD")
  dbNum := 0
  if v := os.Getenv("REDIS_DB"); v != "" {
    if i, err := strconv.Atoi(v); err == nil {
      dbNum = i
    }
  }

  Client = redis.NewClient(&redis.Options{
    Addr:     host + ":" + port,
    Password: pass,
    DB:       dbNum,
  })
}
