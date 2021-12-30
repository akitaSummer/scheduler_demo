package main

import (
	"fmt"
	"time"

	"context"
	"github.com/coreos/etcd/clientv3"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		putResp *clientv3.PutResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 发起链接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)

	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "hello"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Revision", putResp.Header.Revision)
	}

	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "etcd"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Revision", putResp.Header.Revision)
	}
}
