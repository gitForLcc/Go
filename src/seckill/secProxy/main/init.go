package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	etcd_client "go.etcd.io/etcd/clientv3"
)

var (
	redisPool  *redis.Pool
	etcdClient *etcd_client.Client
)

func initRedis() error {
	redisPool = &redis.Pool{
		MaxIdle:     secKillConf.redisConf.redisMaxIdle,   //空闲连接数
		MaxActive:   secKillConf.redisConf.redisMaxActive, //活跃链接数|最大链接数 0表示没有限制
		IdleTimeout: time.Duration(secKillConf.redisConf.redisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.redisConf.redisAddr)
		},
	}
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err: %v", err)
		return err
	}

	return nil
}

func initEtcd() error {
	etcdClient, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{secKillConf.etcdConf.etcdAddr},
		DialTimeout: time.Duration(secKillConf.etcdConf.timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err: ", err)
		return err
	}
	defer etcdClient.Close()
	return nil
}

func converLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	default:
		return logs.LevelDebug
	}
}

func initLogs() error {
	config := make(map[string]interface{})
	config["filename"] = secKillConf.logPath
	config["level"] = converLogLevel(secKillConf.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Printf("marshal failed, err: %v\n", err)
		return err
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))

	return nil
}

// 初始化
func initSec() error {
	err := initLogs()
	if err != nil {
		logs.Error("init logger failed, err: %v", err)
		return nil
	}
	err = initRedis()
	if err != nil {
		logs.Error("init redis failed, err: %v", err)
		return err
	}

	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed, err: %v", err)
		return err
	}

	logs.Info("init sec succ")

	return nil
}
