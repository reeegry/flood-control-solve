package floodControl

import (
	"context"
	"github.com/reeegry/flood-control-solve/config"
	"github.com/reeegry/flood-control-solve/internal/db"
	"strconv"
)

type FloodController struct {
	Db *db.Redis
	N  int
	K  int
}

func NewController(rd *db.Redis, cfg config.Config) *FloodController {
	return &FloodController{
		Db: rd,
		K:  cfg.K,
	}
}

func (f *FloodController) Check(ctx context.Context, userID int64) (bool, error) {
	InDB, err := f.Db.Client.Exists(ctx, strconv.Itoa(int(userID))).Result()
	if err != nil {
		return false, err
	}

	if InDB == 0 {
		err := f.Db.AddUser(ctx, userID)
		if err != nil {
			return false, err
		}

	}

	err = f.Db.IncrTime(ctx, userID)
	if err != nil {
		return false, err
	}

	if cnt, _ := f.Db.GetVal(ctx, userID); cnt > f.K {
		return false, nil
	}

	return true, nil
}
