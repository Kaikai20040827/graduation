package config 

import (
    "time"
	"github.com/spf13/viper"
)

// 数据库类型（名称），对应ServerConfig结构体的Driver
type DatabaseType string

const (
    MySQL DatabaseType = "mysql"
    //目前只支持mysql
    //等待后续开发 => 仓库：https://github.com/Kaikai20040827/graduation
)

type ServerConfig struct {
    Appname string `mapstructure:"APP_NAME" json:"app_name" yaml:"app_name"` 
    Env     string `mapstructure:"ENV" json:"env" yaml:"env"`

    Debug   bool   `mapstructure:"DEBUG" json:"debug" yaml:"debug"`

    Host    string `mapstructure:"HOST" json:"host" yaml:"host"`
    Port    string `mapstructure:"PORT" json:"port" yaml:"port"`
}

type DatabaseConfig struct {
    Driver DatabaseType `mapstructure:"DRIVER" json:"driver" yaml:"driver"`

    Host     string `mapstructure:"DB_HOST" json:"db_host" yaml:"db_host"`
    Port     string `mapstructure:"DB_PORT" json:"db_port" yaml:"db_port"`
    User     string `mapstructure:"DB_USER" json:"db_user" yaml:"db_user"`
    Password string `mapstructure:"DB_PASSWORD" json:"db_password" yaml:"db_password"`
    DBName   string `mapstructure:"DB_NAME" json:"db_name" yaml:"db_name"`

    Debug    bool   `mapstructure:"DB_DEBUG" json:"debug" yaml:"debug"`
}

type JWTConfig struct {
    Secret      string        `mapstructure:"JWT_SECRET" json:"secret" yaml:"secret"`
    Issuer      string        `mapstructure:"JWT_ISSUER" json:"issuer" yaml:"issuer"`
    Audience    string        `mapstructure:"JWT_AUDIENCE" json:"audience" yaml:"audience"`

}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}


func LoadConfig() *Config {
    //设置配置文件名、文件类型、文件路径
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

    // 读取配置文件
    err := viper.ReadInConfig()
    if err != nil {
        panic("无法读取配置文件: " + err.Error())
    }

    //解析配置文件
    var cfg *Config
    err = viper.Unmarshal(&cfg)
    if err != nil {
        panic("无法解析配置文件: " + err.Error())
    }
    return cfg
}

