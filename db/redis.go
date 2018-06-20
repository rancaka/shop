package db

import (
	"github.com/garyburd/redigo/redis"
	"github.com/rancaka/shop/config"
)

func NewRedis(cfg *config.Config) (redis.Conn, error) {
	return redis.Dial("tcp", cfg.DB.Redis.Address)
}
