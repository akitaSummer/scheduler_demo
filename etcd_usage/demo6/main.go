package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	var (
		config           clientv3.Config
		client           *clientv3.Client
		err              error
		lease            clientv3.Lease
		leaseCreatetResp *clientv3.LeaseCreateResponse
		leaseId          int64
		kv               clientv3.KV
		keepResp         *clientv3.LeaseKeepAliveResponse
		keepRespChan     <-chan *clientv3.LeaseKeepAliveResponse
		putResp          *clientv3.PutResponse
		getResp          *clientv3.GetResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 发起链接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 申请一个lease
	lease = clientv3.NewLease(client)

	if leaseCreatetResp, err = lease.Create(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	leaseId = leaseCreatetResp.ID

	// 自动续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), clientv3.LeaseID(leaseId)); err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepResp == nil {
					fmt.Println("lease is nil")
					goto END
				} else {
					fmt.Println(keepResp.ID)
				}
			}
		}

	END:
	}()

	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "hello", clientv3.WithLease(clientv3.LeaseID(leaseId))); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Revision", putResp.Header.Revision)

	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}

		if len(getResp.Kvs) > 0 {
			fmt.Println(getResp.Kvs)
		}

		time.Sleep(2 * time.Second)
	}
}
