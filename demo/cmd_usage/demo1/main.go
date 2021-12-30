package main

import (
	"fmt"
	"os/exec"
)

func main() {

	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	cmd = exec.Command("~/bin/bash", "-C", "sleep 5;ls -l")

	// 执行子进程并捕获输出
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
