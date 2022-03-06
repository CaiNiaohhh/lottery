package utils

import (
	"fmt"
	"log"
	"strings"
)

func InitDBByName(DBName string) string {
	sql := strings.Join([]string{"create database if not exists ", DBName, ";"}, "")
	return sql
}

/* DBName 只能是money or user
*/
func InitTableByName(tableName string) string {
	if tableName != "money" && tableName != "user" {
		return "error"
	}
	sql := ""
	if tableName == "money" {
		sql = `
		create table if not exists money(
		id int auto_increment primary key not null,
		MoneyId varchar(20),
		Num int,
		time varchar(20),
		Money int);
		`
	} else {
		sql = `
		create table if not exists user(
		id int auto_increment primary key not null,
		UserId varchar(20),
		time varchar(20),
		isLucky tinyint(1),
		Money int,
		MoneyId varchar(20)
		);
		`
	}
	return sql
}

func InsertData(tableName string, mapData map[string]interface{}) bool {
	if tableName != "money" && tableName != "user" {
		return false
	}
	sql := ""
	if tableName == "money" {
		sql = fmt.Sprintf(
			"insert into money(MoneyId, Num, time, Money) values ('%s', %d, '%s', %d);",
			mapData["MoneyId"].(string),
			mapData["Num"].(int),
			mapData["time"].(string),
			mapData["Money"].(int),
		)
	} else {
		sql = fmt.Sprintf(
			"insert into user(UserId, isLucky, time, Money, MoneyId) values ('%s', %v, '%s', %d, '%s');",
			mapData["UserId"].(string),
			mapData["isLucky"].(bool),
			mapData["time"].(string),
			mapData["Money"].(int),
			mapData["MoneyId"].(string),
		)
	}
	if msg, err := DB.Exec(sql); err != nil {
		log.Println("insert data fail, the sql is ", sql)
		return false
	} else {
		log.Println("insert data success, the msg is ", msg)
		return true
	}
}

// 查询user表中MoneyId对应的所有isLucky为true的userId:Money键值对
func FindUserData(MoneyId string) (res map[string]int) {
	res = make(map[string]int)
	var (
		UserId string
		Money int
	)
	sql := fmt.Sprintf(`select UserId, Money from user where MoneyId = '%s' and isLucky = 1`, MoneyId)
	rows, err := DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println("查询数据失败")
		return nil
	}
	for rows.Next() {
		rows.Scan(&UserId, &Money)
		res[UserId] = Money
	}
	return

}