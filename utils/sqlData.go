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
		Money int);
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
			"insert into user(UserId, isLucky, time, Money) values ('%s', %v, '%s', %d);",
			mapData["UserId"].(string),
			mapData["isLucky"].(bool),
			mapData["time"].(string),
			mapData["Money"].(int),
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