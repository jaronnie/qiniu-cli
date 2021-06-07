package initalize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var ctx = context.Background()

func InitMAC() *qbox.Mac {
	ak, sk := viper.GetString("qiniu.ak"), viper.GetString("qiniu.sk")
	mac := qbox.NewMac(ak, sk)
	return mac
}

//func GormMysql() *gorm.DB {
//	username := viper.GetString("mysql.username")
//	password := viper.GetString("mysql.password")
//	ip := viper.GetString("mysql.ip")
//	port := viper.GetString("mysql.port")
//	database := viper.GetString("mysql.database")
//	path := strings.Join([]string{username, ":", password,
//		"@(", ip, ":", port, ")/", database, "?charset=utf8&parseTime=true&loc=Local"}, "")
//
//	if db, err := gorm.Open(mysql.Open(path), &gorm.Config{}); err != nil {
//		os.Exit(0)
//		return nil
//	} else {
//		sqlDB, _ := db.DB()
//		sqlDB.SetMaxIdleConns(10)
//		sqlDB.SetMaxOpenConns(100)
//		return db
//	}
//}
//
//func MysqlTables(db *gorm.DB) {
//	err := db.AutoMigrate(&models.Login{})
//	if err != nil {
//		fmt.Println("register table error")
//	}
//}

func ConnectRedis() *redis.Client {
	ip := viper.GetString("redis.ip")
	port := viper.GetString("redis.port")
	client := redis.NewClient(&redis.Options{
		Addr:ip + ":" + port,
		DB:0,
		PoolSize: 100,
		MinIdleConns: 20,
		DialTimeout: 5 * time.Second,
	})
	pong, err := client.Ping(ctx).Result()

	if err != nil {
		log.Fatal(err)
	}

	if pong != "PONG" {
		log.Fatal("连接redis失败")
	} else {
		log.Println("已连接redis服务")
	}
	return client
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read error")
		os.Exit(0)
	}
}
