package libs

import (
    "github.com/go-redis/redis"
    "github.com/astaxie/beego"
    "time"
)

var Redis *redis.Client

func init() {
    host := beego.AppConfig.String("redis.host")
    port := beego.AppConfig.String("redis.port")
    db, _ := beego.AppConfig.Int("redis.db")
    password := beego.AppConfig.String("redis.password")

    Redis = redis.NewClient(&redis.Options{
        Addr:         host + ":" + port,
        Password:     password, // no password set
        DB:           db,       // use default DB
        DialTimeout:  10 * time.Second,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        PoolSize:     10,
        PoolTimeout:  30 * time.Second,
    })
    Redis.FlushDB()
}
