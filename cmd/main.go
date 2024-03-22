package main

import (
	"context"
	"fmt"
	"github.com/reeegry/flood-control-solve/config"
	"github.com/reeegry/flood-control-solve/internal/db"
	"github.com/reeegry/flood-control-solve/internal/floodControl"
	"log"
	"math/rand"
	"time"
)

func messageSimulate(ctx context.Context, cfg config.Config, c *floodControl.FloodController) {
	for {
		ind := rand.Int() % len(cfg.UserIds)
		for i := 0; i < 3; i++ {
			if rand.Float64() > 0.5 {
				ok, err := c.Check(ctx, cfg.UserIds[ind])
				if err != nil {
					log.Printf("can't check user %d", cfg.UserIds[ind])
				}
				val, err := c.Db.GetVal(ctx, cfg.UserIds[ind])
				log.Printf("user %d %d messages", cfg.UserIds[ind], val)
				if !ok {
					log.Printf("user %d if fludding", cfg.UserIds[ind])
				}
			}
			time.Sleep(time.Second)
		}
	}
}

func main() {
	cfg, _ := config.Read("../config/config.json")
	fmt.Println(cfg)

	ctx := context.Background()

	rd := db.NewRd(cfg)
	controller := floodControl.NewController(rd, cfg)
	fmt.Println(controller)

	messageSimulate(ctx, cfg, controller)

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
