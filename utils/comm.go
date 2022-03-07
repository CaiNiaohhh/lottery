package utils

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"strings"
	"time"
)

//数据库配置
const (
	userName = "root"
	password = ""
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "lottery"
)

var (
	packageList []int
	userCh chan string
	userList map[string]bool
	moneyId string
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

	// 初始化

	if msg, err := DB.Exec(InitTableByName("money")); err != nil{ // ignore_security_alert;
		fmt.Printf("error is %s\n", err)
	} else {
		fmt.Printf("create table money success, Msg is %v\n", msg)
	}

	if msg, err := DB.Exec(InitTableByName("user")); err != nil{ // ignore_security_alert;
		fmt.Printf("error is %s\n", err)
	} else {
		fmt.Printf("create table user success, Msg is %v\n", msg)
	}
}

type moneyData struct {
	Total int `json:"total" form:"total" binding:"required"`
	Num int `json:"num" form:"num" binding:"required"`
}

type userData struct {
	Uid string `form:"uid" binding:"required"`
}

func Set(ctx *gin.Context) {
	var (
		money moneyData
		err error
	)
	if err = ctx.ShouldBindQuery(&money); err != nil {
		ctx.JSON(200, gin.H{
			"msg": "参数错误",
			"data": "nil",
		})
		return
	}
	// 定义一个红包分发通道
	userCh = make(chan string, 1)
	// 重置抢到红包的用户白名单
	userList = make(map[string]bool)
	log.Println("money: ", money)
	//now := strconv.FormatInt(time.Now().Unix(),10)
	now := time.Now().Format("2006-01-02 15:04:05")
	//log.Println("Now", now)

	tmp := md5.Sum([]byte(now))
	moneyId = hex.EncodeToString(tmp[:])[:20]

	// 创建一个map来存储数据
	insertData := map[string]interface{}{
		"MoneyId": moneyId,
		"Num": money.Num,
		"time": now,
		"Money": money.Total,
	}
	// 插入数据库
	if InsertData("money", insertData) {
		ctx.JSON(200, gin.H{
			"msg": "insert data success",
			"data": insertData,
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg": "insert data fail",
			"data": insertData,
		})
	}

	// 计算随机算法，将Total 分为 Num个 红包
	packageList = generalPackageList(money.Total, money.Num)
	log.Println("packageList is:", packageList)

}

func generalPackageList(total int, num int) []int {
	// 每个红包的范围是[(total/num)/10, (total/num)*2]
	var (
		count = num
		leave = total
		res []int
		max = (total / num) * 2
		min = (total / num) / 10
	)
	for count > 0 {
		if count == 1 {
			res = append(res, leave)
		} else {
			tmp := rand.Intn(max - min) + min
			if tmp >= leave {
				continue
			}
			res = append(res, tmp)
			leave = leave - tmp
		}
		count -= 1
	}
	return res
}

func Get(ctx *gin.Context) {
	// 使用channel来存储
	var (
		user userData
		err error
	)
	isLucky := false
	if err = ctx.ShouldBindQuery(&user); err != nil {
		ctx.JSON(200, gin.H{
			"msg": "参数错误",
			"data": "nil",
		})
		return
	}
	userCh <- user.Uid
	money := 0
	if len(packageList) == 0 {
		ctx.JSON(200, gin.H{
			"msg": "红包已经被抢完",
		})
	} else if _, ok := userList[user.Uid]; ok {
		ctx.JSON(200, gin.H{
			"msg": "您已经抢过红包",
		})
	} else {
		money = packageList[0]
		packageList = packageList[1:]
		ctx.JSON(200, gin.H{
			"msg": "成功抢到红包",
			"money": money,
		})
		isLucky = true
	}
	<- userCh
	userList[user.Uid] = true
	//now := strconv.FormatInt(time.Now().Unix(),10)
	now := time.Now().Format("2006-01-02 15:04:05")
	// 创建一个map来存储数据
	insertData := map[string]interface{}{
		"UserId": user.Uid,
		"isLucky": isLucky,
		"time": now,
		"Money": money,
		"MoneyId": moneyId,
	}
	// 将数据写入数据库
	if InsertData("user", insertData) {
		log.Println(200, gin.H{
			"msg": "insert user data success",
			"data": insertData,
		})
	} else {
		log.Println(200, gin.H{
			"msg": "insert user data fail",
			"data": insertData,
		})
	}
}
