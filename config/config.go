package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type ServerConf struct {
	RunMode    string `flag:"run_mode" toml:"run_mode" json:"run_mode"`
	ServerPort string `flag:"server_port" toml:"server_port" json:"server_port"`
}

type MysqlConf struct {
	Db         string `flag:"db" toml:"db" json:"db"`
	DbHost     string `flag:"db_host" toml:"db_host" json:"db_host"`
	DbPort     string `flag:"db_port" toml:"db_port" json:"db_port"`
	DbUser     string `flag:"db_user" toml:"db_user" json:"db_user"`
	DbPassWord string `flag:"db_password" toml:"db_password" json:"db_password"`
	DbName     string `flag:"db_name" toml:"db_name" json:"db_name"`
}

type RedisConf struct {
	RedisDb     string `flag:"redis_db" toml:"redis_db" json:"redis_db"`
	RedisAddr   string `flag:"redis_addr" toml:"redis_addr" json:"redis_addr"`
	RedisPw     string `flag:"redis_pw" toml:"redis_pw" json:"redis_pw"`
	RedisDbName string `flag:"redis_dbname" toml:"redis_dbname" json:"redis_dbname"`
}

type Config struct {
	Server ServerConf `flag:"server" toml:"server" json:"server"`
	Mysql  MysqlConf  `flag:"mysql" toml:"mysql" json:"mysql"`
	Redis  RedisConf  `flag:"redis" toml:"redis" json:"redis"`
}

var Conf = Config{}

const (
	CONFIG_FILE = "../config/ji.conf"
)

func LoadConfig() {
	_, err := toml.DecodeFile(CONFIG_FILE, &Conf)
	if err != nil {
		fmt.Println("config file load failed ,please check the file path:", err)
		panic(err)
	}
	fmt.Printf("config file load success,the content is :%v ", Conf)
}
