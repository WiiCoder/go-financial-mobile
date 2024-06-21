package mysql

import (
	"fmt"
	"github.com/wiiCoder/installment/pkg/common/config"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
)

func NewMysql() (*gorm.DB, error) {
	conf := config.Config.Mysql
	// 格式化数据库链接
	mysqlDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset, conf.ParseTime,
	)
	return gorm.Open(gormMysql.Open(mysqlDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.TablePrefix, // 表前缀
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用AutoMigrate创建外键
	})
}
