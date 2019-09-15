package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	secKillConf *SecSkillConf
)

type RedisConf struct {
	redisAddr        string
	redisMaxIdle     int
	redisMaxActive   int
	redisIdleTimeout int
}

type EtcdConf struct {
	etcdAddr string
	timeout  int
}

type SecSkillConf struct {
	redisConf RedisConf
	etcdConf  EtcdConf
	logPath   string
	logLevel  string
}

func initConfig() error {
	secKillConf = &SecSkillConf{}

	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("read config secc, redis addr: %v", redisAddr)
	logs.Debug("read config secc, etcd addr: %v", etcdAddr)

	secKillConf.etcdConf.etcdAddr = etcdAddr
	secKillConf.redisConf.redisAddr = redisAddr

	if len(etcdAddr) == 0 || len(redisAddr) == 0 {
		err := fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return err
	}

	redisMaxIdle, err := beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		return err
	}

	redisMaxActive, err := beego.AppConfig.Int("redis_max_active")
	if err != nil {
		return err
	}

	redisIdleTimeout, err := beego.AppConfig.Int("redis_idle_timeout")
	if err != nil {
		return err
	}

	secKillConf.redisConf.redisMaxIdle = redisMaxIdle
	secKillConf.redisConf.redisMaxActive = redisMaxActive
	secKillConf.redisConf.redisIdleTimeout = redisIdleTimeout

	etcdTimeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		return err
	}

	secKillConf.etcdConf.timeout = etcdTimeout

	secKillConf.logPath = beego.AppConfig.String("logs_path")
	secKillConf.logLevel = beego.AppConfig.String("logs_level")

	return nil
}
