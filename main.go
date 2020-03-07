package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
	"regexp"
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
	CronStErr  string `mapstructure:"cronStErr"`
}

func main() {
	file:="/Users/ck/Documents/go/src/go-cron/cron4/cron3.texa"
	//blueurl := regexp.Match()'', file)
	//fmt.Println(blueurl)
	//
	//[^/]+(?!.*/)
	//one1, _ := regexp.Compile(`/([/][^/]+)$/`)
	//index1 := one1.FindIndex([]byte(file))

	reg := regexp.MustCompile(`(^/)$`)
	fmt.Println(reg.FindAllString(file, -1))

	//fmt.Println("FindIndex", index1)


	//flysnowRegexp := regexp.MustCompile(`/([/][^/]+)$/`)
	//params := flysnowRegexp.FindStringSubmatch(file)
	//fmt.Println(params)
	//for _,param :=range params {
	//	fmt.Println(param)
	//}


	//f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	//defer f.Close()
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	f.Write([]byte("黄河之水天上来，奔流到海不复回。\r\n"))
	//}

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

func WriteLog(file, msg string) error {
	if !IsExist(file) {
		return CreateDirFile(file)
	}
	var (
		err error
		f   *os.File
	)

	f, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f, LineFeed+msg)

	defer f.Close()
	return err
}

//CreateDir  文件夹创建
func CreateDirFile(filePath string) error {
	f, err := os.Create(filePath) //传递文件路径
	if err != nil {           //有错误
		fmt.Println("err = ", err)
		return err
	}
	defer f.Close()
	os.Chmod(filePath, os.ModePerm)
	return nil
}

//IsExist  判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}