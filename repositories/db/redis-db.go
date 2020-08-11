package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/config"
)

type redisService struct {
	Client *redis.Client
}

type RedisRepository interface {
	Send(ctx context.Context, match *api.Match) error
	Get(ctx context.Context, key int) *api.Match
}

func NewRedisCache(conf *config.Config) (RedisRepository, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})

	return &redisService{Client: redisClient}, nil

}

func (r *redisService) Send(ctx context.Context, match *api.Match) error {
	value, err := json.Marshal(match)
	if err != nil {
		log.Println("Could not unmarshal message", err)
	}

	key := strconv.Itoa(match.MatchDTO.GameId)
	fmt.Println("Keyyy: ", key)
	err = r.Client.Set(key, value, 10*time.Second).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func (r *redisService) Get(ctx context.Context, key int) *api.Match {

	value, err := r.Client.Get(strconv.Itoa(key)).Result()
	if err != nil {
		log.Println("Could not get match from redis", err)
		return nil
	}

	match := api.Match{}
	err = json.Unmarshal([]byte(value), &match.MatchDTO)
	if err != nil {
		log.Println("Unable to Unmarshal value from redis")
		return nil
	}
	return &match

}
