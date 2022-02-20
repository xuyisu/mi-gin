package session

import (
	"go-gin/models"
	"time"
)

func Set(key string, value interface{}, expire int) {
	models.RedisDB.Set(key, value, time.Second*time.Duration(expire))
}

func Get(key string) string {
	cmd := models.RedisDB.Get(key)
	return cmd.Val()
}

func Delete(key string) {
	models.RedisDB.Del(key)
}
