package main

/* TODO:
1 обновлять таймер через контекст
2 через горутины цикл запустить
3 добавлять пользователя сразу а не в чеке
4 сделать readme
5 билдить как расписать
6 почистить код сделать все по красоте
7 рефакторинг
*/
import (
	"context"
	"fmt"
	"strconv"
	"test/db"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	//db.DbConnect()
	rd := db.NewRd()
	controller := NewController(rd)
	fmt.Println(controller)

	//controller.db.AddUser(ctx, 123)

	for i := 0; i < 10; i++ {
		ok, err := controller.Check(ctx, 123)
		if err != nil {
			fmt.Println("err: ", err)
		}

		if ok {
			fmt.Println("OK")
		} else {
			fmt.Println("NOT OK")
			//controller.db.AddUser(ctx, 123)
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

type FloodController struct {
	db *db.Redis
	N  int
	K  int
}

func NewController(rd *db.Redis) *FloodController {

	return &FloodController{
		db: rd,
		N:  4,
		K:  2,
	}
}

func (f *FloodController) Check(ctx context.Context, userID int64) (bool, error) {
	InDB, err := f.db.Client.Exists(ctx, strconv.Itoa(int(userID))).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(InDB)

	if InDB == 0 {
		// TODO: изменить на addUser
		_, err = f.db.Client.Set(ctx, strconv.Itoa(int(userID)), 0, time.Duration(f.N)*time.Second).Result()

	}

	val, err := f.db.Client.Incr(ctx, strconv.Itoa(int(userID))).Result()
	if err != nil {
		return false, err
	}

	fmt.Println(val)

	if cnt, _ := f.db.Client.Get(ctx, strconv.Itoa(int(userID))).Int(); cnt > f.K {
		return false, nil
	}

	return true, nil
}
