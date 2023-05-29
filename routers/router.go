package routers

import (
	"net/http"
	"ru-library-api/entity"
	"ru-library-api/handler"
	"ru-library-api/repository"
	"ru-library-api/service"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func Setup(router *echo.Echo, postgres_db *sqlx.DB, redis_cache *redis.Client) {

	//router.Use(middlewares.NewCorsAccessControl().CorsAccessControl())

	router.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, entity.Response{
			Code:         "healthz",
			Data:         "ระบบกำลังทำงาน.",
			HttpCode:     http.StatusOK,
			ErrorMessage: "",
		})
	})

	sierra := router.Group("/sierra")
	sierrarepo := repository.NewSierraRepo(postgres_db)
	sierraservice := service.NewSierraService(sierrarepo, redis_cache)
	sierrahandler := handler.NewSierraHandler(sierraservice)
	sierra.GET("/fine/:id", sierrahandler.Fine)
	sierra.GET("/patron/:id", sierrahandler.Patron)

	PORT := viper.GetString("app.port")
	router.Logger.Fatal(router.Start(PORT))

}
