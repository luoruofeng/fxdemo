# 该项目是fx-tool(https://github.com/luoruofeng/fx-tool)项目的模版,fx-tool是快速搭建包含Fx框架的go项目的脚手架。

## 基础分支包含的内容
* **HTTP Server Router** gorilla/muxl路由支持,可自定义middleware。
* **Logger** UBER/zap日志系统。
* **Flag** go原生的flag支持。
* **Config File** 支持读取配置文件 


# 项目启动后可访问
```shell
make run
# or
go run  . -cnf="./conf.json"
```

```shell
$ curl -X POST -d 'hello' http://localhost:8080/echo
hello

$ curl -X POST -d 'gopher' http://localhost:8080/hello
Hello, gopher
```