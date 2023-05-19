# 对FX的使用

## 项目包括的内容
* **HTTP Server Router** gorilla/muxl路由支持,可自定义middleware。
* **Logger** UBER/zap日志系统。
* **Flag** go原生的flag支持。
* **Config File** 支持读取配置文件 
* **Consul** 注册服务


# 项目启动后可访问
```shell
make run
```

```shell
$ curl -X POST -d 'hello' http://localhost:8080/echo
hello

$ curl -X POST -d 'gopher' http://localhost:8080/hello
Hello, gopher
```