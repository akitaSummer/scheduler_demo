package main

import (
	"flag"
	"fmt"
	"runtime"
	"scheduler_demo/master"
)

var (
	confFile string
)

func initArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	var (
		err error
	)

	initArgs()

	initEnv()

	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	return

ERR:
	fmt.Println(err)

}
