package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

/*
 提供redis的一些方法
*/

// 接收的数据格式
type moneyID struct {
	MoneyId string `form:"money_id" binding:"required"`
}

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func InitRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.163.129:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// 注意redis的V8版本，需要传入context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

// 查询需求，查询某一期的中奖名单，redis的键值对存的应该是 string: "{user: money}"
// 拿到redis的数据后要先用json转一下
func getLuckyUserList(Key string) (userList map[string]int) {
	ctx := context.Background()
	// 先查询redis中是否存在Key,有的话直接返回，没有的话先访问MySQL，返回的同时写到redis
	fmt.Println("key ", Key)
	val, err := rdb.Get(ctx, Key).Result()
	if err == redis.Nil {
		// redis找不到，先查询MySQL,再刷新redis
		tmp := FindUserData(Key)
		fmt.Println("tmp", tmp)
		if len(tmp) == 0 {
			log.Println("数据不在redis和MySQL")
			return nil
		}
		if err := rdb.Set(ctx, Key, mapToString(tmp), 0).Err(); err != nil {
			panic(err)
		}
		userList = tmp
		log.Printf("mysql查询到key: %s 对应的值是: %s\n", Key, val)
	} else if err != nil {
		panic(err)
	} else {
		userList = stringToMap(val)
		log.Printf("redis查询到key: %s 对应的值是: %s\n", Key, val)
	}
	return
}

// 自定义string转map
func stringToMap(Data string) (res map[string]int) {
	err := json.Unmarshal([]byte(Data), &res)
	if err != nil {
		panic(err)
	}
	return
}

// 自定义map转string
func mapToString(Data map[string]int) (res string) {
	if dataType , err := json.Marshal(Data); err != nil {
		panic(err)
	} else {
		res = string(dataType)
	}
	return
}

// 定期刷新redis中的key值，和MySQL对齐
func SyncRedisAndMysql(duration time.Duration) {
	var cursor uint64
	ctx := context.Background()
	for {
		if keys, _, err := rdb.Scan(ctx, cursor, "*", 10).Result(); err != nil {
			panic(err)
		} else {
			for _, key := range keys {
				mysqlVal := FindUserData(key)
				if redisVal, err := rdb.Get(ctx, key).Result(); err != nil {
					panic(err)
				} else {
					if len(mysqlVal) != 0 && mapToString(mysqlVal) != redisVal {
						if err := rdb.Set(ctx, key, mapToString(mysqlVal), 0).Err(); err != nil {
							panic(err)
						}
					}
				}
			}
		}
		log.Println("redis sync with mysql")
		time.Sleep(duration)
	}
}

func GetRedisData(ctx *gin.Context)  {
	var (
		MoneyId  moneyID
		err error
	)
	if err = ctx.ShouldBindQuery(&MoneyId); err != nil {
		ctx.JSON(200, gin.H{
			"msg": "参数错误",
			"data": "nil",
		})
		return
	}
	if val := getLuckyUserList(MoneyId.MoneyId); val != nil {
		ctx.JSON(200, val)
	} else {
		ctx.JSON(200, gin.H{
			"msg": "查询不到对应的数据",
		})
	}

}