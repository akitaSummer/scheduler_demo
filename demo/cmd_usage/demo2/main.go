package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {

	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)

	resultChan = make(chan *result, 1000)

	// context: chan byte
	// cancelFunc: close(chan byte)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)
		// select { case <- ctx.Done() }
		// kill pid
		cmd = exec.CommandContext(ctx, "~/bin/bash", "-C", "sleep 5;ls -l;")

		output, err = cmd.CombinedOutput()

		resultChan <- &result{
			output: output,
			err:    err,
		}
	}()

	time.Sleep(1 * time.Second)
	cancelFunc()

	res = <-resultChan

	fmt.Println(res.err, string(res.output))
}
