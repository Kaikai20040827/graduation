package config

import (
    "fmt"
    "strings"

    "github.com/spf13/viper"
)

type ServerConfig struct {
    AppName  string `mapstructure:"app_name"`
    Env      string `mapstructure:"env"`
    Debug    bool   `mapstructure:"debug"`
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    TimeZone string `mapstructure:"time_zone"`
}

type DatabaseConfig struct {
    Driver   string `mapstructure:"driver"`
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    Name     string `mapstructure:"name"`
    Debug    bool   `mapstructure:"debug"`
}

type JWTConfig struct {
    Secret   string `mapstructure:"secret"` //签名密钥
    Issuer   string `mapstructure:"issuer"` // 签发者
    Audience string `mapstructure:"audience"`
}

type Config struct {
    Server   ServerConfig   `mapstructure:"server"` 
    Database DatabaseConfig `mapstructure:"database"`
    JWT      JWTConfig      `mapstructure:"jwt"`
}

func LoadConfig() (*Config, error) {
    v := viper.New()

    // 配置文件名
    v.SetConfigName("config")
    v.SetConfigType("yaml")
    v.AddConfigPath("../../") // 根目录读取 config.yaml
    fmt.Println("✓ Loading config file done")

    // 支持环境变量覆盖，例如：SERVER_APP_NAME
    v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    v.AutomaticEnv()

    setDefaults(v)
    fmt.Println("✓ Setting default configuration done")

    // 读取 YAML
    if err := v.ReadInConfig(); err != nil {
        fmt.Println("⚠ config.yaml 未找到，默认使用默认值 + 环境变量")
    } else {
        fmt.Printf("✓ Using config file: %s, not using default configuration\n", v.ConfigFileUsed())
    }
    
    var cfg Config
    if err := v.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("解析配置失败: %w", err)
    }
    fmt.Println("✓ Unmarshalling config file done")

    // 校验
    if err := validateConfig(&cfg); err != nil {
        return nil, err
    }
    fmt.Println("✓ Validating config file done")

    return &cfg, nil
}

func setDefaults(v *viper.Viper) {
    v.SetDefault("server.app_name", "secure_file_box")
    v.SetDefault("server.env", "development")
    v.SetDefault("server.debug", true)
    v.SetDefault("server.host", "127.0.0.1")
    v.SetDefault("server.port", 8080)
    v.SetDefault("server.time_zone", "Asia/Shanghai")

    v.SetDefault("database.driver", "mysql")
    v.SetDefault("database.host", "localhost")
    v.SetDefault("database.port", 3306)
    v.SetDefault("database.user", "root")
    v.SetDefault("database.password", "0827")
    v.SetDefault("database.name", "secure_file_box")
    v.SetDefault("database.debug", false)

    v.SetDefault("jwt.secret", "PLEASE_CHANGE_ME_32_CHARS_MINIMUM")
    v.SetDefault("jwt.issuer", "secure_file_box")
    v.SetDefault("jwt.audience", "secure_users")
}

func validateConfig(cfg *Config) error {
    if len(cfg.JWT.Secret) < 32 {
        return fmt.Errorf("jwt.secret 必须 ≥ 32 字符")
    }
    if cfg.Database.Name == "" {
        return fmt.Errorf("database.name 不能为空")
    }
    return nil
}
