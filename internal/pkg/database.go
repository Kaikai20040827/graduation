package pkg

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/Kaikai20040827/graduation/internal/config"
)

func NewDatabase(cfg *config.DatabaseConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.User,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.Name,
    )

    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
