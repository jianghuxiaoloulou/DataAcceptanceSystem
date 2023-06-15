package model

import (
	"WowjoyProject/DataAcceptanceSystem/pkg/setting"
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// Pacs
func NewPacsDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.PacsDBType, databaseSetting.PacsDBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.DBMaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.DBMaxOpenConns)
	db.SetMaxIdleConns(databaseSetting.DBMaxIdleConns)

	return db, nil
}

//门诊
func NewMZApplyDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.MZApplyDBType, databaseSetting.MZApplyDBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.DBMaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.DBMaxOpenConns)
	db.SetMaxIdleConns(databaseSetting.DBMaxIdleConns)

	return db, nil
}

// 住院
func NewZYApplyDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.ZYApplyDBType, databaseSetting.ZYApplyDBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.DBMaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.DBMaxOpenConns)
	db.SetMaxIdleConns(databaseSetting.DBMaxIdleConns)

	return db, nil
}

// 体检
func NewTJApplyDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.TJApplyDBType, databaseSetting.TJApplyDBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.DBMaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.DBMaxOpenConns)
	db.SetMaxIdleConns(databaseSetting.DBMaxIdleConns)

	return db, nil
}
