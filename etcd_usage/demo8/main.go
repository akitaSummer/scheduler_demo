package main

import (
	"fmt"
	"time"

	"context"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putOp  clientv3.Op
		getOp  clientv3.Op
		OpResp clientv3.OpResponse
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

	putOp = clientv3.OpPut("/cron/jobs/job8", "")

	if OpResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(OpResp)

	getOp = clientv3.OpGet("/cron/jobs/job8")

	if OpResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(OpResp)
}
