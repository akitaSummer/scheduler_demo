package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {

	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheulerTable map[string]*CronJob
	)

	scheulerTable = make(map[string]*CronJob)

	now = time.Now()

	expr = cronexpr.MustParse("*/5 * * * * * *")

	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}

	scheulerTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/5 * * * * * *")

	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}

	scheulerTable["job2"] = cronJob

	// 调度协程，谁过期了就执行谁
	go func() {
		var (
			now     time.Time
			cronJob *CronJob
			jobName string
		)
		for {
			now = time.Now()

			for jobName, cronJob = range scheulerTable {
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					go func(jobName string) {
						fmt.Println(jobName)
					}(jobName)
				}

				cronJob.nextTime = cronJob.expr.Next(now)
			}

			select {
			case <-time.NewTimer(100 * time.Millisecond).C:
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(100 * time.Minute)
}
