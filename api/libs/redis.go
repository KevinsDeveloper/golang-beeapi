package libs

import (
    "github.com/garyburd/redigo/redis"
    "github.com/astaxie/beego"
    "time"
)

var client *redis.Pool

func init() {
    host := beego.AppConfig.String("redis.host")
    port := beego.AppConfig.String("redis.port")
    dbNum, _ := beego.AppConfig.Int("redis.db")
    password := beego.AppConfig.String("redis.password")

    dialFunc := func() (c redis.Conn, err error) {
        c, err = redis.Dial("tcp", host+":"+port)
        if err != nil {
            return nil, err
        }

        if password != "" {
            if _, err := c.Do("AUTH", password); err != nil {
                c.Close()
                return nil, err
            }
        }

        _, selecterr := c.Do("SELECT", dbNum)
        if selecterr != nil {
            c.Close()
            return nil, selecterr
        }
        return
    }
    // initialize a new pool
    client = &redis.Pool{
        MaxIdle:     3,
        IdleTimeout: 180 * time.Second,
        Dial:        dialFunc,
    }

}

// actually do the redis cmds
func Redis(commandName string, args ...interface{}) (reply interface{}, err error) {
    c := client.Get()
    defer c.Close()
    return c.Do(commandName, args...)
}
