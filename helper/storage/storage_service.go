package storage

import (
	"context"
	"fmt"
	"github.com/anggunpermata/url-shortener/config"
	"github.com/go-redis/redis"
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
	opt, err := redis.ParseURL(config.LoadEnv("REDIS_URL"))
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect redis")
	}

	redisClient := redis.NewClient(opt)

	fmt.Println(config.LoadEnv("REDIS_URL \n"))
	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storageService.redisClient = redisClient
	return storageService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string)error{
	if err := storageService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err(); err != nil{
		return err
		//panic(fmt.Sprintf("failed saving key url | error : %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
	return nil
}

func RetrieveInitialUrl(shortUrl string) (string, error){
	result, err := storageService.redisClient.Get(shortUrl).Result()
	if err != nil {
		return result, err
		//panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result, nil
}