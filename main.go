package main

import (
	"os"

	token "github.com/frank1995alfredo/api/controllers/token"
	rutas "github.com/frank1995alfredo/api/routes"
	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
)

func main() {
	rutas.Rutas()

}
func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	token.Client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := token.Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
