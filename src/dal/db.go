/*
封装数据库对象的包
*/
package dal

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	modelDB *sql.DB
)

// 初始化数据库连接相关的配置
// modelDBConnection：模型数据库连接字符串
// modelDBMaxOpenConns：最大开启连接数
// modelDBMaxIdleConns：最大空闲连接数
func InitDB(modelDBConnection string, modelDBMaxOpenConns, modelDBMaxIdleConns int) {
	modelDB = openMysqlConnection(modelDBConnection, modelDBMaxOpenConns, modelDBMaxIdleConns)
}

// 获取模型数据库对象
// 返回值：
// 模型数据库对象
func ModelDB() *sql.DB {
	return modelDB
}

func openMysqlConnection(connectionString string, maxOpenConns, maxIdleConns int) *sql.DB {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(errors.New(fmt.Sprintf("打开数据库失败,连接字符串为：%s", connectionString)))
	}

	if maxOpenConns > 0 && maxIdleConns > 0 {
		db.SetMaxOpenConns(maxOpenConns)
		db.SetMaxIdleConns(maxIdleConns)
	}

	if err := db.Ping(); err != nil {
		panic(errors.New(fmt.Sprintf("Ping数据库失败,连接字符串为：%s", connectionString)))
	}

	return db
}
