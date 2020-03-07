package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
	"time"
	"github.com/spf13/viper"
)

var Config config

type config struct {
	CronTabContents []cronTabContent `mapstructure:"cronConfigs"`
}

const (
	LineFeed = "\r\n"
)

type cronTabContent struct {

	Path       string `mapstructure:"path"`
	CronTime   string `mapstructure:"cronTime"`
	CronStdOut string `mapstructure:"cronStdOut"`
}

func main() {

	viper.AddConfigPath("./config")
	viper.SetConfigName("cron")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("Fatal error on set value to Config: %s \n", err))
	}




	//file:="/Users/ck/Documents/go/src/go-cron/cron4/cron3/cron3.text"

	//res:=WriteLog(file,"ok")
	//fmt.Println(res)

	//fmt.Println(os.Getenv("GOPATH"))
	//c := cron.New()
	//c.AddFunc("* * * * * *", func() { fmt.Println("Every hour on the half hour") })
	//c.AddFunc("@every 5s",test1)

	//c.AddFunc("@every 5s",test2)
	//c.Start()

	//select {
	//
	//}
	//c.Stop()

}

func test1() {
	order:="echo 'ojbk1' >> cron1.text"
	ttt:="3s"
	fmt.Println(ttt)

	res:=string(Cmd(order,true))

	fmt.Println(res)
}

func test2() {
	order:="ls -al"
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

func GetPath(file string) string{
	file_byte:=[]byte(file)
	var stop int

	for x,s:=range file_byte{
		if s== 47{
			stop=x
		}
	}

	return string(file_byte[0:stop])
}

func WriteLog(file, msg string) error {
	path:=GetPath(file)
	if !IsExist(path) {
		 CreateDirFile(path)
	}

	var (
		err error
		f   *os.File
	)

	timeUnix:=time.Now().Unix()
	formatTimeStr:=time.Unix(timeUnix,0).Format("2006-01-02 15:04:05")

	f, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f, formatTimeStr+LineFeed+msg+LineFeed)

	 f.Close()
	return err
}

func CreateDirFile(filePath string) error {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(filePath, os.ModePerm)
	return nil
}


func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}