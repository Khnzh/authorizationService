package models

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type ServiceApis struct {
	DB    *sql.DB
	Redis *redis.Client
}
