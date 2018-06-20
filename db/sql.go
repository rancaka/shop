package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/rancaka/shop/config"
)

func NewSQL(cfg *config.Config) (*sqlx.DB, error) {
	return sqlx.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			cfg.DB.SQL.Username,
			cfg.DB.SQL.Password,
			cfg.DB.SQL.Host,
			cfg.DB.SQL.Port,
			cfg.DB.SQL.DatabaseName),
	)
}
