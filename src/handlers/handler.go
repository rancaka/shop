package handlers

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rancaka/shop/config"
	"github.com/rancaka/shop/src/models"
)

type Handler struct {
	Cfg       *config.Config
	RedisConn redis.Conn

	Model *models.Model
}

func New(cfg *config.Config, db *sqlx.DB, redisConn redis.Conn) *Handler {
	return &Handler{Cfg: cfg, Model: models.New(cfg, db), RedisConn: redisConn}
}

func (h *Handler) Index(c *gin.Context) {
	c.JSON(200, gin.H{"Hellow": "Anca!"})
}
