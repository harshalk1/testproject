package models

import (
	"github.com/go-redis/redis"
	//"github.com/couchbase/go-couchbase"
	"github.com/astaxie/beego"
	//"reflect"
)

func initDb() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redisHost"),
		Password: "", // no password set
		DB:       0,  // use default DB
		//DialTimeout: 50000,
	})
	return client
}

var client = initDb()

//RedisGet used to get doc from redis db
func RedisGet(key string) (string, error) {
	val, err := client.Get(key).Result()
	//fmt.Println("client", client, "\nErr:-->", err, "\nval--->", val)
	return val, err
}

//RedisSet used to set doc in redis db
func RedisSet(key string, val string, expTime int) (string, error) {
	val, err := client.Set(key, val, 0).Result()
	return val, err
}
