package configs

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //mysql 驱动
	"github.com/jinzhu/gorm"
)

// MySQLConfig 数据库配置
type MySQLConfig struct {
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Database   string `json:"database"`
	Host       string `json:"host"`
	Parameters string `json:"parameters"`
	MaxIdle    int    `json:"maxidle"`
	MaxOpen    int    `json:"maxopen"`
	Debug      bool   `json:"debug"`
}

//       db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
// String 获取数据库配置参数
func (m *MySQLConfig) String() string {
	return fmt.Sprintf("%s:%s@/%s?%s", m.UserName, m.Password, m.Database, m.Parameters)
}

// New 打开数据库
func (m *MySQLConfig) New() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", m.String())
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(m.MaxIdle)
	db.DB().SetMaxOpenConns(m.MaxOpen)
	db.LogMode(m.Debug)
	return db, nil
}
