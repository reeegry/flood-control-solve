package main

import (
	"context"
	"fmt"

	"github.com/reeegry/flood-control-solve/config"
	"github.com/reeegry/flood-control-solve/internal/db"
	"github.com/reeegry/flood-control-solve/internal/floodControl"
	"time"
)

func main() {
	cfg, _ := config.Read("../config/config.json")
	fmt.Println(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*11)
	defer cancel()

	rd := db.NewRd(cfg)
	controller := floodControl.NewController(rd, cfg)
	fmt.Println(controller)

	for i := 0; i < 25; i++ {
		ok, err := controller.Check(ctx, 123)
		if err != nil {
			fmt.Println("err: ", err)
		}

		if !ok {
			fmt.Printf("%d is fluding", 123)
		}

		time.Sleep(time.Second)
	}

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
