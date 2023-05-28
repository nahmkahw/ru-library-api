package main

import (
	"ru-library-api/environments"
	"ru-library-api/databases"
	"ru-library-api/routers"

		"github.com/go-redis/redis/v8"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo"
)

var oracle_db *sqlx.DB
var redis_cache *redis.Client

func init(){
	environments.TimeZoneInit()
	environments.EnvironmentInit()
	oracle_db = databases.NewDatabases().OracleInit()
	redis_cache = databases.NewDatabases().RedisInint()
} 

func main() {
	// Create a new Fiber app
	defer oracle_db.Close()
	defer redis_cache.Close()
	router := echo.New()
	routers.Setup(router, oracle_db, redis_cache)
}
