package main

import (
	"fmt"
	"time"
	"github.com/robfig/cron"
	"os"
	log "github.com/cihub/seelog"
)

var ttt = 1

var Config config

type config struct {
	CronTabContents []cronTabContent `mapstructure:"cronConfigs"`
}

type cronTabContent struct {
	Path       string `mapstructure:"path"`
	CronTime   string `mapstructure:"cronTime"`
	CronStdOut string `mapstructure:"cronStdOut"`
	CronStErr  string `mapstructure:"cronStErr"`
}

func main() {
	fmt.Println(os.Getenv("GOPATH"))
	c := cron.New()
	c.AddFunc("*/5 * * * * ?", func() { log.Info("Every hour on the half hour") })
	c.AddFunc("0/5 * * * * ?", test)
	c.Start()
	time.Sleep(time.Minute)
	fmt.Println("aaa")

}

func test() {

	fmt.Printf("test=%d\n", ttt)
	ttt++
}
