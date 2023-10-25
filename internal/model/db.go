package model

import (
	"WowjoyProject/DataAcceptanceSystem/global"
	"WowjoyProject/DataAcceptanceSystem/pkg/setting"
	"database/sql"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// PACS集成平台数据库操作
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.DBType, databaseSetting.DBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.DBMaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.DBMaxOpenConns)
	db.SetMaxIdleConns(databaseSetting.DBMaxIdleConns)

	return db, nil
}

// 创建临时数据库连接
func NewTempDBEngine(dbType, dbConn string) (*sql.DB, error) {
	db, err := sql.Open(dbType, dbConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(global.DatabaseSetting.DBMaxLifetime) * time.Minute)
	db.SetMaxOpenConns(global.DatabaseSetting.DBMaxOpenConns)
	db.SetMaxIdleConns(global.DatabaseSetting.DBMaxIdleConns)
	return db, nil
}
