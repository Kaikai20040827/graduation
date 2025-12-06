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

//服务器配置
type ServerConfig struct {
    Appname string `mapstructure:"APP_NAME" json:"app_name" yaml:"app_name"` 
    Env     string `mapstructure:"ENV" json:"env" yaml:"env"`

    Debug   bool   `mapstructure:"DEBUG" json:"debug" yaml:"debug"`

    Host    string `mapstructure:"HOST" json:"host" yaml:"host"`
    Port    string `mapstructure:"PORT" json:"port" yaml:"port"`
}

//数据库配置
type DatabaseConfig struct {
    Driver DatabaseType `mapstructure:"DRIVER" json:"driver" yaml:"driver"`

    Host     string `mapstructure:"DB_HOST" json:"db_host" yaml:"db_host"`
    Port     string `mapstructure:"DB_PORT" json:"db_port" yaml:"db_port"`
    User     string `mapstructure:"DB_USER" json:"db_user" yaml:"db_user"`
    Password string `mapstructure:"DB_PASSWORD" json:"db_password" yaml:"db_password"`
    DBName   string `mapstructure:"DB_NAME" json:"db_name" yaml:"db_name"`

    Debug    bool   `mapstructure:"DB_DEBUG" json:"debug" yaml:"debug"`
}

//JWT 认证配置
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

// 服务器默认配置
func setDefaults() {
    viper.SetDefault("server.app_name", "secure_file_box")
    viper.SetDefault("server.env", "development")
    viper.SetDefault("server.debug", true)
    viper.SetDefault("server.host", "127.0.0.1")
    viper.SetDefault("server.port", "8080") //默认端口 8080，也可用443
    viper.SetDefault("server.time_zone", "Asia/Shanghai")
    // viper.SetDefault("server.read_timeout", 30*time.Second)
    // viper.SetDefault("server.write_timeout", 30*time.Second)
    // viper.SetDefault("server.idle_timeout", 120*time.Second)
    // viper.SetDefault("server.max_multipart_memory", 32<<20) // 32 MB
    
    // 数据库默认配置
    viper.SetDefault("database.driver", "mysql")
    viper.SetDefault("database.host", "root") //localhost 的 root 用户
    viper.SetDefault("database.port", "3306")
    viper.SetDefault("database.charset", "utf8mb4")
    viper.SetDefault("database.timezone", "Local")
    // viper.SetDefault("database.max_open_conns", 100)
    // viper.SetDefault("database.max_idle_conns", 10)
    // viper.SetDefault("database.conn_max_lifetime", time.Hour)
    // viper.SetDefault("database.conn_max_idle_time", 30*time.Minute)
    // viper.SetDefault("database.log_level", "warn")
    
    // JWT 默认配置
    viper.SetDefault("jwt.secret", "secret_key") //密钥
    viper.SetDefault("jwt.issuer", "gin-app") 
    viper.SetDefault("jwt.audience", "audience")
    // viper.SetDefault("jwt.access_expire", 24*time.Hour)
    // viper.SetDefault("jwt.refresh_expire", 7*24*time.Hour)
    // viper.SetDefault("jwt.enable_refresh", true)
    // viper.SetDefault("jwt.refresh_cookie_name", "refresh_token")
    // viper.SetDefault("jwt.token_lookup", "header:Authorization")
    // viper.SetDefault("jwt.token_head", "Bearer")
    // viper.SetDefault("jwt.cookie_http_only", true)
    // viper.SetDefault("jwt.cookie_secure", false)
    // viper.SetDefault("jwt.cookie_same_site", "Lax")
}