package db

import (
    "context"
    "log"
    "time"

    "github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", 
        Password: "",               
        DB:       0,                
    })

    if err := RedisClient.Ping(ctx).Err(); err != nil {
        return err
    }

    log.Println("✅ اتصال به Redis برقرار شد")
    return nil
}
