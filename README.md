### 说明
go-cron小工具，解决众多crontab任务添加问题
### 编译
> go build main.go
### 使用方法
> 在config/cron.yaml中添加执行的定时任务
---
``` 
  -
    path: "echo 'ojbk1111' >> cron1.text"
    cronTime: "@every 5s"
    cronStdOut: "your/path/ojbk11.log"
  -
    path: "echo 'ojbk22' >> cron2.text"
    cronTime: "@every 8s"
    cronStdOut: "your/path/ojbk22.log"
```     
