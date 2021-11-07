package dao

import (
	"github.com/pkg/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDaoServices(dsn, logLevel string) (SwapPairDaoInterface, SwapDaoInterface, error) {
	mysqlConn := mysql.New(mysql.Config{
		DSN: dsn,
	})
	db, err := gorm.Open(mysqlConn, &gorm.Config{
		Logger: logger.Default.LogMode(dbLogLevel(logLevel)),
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "[NewDaoServices]: failed to open db")
	}

	return NewSwapPairDao(db), NewSwapDao(db), nil
}

func dbLogLevel(level string) logger.LogLevel {
	switch level {
	case "SILENT":
		return logger.Silent
	case "ERROR":
		return logger.Error
	case "WARN":
		return logger.Warn
	case "INFO":
		return logger.Info
	}

	return logger.Warn
}
