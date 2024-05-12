package db

import (
	"log"
	"runtime"

	C "github.com/eleynes/MyK3y/config"
	"github.com/gofiber/storage/redis/v3"
)

func RedisConnect() (storage *redis.Storage, err error) {

	// Initialize default config
	// store := redis.New()

	// Initialize custom config
	store := redis.New(redis.Config{
		Host:      C.Conf.RedisHost,
		Port:      C.Conf.RedisPort,
		Username:  C.Conf.RedisUsername,
		Password:  C.Conf.RedisPassword,
		Database:  C.Conf.RedisDatabase,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})

	log.Println("RedisDB Connected")
	return store, nil

	// // Initialize Redis Failover Client
	// store := redis.New(redis.Config{
	// 	MasterName: "master-name",
	// 	Addrs:      []string{":6379"},
	// })

	// // Initialize Redis Cluster Client
	// store := redis.New(redis.Config{
	// 	Addrs: []string{":6379", ":6380"},
	// })

	// // Create a client with support for TLS
	// cer, err := tls.LoadX509KeyPair("./client.crt", "./client.key")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// tlsCfg := &tls.Config{
	// 	MinVersion:         tls.VersionTLS12,
	// 	InsecureSkipVerify: true,
	// 	Certificates:       []tls.Certificate{cer},
	// }
	// store = redis.New(redis.Config{
	// 	URL:       "redis://<user>:<pass>@127.0.0.1:6379/<db>",
	// 	TLSConfig: tlsCfg,
	// 	Reset:     false,
	// })

	// Create a client with a Redis URL with all information.
	// store = redis.New(redis.Config{
	// 	URL:   "redis://<user>:<pass>@127.0.0.1:6379/<db>",
	// 	Reset: false,
	// })
}
