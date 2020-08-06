package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var ctx = context.Background()

func testClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}

// 5000 set/s 我的电脑
// value如果为1MB大小，10000耗时约50秒，200 set/s，16G内存勉强支撑redis工作
func BatchSet(limit int) {
	defer duration(track(GetFunctionName(BatchSet)))
	rdb := NewClient()
	var err error
	for i := 0; i < limit; i++ {
		err = rdb.Set(ctx, strconv.Itoa(i), RepeatStringWithSize(1024*1024, "a", strconv.Itoa(i)),0).Err()
		if err != nil {
			panic(err)
		}
	}
}

// value如果为1MB大小，10000耗时约13秒，1000 set/s
func BatchSetVerify(limit int) {
	defer duration(track(GetFunctionName(BatchSet)))
	rdb := NewClient()
	for i := 0; i < limit; i++ {
		value, err := rdb.Get(ctx, strconv.Itoa(i)).Result()
		if err != nil {
			panic(err)
		}
		s := strings.Split(value, "-")
		if strconv.Itoa(i) != s[1] {
			panic(errors.New(fmt.Sprintf("values not match for key %d", i)))
		}
	}
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func RepeatStringWithSize(size int, repeat, suffix string) string {
	v := strings.Repeat(repeat, size)
	return strings.Join([]string{v, suffix}, "-")
}

// 如何生成一个长字符串
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
// https://stackoverflow.com/questions/33139020/can-golang-multiply-strings-like-python-can
// https://stackoverflow.com/questions/38418171/how-to-generate-unique-random-string-in-a-length-range-using-golang
// https://stackoverflow.com/questions/50498930/fastest-way-to-allocate-a-large-string-in-go
func main() {
	//BatchSet(10000)
	BatchSetVerify(10000)
}