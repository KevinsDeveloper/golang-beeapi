package libs

import (
    "github.com/go-redis/redis"
    "github.com/astaxie/beego"
)

func Redis() *redis.Client {
    host := beego.AppConfig.String("redis.host")
    port := beego.AppConfig.String("redis.port")
    db, _ := beego.AppConfig.Int("redis.db")
    password := beego.AppConfig.String("redis.password")

    client := redis.NewClient(&redis.Options{
        Addr:     host + ":" + port,
        Password: password, // no password set
        DB:       db,       // use default DB
    })
    client.FlushDB()
    return client
}
