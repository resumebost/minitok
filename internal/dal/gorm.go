package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"minitok/internal/conf"
)

// InitGorm Init TODO: 增加更多配置
func InitGorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})

	if err != nil {
		panic(fmt.Errorf("open db: %w", err))
	}

	return db
}
