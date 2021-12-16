package main

import (
	"fmt"
	"time"

	"context"

	"github.com/coreos/etcd/clientv3"
	storagepb "github.com/coreos/etcd/storage/storagepb"
)

func main() {
	var (
		config             clientv3.Config
		client             *clientv3.Client
		err                error
		kv                 clientv3.KV
		getResp            *clientv3.GetResponse
		watchStartRevision int64
		watcher            clientv3.Watcher
		watchRespChan      <-chan clientv3.WatchResponse
		watchResp          clientv3.WatchResponse
		event              *storagepb.Event
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

	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job1", "hello")
			kv.Delete(context.TODO(), "/cron/jobs/job1")
			time.Sleep(1 * time.Second)
		}
	}()

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1"); err != nil {
		fmt.Println(err)
		return
	}

	if len(getResp.Kvs) != 0 {
		fmt.Println(getResp.Kvs[0].Value)
	}

	watchStartRevision = getResp.Header.Revision + 1

	watcher = clientv3.NewWatcher(client)

	watchRespChan = watcher.Watch(context.TODO(), "/cron/jobs/job1", clientv3.WithRev(watchStartRevision))

	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case storagepb.PUT:
				fmt.Println("PUT", event.Kv.Value, event.Kv.ModRevision)
			case storagepb.DELETE:
				fmt.Println("DELETE", event.Kv.ModRevision)
			}
		}
	}
}
