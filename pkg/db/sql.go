package db

import (
	"errors"
	"fmt"
	"gin-bic/config"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var mysqlDatabase *gorm.DB
var mysqlMock sqlmock.Sqlmock
var SqlMockErr = errors.New("SqlMock 测试使用错误")
var sqlMockOnce sync.Once

func SqlDB() *gorm.DB {
	return mysqlDatabase
}
func SqlDBMock() sqlmock.Sqlmock {
	SetupMock()
	return mysqlMock
}

func MustInitMysql(cfg *config.MySqlConfig) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s`,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	gormCfg := gorm.Config{
		CreateBatchSize: 200,
	}
	db, err := gorm.Open(mysql.Open(dsn), &gormCfg)
	if err != nil {
		panic("mysql connect error " + err.Error())
	}
	DB, err := db.DB()
	if err != nil {
		panic("mysql connect load DB " + err.Error())
	}
	DB.SetMaxIdleConns(20)
	DB.SetMaxOpenConns(200)
	DB.SetConnMaxLifetime(time.Minute * 30)
	mysqlDatabase = db
}

func SetupMock() {
	sqlMockOnce.Do(func() {
		sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp)) // mock sql.DB
		if err != nil {
			panic(err)
		}
		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB, SkipInitializeWithVersion: true}), &gorm.Config{}) // open gorm db
		if err != nil {
			panic(err)
		}
		mysqlDatabase = gdb
		mysqlMock = mock
	})
}
