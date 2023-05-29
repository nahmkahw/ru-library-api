package databases

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type connection struct{}

func NewDatabases() *connection {
	return &connection{}
}

func (c *connection) OracleInit() *sqlx.DB {

	db, err := oracleConnection()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)
	return db
}

func (c *connection) RedisInint() *redis.Client {
	return redisConnection()
}

func redisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		//Addr: viper.GetString("redis_cache.addressInLocal"),
		Addr: viper.GetString("redis_cache.address"),
		// Addr:     viper.GetString("redis_cache.addressInServer"),
		Password: viper.GetString("redis_cache.password"),
		DB:       viper.GetInt("redis_cache.db-num"),
	})
}

func oracleConnection() (*sqlx.DB, error) {

	dns := fmt.Sprintf("%v", viper.GetString("db.connection"))
	driver := viper.GetString("db.openDriver")

	return sqlx.Open(driver, dns)

}

func (c *connection) PostgresInit() *sqlx.DB {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", viper.GetString("postgres.host"), viper.GetInt("postgres.port"), viper.GetString("postgres.user"), viper.GetString("postgres.password"), viper.GetString("postgres.dbname"))

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	return db
}
