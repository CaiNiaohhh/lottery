package utils

/*
	主要存放一些通用方法
*/

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test"
)

//Db数据库连接池
var DB *sql.DB

// 初始化与MySQL的连接
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

// 根据传入的字段名和字段值返回对应的查询结果
func Query(tableName string, key string, value interface{}) interface{} {
	//rows, err := DB.Query("select id,name,address from users")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for rows.Next() {
	//	rows.Scan(&id, &name, &address)
	//	fmt.Println(id, name, address)
	//}
	//defer rows.Close()
	query := strings.Join([]string{"select * from tableName where ", key, " = ?"}, "")
	Row := ""
	err := DB.QueryRow(query, value.(string)).Scan(&Row) // ignore_security_alert
	if err != nil {
		fmt.Printf("err Msg is %v", err)
		return nil
	}
	

}
