package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func main() {
	m := make(map[string]int)
	fmt.Println(reflect.TypeOf(m))
}

// 单独的监控协程
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "收到信号，监控退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println(name, "goroutine监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}