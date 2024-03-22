package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/reeegry/flood-control-solve/config"
	"strconv"
	"time"
)

type Redis struct {
	Client *redis.Client
	N      int
}

func NewRd(cfg config.Config) *Redis {
	serv := cfg.Redis.Host + ":" + strconv.Itoa(cfg.Redis.Port)
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     serv,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
		N: cfg.N,
	}

}

func (r *Redis) AddUser(ctx context.Context, UserID int64) error {
	_, err := r.Client.Set(ctx, strconv.Itoa(int(UserID)), 0, time.Duration(r.N)*time.Second).Result()
	return err
}

func (r *Redis) IncrTime(ctx context.Context, UserID int64) error {
	_, err := r.Client.Incr(ctx, strconv.Itoa(int(UserID))).Result()
	return err
}

func (r *Redis) GetVal(ctx context.Context, UserID int64) (int, error) {
	val, err := r.Client.Get(ctx, strconv.Itoa(int(UserID))).Int()
	return val, err
}
