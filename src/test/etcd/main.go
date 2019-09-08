package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd failed, err: ", err)
		return
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = cli.Put(ctx, "nginx_log", "Hello Etcd")
	cancel()
	if err != nil {
		fmt.Println("put message to etcd failed, err: ", err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Get(ctx, "nginx_log")
	cancel()
	if err != nil {
		fmt.Println("get message for etcd failed, err: ", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("Key: %s, Val: %s\n", ev.Key, ev.Value)
	}
}
