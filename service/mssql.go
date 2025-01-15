package service

import (
	"database/sql"
	"go.uber.org/zap"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"vline/model"
	"vline/util"
)

func initDb(log *zap.Logger) (*gorm.DB, *sql.DB) {
	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:sa123456@192.168.100.80:1433?database=test&encrypt=disable"), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "dbo",
		},
		Logger: &util.ZapGormLogger{Logger: log},
	})

	if err != nil {
		panic(err)
	}

	client, _ := db.DB()

	client.SetMaxIdleConns(5)
	client.SetMaxOpenConns(10)
	client.SetConnMaxIdleTime(time.Second * 60)
	client.SetConnMaxLifetime(time.Minute * 25)

	return db, client
}

func MssqlHandle(logger *zap.Logger) {

	db, client := initDb(logger)

	defer client.Close()

	for {
		t := &[]model.Model{}
		db.Raw("select * from test").Scan(t)

		logger.Info("fetch data: ", zap.Reflect("data", t))

		time.Sleep(time.Second * 2)
	}

}
