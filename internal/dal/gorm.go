package dal

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"minitok/internal/constant"
)

var datasourceConstants = &constant.AllConstants.Datasource

func InitGorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open(datasourceConstants.DSNString()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})

	if err != nil {
		klog.Fatalf("Gorm instance can not be established: %v", err)
	}

	return db
}
