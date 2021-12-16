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
		delResp *clientv3.DeleteResponse
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

	// 读取/cron/jobs下的所有
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job1"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Revision", delResp.Header.Revision)
	}
}
