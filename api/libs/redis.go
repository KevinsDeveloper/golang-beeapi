package libs

import (
    "github.com/astaxie/beego"
    "github.com/garyburd/redigo/redis"
    "strconv"
    "time"
)

// Cache is Redis cache adapter.
type Redis struct {
    Pool     *redis.Pool // redis connection pool
    conninfo string
    dbNum    int
    key      string
    password string
}

func (this *Redis) init() error {
    host := beego.AppConfig.String("redis.host")
    port := beego.AppConfig.String("redis.port")
    db := beego.AppConfig.String("redis.db")
    password := beego.AppConfig.String("redis.password")

    this.conninfo = host + ":" + port
    this.dbNum, _ = strconv.Atoi(db)
    if password != "" {
        this.password = password
    }

    this.connect()

    c := this.Pool.Get()
    defer c.Close()

    return c.Err()
}

// actually do the redis cmds
func (rc *Redis) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
    c := rc.Pool.Get()
    defer c.Close()

    return c.Do(commandName, args...)
}

// connect to redis.
func (rc *Redis) connect() {
    dialFunc := func() (c redis.Conn, err error) {
        c, err = redis.Dial("tcp", rc.conninfo)
        if err != nil {
            return nil, err
        }

        if rc.password != "" {
            if _, err := c.Do("AUTH", rc.password); err != nil {
                c.Close()
                return nil, err
            }
        }

        _, selecterr := c.Do("SELECT", rc.dbNum)
        if selecterr != nil {
            c.Close()
            return nil, selecterr
        }
        return
    }
    // initialize a new pool
    rc.Pool = &redis.Pool{
        MaxIdle:     3,
        IdleTimeout: 180 * time.Second,
        Dial:        dialFunc,
    }
}
