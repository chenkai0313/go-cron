package main

import (
	"fmt"
	"github.com/robfig/cron"
	"os"
	"os/exec"
)

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
	//c.AddFunc("* * * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@every 5s",test1)

	c.AddFunc("@every 5s",test2)
	c.Start()

	select {

	}
	c.Stop()

}

func test1() {
	order:="echo 'ojbk' >> cron1.text"
	ttt:="3s"
	fmt.Println(ttt)

	res:=string(Cmd(order,true))

	fmt.Println(res)
}

func test2() {
	order:="echo 'ojbk' >> cron2.text"
	ttt:="5s"
	fmt.Println(ttt)

	res:=string(Cmd(order,true))

	fmt.Println(res)
}

func Cmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	}
}
