package handlers

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func JWTIsValid(userId, tokenId string) bool {
	fmt.Println("accessed JWTIsValid")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	// err := rdb.Set(ctx, "user:" + userId + ":" + tokenId, "valid", 10).Err()
	//
	// if err != nil {
	// 	panic(err)
	// }

	results, err := rdb.Keys(ctx, "user:"+userId+":"+tokenId).Result()

	if err != nil {
		panic(err)
	}

	if len(results) < 1 {
		return false
	}

	return true
}

func JWTSet(userId, tokenId string) {
	fmt.Println("accessed JWTIsValid")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	err := rdb.Set(ctx, "user:"+userId+":"+tokenId, "valid", 0).Err()

	if err != nil {
		panic(err)
	}

}
