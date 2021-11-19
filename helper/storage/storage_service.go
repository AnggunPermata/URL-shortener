package storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

//Redis Client
type StorageService struct {
	redisClient *redis.Client
}

var (
	storageService = &StorageService{}
	ctx = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storageService.redisClient = redisClient
	return storageService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string)error{
	if err := storageService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err(); err != nil{
		return err
		//panic(fmt.Sprintf("failed saving key url | error : %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
	return nil
}

func RetrieveInitialUrl(shortUrl string) (string, error){
	result, err := storageService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		return result, err
		//panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result, nil
}