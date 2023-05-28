package routers

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/labstack/echo"

)

func Setup(router *echo.Engine, oracle_db *sqlx.DB, redis_cache *redis.Client) {

	//router.Use(middlewares.NewCorsAccessControl().CorsAccessControl())

	router.GET("/healthz", func(c *echo.Context) {
		return c.JSON(http.StatusOK,"The service works normally... kor jaa")
	})

	PORT := viper.GetString("app.port")
	router.Logger.Fatal(router.Start(PORT))

}
