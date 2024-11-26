package biz_content

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/dgdts/UniversalServer/pkg/loader"

	redisClient "github.com/dgdts/UniversalServer/pkg/redis"
	redis "github.com/redis/go-redis/v9"
)

var cacheKey = "contents"
var cacheMap sync.Map

var contentLoader = loader.ChainFunc[*ContentData]{
	loader.SingleFlightLoader("fromCache", fromCache),
	loader.SingleFlightLoader("fromRedis", fromRedis),
	loader.SingleFlightLoader("fromDB", fromDB),
}

// Init cache and create schedule cron job to sync redis from db
func InitCache() {
	// TODO: Add cron job to sync redis from db periodically
}

func GetContent(ctx context.Context, contentID string) (*ContentData, error) {
	return contentLoader.Load(ctx, contentID)
}

func fromCache(ctx context.Context, contentID string) (*ContentData, error) {
	if content, ok := cacheMap.Load(contentID); ok {
		return content.(*ContentData), nil
	}
	return nil, loader.ErrNext
}

func fromRedis(ctx context.Context, contentID string) (*ContentData, error) {
	content, err := redisClient.GetConnection().Get(ctx, cacheKey+":"+contentID).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, loader.ErrNext
		}
		return nil, err
	}

	var contentData ContentData
	err = json.Unmarshal([]byte(content), &contentData)
	if err != nil {
		return nil, err
	}

	cacheMap.Store(contentID, &contentData)
	return &contentData, nil
}

func fromDB(ctx context.Context, contentID string) (*ContentData, error) {
	content, err := ContentByID(ctx, contentID)
	if err != nil {
		return nil, err
	}

	err = redisClient.GetConnection().Set(ctx, cacheKey+":"+contentID, content, 24*time.Hour).Err()
	if err != nil {
		log.Printf("Failed to set Redis cache: %v", err)
	}

	cacheMap.Store(contentID, content)
	return content, nil
}
