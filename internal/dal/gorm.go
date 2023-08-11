package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"minitok/internal/conf"
)

var Gorm *gorm.DB

// InitGorm Init TODO: 增加更多配置
func InitGorm() {
	db, err := gorm.Open(mysql.Open(conf.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})

	if err != nil {
		panic(fmt.Errorf("open db:%w", err))
	}

	Gorm = db
}
