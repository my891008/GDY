package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

type MysqlConf struct{
	Host 	 string
	User 	 string
	Password string
	Database string
	Port 	 string
	Prefix   string
	Debug 	 bool
}
type RedisConf struct{
	Host 	 string
	Port 	 string
	Auth	 string
	Database string
}

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
	Mysql 		MysqlConf
	Redis       RedisConf
}



func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
