package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type base struct {
	Port      string `json:"port,omitempty"`
	Model     string `json:"model,omitempty"`
	ServeName string `json:"serve_name,omitempty"`
	Banner    string `json:"banner,omitempty"`
}

type log struct {
	Level    string `json:"level,omitempty"`
	FilePath string `json:"filePath,omitempty"`
}

type mysql struct {
	Username                  string `json:"username,omitempty"`
	Password                  string `json:"password,omitempty"`
	Host                      string `json:"host,omitempty"`
	Port                      string `json:"port,omitempty"`
	DB                        string `json:"db,omitempty"`
	Charset                   string `json:"charset,omitempty"`
	ParseTime                 string `json:"parseTime,omitempty"`
	Loc                       string `json:"loc,omitempty"`
	DefaultStringSize         int    `json:"defaultStringSize,omitempty"`
	DisableDatetimePrecision  bool   `json:"disableDatetimePrecision,omitempty"`
	DontSupportRenameIndex    bool   `json:"dontSupportRenameIndex,omitempty"`
	DontSupportRenameColumn   bool   `json:"dontSupportRenameColumn,omitempty"`
	SkipInitializeWithVersion bool   `json:"skipInitializeWithVersion,omitempty"`
	MaxIdleConns              int    `json:"maxIdleConns,omitempty"`
	MaxOpenConns              int    `json:"maxOpenConns,omitempty"`
	ConnMaxLifetime           int    `json:"connMaxLifetime,omitempty"`
}

type postgre struct {
	Username        string `json:"username,omitempty"`
	Password        string `json:"password,omitempty"`
	Host            string `json:"host,omitempty"`
	Port            string `json:"port,omitempty"`
	Charset         string `json:"charset,omitempty"`
	ParseTime       string `json:"parseTime,omitempty"`
	Loc             string `json:"loc,omitempty"`
	MaxIdleConns    int    `json:"maxIdleConns,omitempty"`
	MaxOpenConns    int    `json:"maxOpenConns,omitempty"`
	ConnMaxLifetime int    `json:"connMaxLifetime,omitempty"`
}

type redis struct {
	Host     string `json:"host,omitempty"`
	Password string `json:"password,omitempty"`
	DB       uint   `json:"db,omitempty"`
	Prefix   string `json:"prefix,omitempty"`
}

type email struct {
	User           string `json:"user,omitempty"`
	Password       string `json:"password,omitempty"`
	Address        string `json:"address,omitempty"`
	Host           string `json:"host,omitempty"`
	Authentication string `json:"authentication,omitempty"`
}

type Config struct {
	Base    base    `json:"base,omitempty"`
	Log     log     `json:"log,omitempty"`
	Mysql   mysql   `json:"mysql,omitempty"`
	Postgre postgre `json:"postgre,omitempty"`
	Redis   redis   `json:"redis,omitempty"`
	Email   email   `json:"email,omitempty"`
}

var config *Config

// GetConfig 获取配置文件
func GetConfig() *Config {
	if config == nil {
		initConfig()
	}
	return config
}

// initConfig 初始化配置文件
func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../steward/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unmarshal config error: %w", err))
	}
}

// InitConfig 初始化配置文件
func InitConfig() {
	initConfig()
}
