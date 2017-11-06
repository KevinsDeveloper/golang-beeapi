package libs

import (
    "github.com/garyburd/redigo/redis"
    "time"
    "github.com/astaxie/beego"
    "strconv"
)

// Cache is Redis cache adapter.
type Redis struct {
    p        *redis.Pool // redis connection pool
    conninfo string
    dbNum    int
    key      string
    password string
}

func (this *Redis) Client() error {
    host := beego.AppConfig.String("redis.host")
    port := beego.AppConfig.String("redis.port")
    db := beego.AppConfig.String("redis.db")
    password := beego.AppConfig.String("redis.password")

    this.conninfo = host + ":" + port
    this.dbNum, _ = strconv.Atoi(db)
    if password != "" {
        this.password = password
    }

    dialFunc := func() (c redis.Conn, err error) {
        c, err = redis.Dial("tcp", this.conninfo)
        if err != nil {
            return nil, err
        }

        if this.password != "" {
            if _, err := c.Do("AUTH", this.password); err != nil {
                c.Close()
                return nil, err
            }
        }

        _, selecterr := c.Do("SELECT", this.dbNum)
        if selecterr != nil {
            c.Close()
            return nil, selecterr
        }
        return
    }
    // initialize a new pool
    this.p = &redis.Pool{
        MaxIdle:     3,
        IdleTimeout: 180 * time.Second,
        Dial:        dialFunc,
    }

    c := this.p.Get()
    defer c.Close()

    return c.Err()
}
