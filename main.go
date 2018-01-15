package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"

	"database/sql"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
var db *sql.DB
var pool *redis.Pool
func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	db = newSqlPool(conf.Server.Mysql) //数据库连接
	pool = newPool(conf.Server.Redis) //redis连接
	defer db.Close()
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)


}

//连接数据库
func newSqlPool(dconf conf.MysqlConf) *sql.DB {
	sqldb := fmt.Sprint(dconf.User, ":", dconf.Password, "@tcp(", dconf.Host, ":", dconf.Port, ")/", dconf.Database)
	db, _ := sql.Open("mysql", sqldb)

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.Ping()
	return db
}

//生成连接池方法,连接redis
func newPool(rconf conf.RedisConf) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   1000,
		MaxActive: 5000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprint(rconf.Host, ":", rconf.Port))
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", rconf.Auth); err != nil {
				return nil, err
			}

			c.Do("SELECT", rconf.Database)
			return c, err

		},
	}
}
