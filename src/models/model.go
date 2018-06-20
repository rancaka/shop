package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/rancaka/shop/config"
)

type Model struct {
	Cfg *config.Config
	DB  *sqlx.DB
}

func New(cfg *config.Config, db *sqlx.DB) *Model {
	return &Model{Cfg: cfg, DB: db}
}
