package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/google/wire"
)

type ServerConf struct {
	RunMode    string `flag:"run_mode" toml:"run_mode" json:"run_mode"`
	ServerHost string `flag:"server_host" toml:"server_host" json:"server_host"`
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
	RedisDb   string `flag:"redis_db" toml:"redis_db" json:"redis_db"`
	RedisAddr string `flag:"redis_addr" toml:"redis_addr" json:"redis_addr"`
	RedisPw   string `flag:"redis_pw" toml:"redis_pw" json:"redis_pw"`
}


type QiNiuConf struct {
	AccessKey string `flag:"access_key" toml:"access_key" json:"access_key"`
	SerectKey string `flag:"serect_key" toml:"serect_key" json:"serect_key"`
	Bucket    string `flag:"bucket" toml:"bucket" json:"bucket"`
	Domain    string `flag:"domain" toml:"domain" json:"domain"`
}

type MailConf struct {
	MailUsername string `flag:"mail_username" toml:"mail_username" json:"username"`
	MailPasswd   string `flag:"mail_passwd" toml:"mail_passwd" json:"mail_passwd"`
	MailHost     string `flag:"mail_host" toml:"mail_host" json:"mail_host"`
	MailAddress  string `flag:"mail_address" toml:"mail_address" json:"mail_address"`
}

type Config struct {
	Server ServerConf `flag:"server" toml:"server" json:"server"`
	Mysql  MysqlConf  `flag:"mysql" toml:"mysql" json:"mysql"`
	Redis  RedisConf  `flag:"redis" toml:"redis" json:"redis"`
	QiNiu  QiNiuConf  `flag:"qiniu" toml:"qiniu" json:"qiniu"`
	Mail   MailConf   `flag:"mail" toml:"mail" json:"mail"`
}

func NewConfig() *Config {
	var Conf Config
	CONFIG_FILE := "../config/gin_api_frame.conf"
	_, err := toml.DecodeFile(CONFIG_FILE, &Conf)
	if err != nil {
		fmt.Println("config file load failed ,please check the file path:", err)
		panic(err)
	}
	fmt.Printf("config file load success,the content is :%v ", Conf)
	return &Conf
}

var ConfigProviderSet = wire.NewSet(NewConfig)
