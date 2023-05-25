# FXDEMO

该项目是`fx-tool`(https://github.com/luoruofeng/fx-tool)项目的模版,fx-tool是快速搭建go项目的脚手架。

该项目使用了UBER开源的Fx框架作为实例管理以及依赖注入框架。

## 分支说明
项目中的`init_project`, `http_server`, `register_handler`, `many_handlers`, `logger`, `decouple_registration`, `many_handlers分支`为Fx的教学示例。    
`basic分支`最fx-tool脚手架所搭建的基础版本项目。   
`其他分支`为常用三方中间件与basic基础版本项目融合的项目。

Fx的基础使用方法还可以参考官网docs(https://uber-go.github.io/fx/get-started/)。

## basic分支提供的功能
* **HTTP Server Router** gorilla/muxl路由支持,可自定义middleware。
* **Logger** UBER/zap日志系统。
* **Flag** go原生的flag支持。
* **Config File** 支持读取配置文件 
* 将来会开辟不同分支来将常用的三方中间件与Fx框架融合。

## 编译
若需要二开可以进行编译
```shell
# 使用makefile可以对linux和windows进行编译
make build-linux
make build-windows
```

## 运行
```shell
## 有两种方式启动项目
make run #这种方式请确保不是在windows环境，应该运行前会编译成linux版本。
go run  . -cnf="./conf/conf.json" #该方式linux和windows皆可使用。
```

```shell
## 测试是否启动成功访问以下URL测试
curl -X POST -d 'hello' http://localhost:8080/echo
curl -X POST -d 'gopher' http://localhost:8080/hello
```