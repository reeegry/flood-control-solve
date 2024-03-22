package floodControl

import (
	"context"
	"github.com/reeegry/flood-control-solve/config"
	"github.com/reeegry/flood-control-solve/internal/db"
	"strconv"
)

type FloodController struct {
	db *db.Redis
	N  int
	K  int
}

func NewController(rd *db.Redis, cfg config.Config) *FloodController {
	return &FloodController{
		db: rd,
		//N:  cfg.N,
		K: cfg.K,
	}
}

func (f *FloodController) Check(ctx context.Context, userID int64) (bool, error) {
	InDB, err := f.db.Client.Exists(ctx, strconv.Itoa(int(userID))).Result()
	if err != nil {
		return false, err
	}

	if InDB == 0 {
		err := f.db.AddUser(ctx, userID)
		if err != nil {
			return false, err
		}

	}

	err = f.db.IncrTime(ctx, userID)
	if err != nil {
		return false, err
	}

	if cnt, _ := f.db.GetVal(ctx, userID); cnt > f.K {
		return false, nil
	}

	return true, nil
}
