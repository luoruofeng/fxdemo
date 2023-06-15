# FXDEMO

该项目是[fx-tool](https://github.com/luoruofeng/fx-tool)项目的模版。  
`fx-tool`是快速搭建go项目的`脚手架`。     
所生产的项目`模块化，超轻量，少封装`。


<br/>
<br/>

# 使用指南

## 运行
```shell
# 进入项目中。例如：
cd proj_name

#下载所需的库
go mod tidy
```
有如下两种方式启动项目，项目启动成功后会运行一个http服务器进程。
```shell
#该方式linux和windows皆可使用
go run  . -cnf="./conf/conf.json"

#这种方式运行前会编译成linux版本。所有请确保不是在windows环境
make run 
```

```shell
## windows中测试是否运行成功，打开浏览器测试
http://localhost:8080/echo
http://localhost:8080/hello

## linux中测试是否启动成功访问以下URL测试
curl -X POST -d 'hello' http://localhost:8080/echo
curl -X POST -d 'gopher' http://localhost:8080/hello

## ctrl + c停止服务器
```


若想修改项目的相关参数，请修改`/conf/conf.json`文件后，再运行。
```json
{
    "log_level":"日志级别", 
    "log_file":"日志存储位置", 
    "http_addr":"http服务器监听地址", 
    "http_read_over_time":"http读取超时（秒）", 
    "http_write_over_time":"http写入超时（秒）"
}
```
<br>

---

## 教程
 
* 预先说明下面教程会用到的文件夹：  
  `http/handler`是放handler实现的文件夹。  
  `http/middleware`是放http的拦截器实现的文件夹。  
  `fx_opt/srv.go`是http配置的文件夹。  
  `fx_opt/var.go`是注册fx实例配置的文件夹。  

* *`强烈建议`阅读以下内容前，先花10分钟查看[fx基本教程](https://uber-go.github.io/fx/get-started/)*  

* *我们的fx中将会包含以下实例，所以可以给任何一个注册到fx的实例的New函数传递这些实例作为参数。*
```
*http.Server
*go.uber.org/zap.Logger
*github.com/gorilla/mux.Router
*github.com/luoruofeng/fxdemo/onf.Config
```

<br><br>

### `HTTP添加Handler`   
举例说明,只需两步即可完成：   
1. 在`http/handler`中新建文件`hello.go`。  
```go
package handler

import (
	"net/http"

	"go.uber.org/zap"
)

type HelloHandler struct {
	log *zap.Logger
}

func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("这里可以使用h中的任何属性，这是使用了log属性记录日志")
	w.Write([]byte("hello"))
}
```

* *HelloHandler*定义结构体,可以包含任何需要的其他实例(fx中注册的其他任何provider实例)  
* *NewHelloHandler*是fx的provider的构造器需的New构造方法，返回的*HelloHandler将会注册到fx中，参数为所需的任何实例（fx中注册的实例）。
* *Pattern*是路由url配置。    
* *ServeHTTP*是handler的具体实现

<br>

2. 打开`fx_opt/srv.go`,在`Setup`方法的`handlerProv`变量一行中的`fxhttp.AllAsRoute`函数中添加新的参数`handler.NewHelloHandler`即可(`handler.NewHelloHandler`是上面的代码创建的函数)。
```go
func (f *FxSrv) Setup() {
	handlerProv := fx.Provide(
		fxhttp.AllAsRoute(handler.NewEchoHandler, handler.NewHelloHandler)...,
	)
//其他代码
```
<br>

3. 启动服务验证结果  
浏览器中输入http://localhost:8080/hello即可访问。

<br>
<br>

### `HTTP添加拦截器`
举例说明，若想给http添加新的日志拦截器（middleware）只需以下两步。
1. 在`http/middleware`文件夹中新建`log2.go`，内容如下：  
```go 
package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type LogMiddleware2 struct {
	logger *zap.Logger
}

func NewLogMiddleware2(logger *zap.Logger) *LogMiddleware2 {
	return &LogMiddleware2{logger: logger}
}

func (l *LogMiddleware2) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.logger.Info("正在执行log2拦截器")
		next.ServeHTTP(w, r)
        l.logger.Info("log2拦截器即将执行完成")

	})
}

```
* *LogMiddleware2 struct*定义结构体，结构体中定义我们所需要的实例(该实例已经在fx中注册)，例如我们这里需要*zap.Logger。
* *NewLogMiddleware2*定义上面结构体的New函数，参数为我们所需要的实例(该实例已经在fx中注册)，例如我们这里需要*zap.Logger。
* *Middleware*定义拦截器，该拦截器会在每次HTTP请求中被调用，*next.ServeHTTP(w, r)*为handler方法本身。可以在该方法上下去编写拦截器逻辑。

<br>

2. 打开`fx_opt/srv.go`,在`Setup`方法的`middlewareProv`变量的`fxhttp.AllAsMiddleware`数中添加新的参数`middleware.NewLogMiddleware2`即可(*middleware.NewLogMiddleware2*是上面的代码创建的函数)。
```go
func (f *FxSrv) Setup() {
    //其他代码
	middlewareProv := fx.Provide(
		fxhttp.AllAsMiddleware(middleware.NewLogMiddleware, middleware.NewLogMiddleware2)...,
```

<br>

3. 启动服务验证结果    
浏览器中输入http://localhost:8080/hello即可访问。   
```shell
# 项目控制台将会输出以下日志
...
{"level":"info","message":"正在执行log2拦截器"}
...
{"level":"info","message":"log2拦截器即将执行完成"}
...
```

<br>
<br>

### `添加新的实例到fx`
举例说明，若想在fx中创建实例只需以下两步。
1. 在`srv/`文件夹中新建`srv1.go`，内容如下：  
```go
package srv

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Abc struct {
	logger *zap.Logger
}

func NewAbc(lc fx.Lifecycle, logger *zap.Logger) Abc {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Abc开始构建")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Abc开始销毁")
			return nil
		},
	})

	return Abc{logger: logger}
}
```
* *Abc*是该案例中的结构体，该结构体产生的实例最终会被包含到fx中。也可以被fx中的其他实例使用。
* *NewAbc*函数的返回值Abc会创建到fx中。  
其中参数是我们已经注册到了fx中的实例。当然我们也可以不同传递任何参数，如果你用不到这些参数。  
传递第一个参数*lc fx.Lifecycle*的原因是因为要用*OnStart*和*OnStop*，如果不需要在该实例创建和销毁时候调用这两个函数，第一个参数可以不用传递。   
第二个参数*logger \*zap.Logger*是为了log日志记录传递，不需要也可以不传递。  

<br>

2. 打开`fx_opt/var.go`,在`ConstructorFuncs`变量中添加新的参数`srv.NewAbc,`即可(*srv.NewAbc*是上面的代码创建的函数)。
```go
var ConstructorFuncs = []interface{}{
	srv.NewAbc,
```
* 如果*NewAbc*函数使用了`lc fx.Lifecycle`，则需要在`fx_opt/var.go`的`InvokeFuncs`变量中添加`func(srv.Abc) {}`函数（该函数将注册到fx的invoke中）。如果没有使用`lc fx.Lifecycle`则不用添加。
```go
var InvokeFuncs = []interface{}{
	func(srv.Abc) {}
//其他代码
```
<br>

3. 启动服务器然后终止服务器，查看服务器控制台如下输出表示成功。
```
...
{"level":"info","message":"Abc开始构建"}
...
{"level":"info","message":"Abc开始销毁"}
...
```


<br>
<br>

### `添加新的实例到fx（结构体带不属于fx中实例的属性，New函数带不属于fx中实例的参数）`
举例说明，若想在fx中创建实例只需以下两步。
1. 在`srv/`文件夹中新建`srv2.go`，内容如下：  
```go
package srv

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Abc2 struct {
	logger  *zap.Logger
	Content string
}

func NewAbc2(lc fx.Lifecycle, logger *zap.Logger, content string) Abc2 {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Abc2开始构建", zap.String("content", content))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Abc2开始销毁")
			return nil
		},
	})

	return Abc2{logger: logger, Content: content}
}
```
* 和上一个案例基本相同，唯一不同之处是*Abc2*结构体中包含了不来自于fx的参数*Content string*,它同样来源于*NewAbc2*函数的参数*content string*，而传递参数的不同之处在下一步中体现。

<br>

2. 打开`fx_opt/var.go`,在`ConstructorFuncs`变量中添加两个新的数据项`fx.Annotate`即可(*srv.NewAbc2*是上面的代码创建的函数)。
```go
var ConstructorFuncs = []interface{}{
	//其他代码
	fx.Annotate(
		func() string {
			return "这是Abc2的Content参数"
		},
		fx.ResultTags(`name:"abc2content"`),
	),

	fx.Annotate(
		srv.NewAbc2,
		fx.ParamTags(``, ``, `name:"abc2content"`),
	),
```
* 上面第一个*fx.Annotate*声明将会作为*srv.NewAbc2*函数的*content string*参数，进行依赖注入。   
* 上面第一个*fx.Annotate*中的*fx.ResultTags(\`name:"abc2content"\`)*会匹配第二个*fx.Annotate*对象的fx.ParamTags(\`\`, \`\`, \`name:"abc2content"\`)的第三项。
* 第二个*fx.Annotate*中的参数\`\`, \`\`, \`name:"abc2content"\`按顺序匹配了*NewAbc2(lc fx.Lifecycle, logger *zap.Logger, content string)*的三个参数，第一项，第二项fx不需要*name*匹配，第三项则是第一个*fx.Annotate*实例的第一个参数的内容。

<br>

* 如果*NewAbc2*函数使用了`lc fx.Lifecycle`，则需要在`fx_opt/var.go`的`InvokeFuncs`变量中添加`func(srv.Abc2) {}`函数（该函数将注册到fx的invoke中）。如果没有使用`lc fx.Lifecycle`则不用添加。
```go
var InvokeFuncs = []interface{}{
//其他代码
	func(srv.Abc2) {}
//其他代码
```
<br>

3. 启动服务器然后终止服务器，查看服务器控制台如下输出表示成功。
```
...
{"level":"info","message":"Abc2开始构建","content":"这是Abc2的Content参数"}
...
{"level":"info","message":"Abc2开始销毁"}
...
```

<br>
<br>

### `添加新的实例到fx（结构体包含刚在fx中注册的实例作为属性，New函数将刚在fx中注册的实例作为参数）`
举例说明，将上一个例子中的Abc2作为`NewAbc3`的参数,若想在fx中创建实例只需以下两步。
1. 在`srv/`文件夹中新建`srv2.go`，内容如下：  
```go
package srv

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Abc3 struct {
	logger  *zap.Logger
	Abc2   Abc2
}

func NewAbc3(lc fx.Lifecycle, logger *zap.Logger, abc2 Abc2) Abc3 {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Abc3开始构建", zap.Any("abc2", abc2))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Abc3开始销毁")
			return nil
		},
	})

	return Abc3{logger: logger, Abc2: abc2}
}
```
* 和上两个案例基本相同，唯一不同之处是*Abc3*结构体中包含了刚在fx中注册的"Abc2"实例,它同样来源于*NewAbc3*函数，而传递参数的不同之处在下一步中体现。

<br>

2. 打开`fx_opt/var.go`,在`ConstructorFuncs`变量中添加新的参数`srv.NewAbc3,`即可(*srv.NewAbc*是上面的代码创建的函数)。
```go
var ConstructorFuncs = []interface{}{
	//其他代码
	srv.NewAbc3,
```
* 如果*NewAbc3*函数使用了`lc fx.Lifecycle`，则需要在`fx_opt/var.go`的`InvokeFuncs`变量中添加`func(srv.Abc2) {}`函数（该函数将注册到fx的invoke中）。如果没有使用`lc fx.Lifecycle`则不用添加。
```go
var InvokeFuncs = []interface{}{
//其他代码
	func(srv.Abc3) {}
//其他代码
```
<br>

3. 启动服务器然后终止服务器，查看服务器控制台如下输出表示成功。
```
...
{"level":"info","message":"Abc3开始构建","abc2":{"Content":"这是Abc2的Content参数"}}
...
{"level":"info","message":"Abc3开始销毁"}
...
```

<br>
<br>

# 其他说明

## 项目依赖的三方库  
UBER开源的[fx](https://github.com/uber-go/fx/)框架作为实例管理以及依赖注入框架    
日志管理使用了UBER开源日志框架的[zap](https://github.com/uber-go/zap/)    
HTTP服务使用了[gorilla/mux](github.com/gorilla/mux)  

<br>

## 项目结构
`用户需要关注的文件夹以 * 提示`
```shell
├── component *存放集成到项目的三方模块*
├── conf
│   ├── conf.json *配置文件*
│   └── model.go 
├── http
│   ├── handler *http的handler文件夹*
│   │   ├── echo.go *案例*
│   │   └── ...
│   └── middleware *http请求的拦截器文件夹*
│       ├── log.go *案例*
│       └── ...
└── srv *需要注册到fx中的服务*
    ├── srv1.go *案例*
    ├── ...
├── fx_opt fx操作文件夹
│   ├── component
│   │   ├── invoke fx的invoke
│   │   │   └── router.go
│   │   └── provide fx的provide
│   │       ├── conf.go 
│   │       ├── http
│   │       │   ├── http_middleware.go
│   │       │   ├── http_mux.go
│   │       │   ├── http_route.go
│   │       │   └── http_server.go
│   │       └── logger.go
│   │       └── context.go
│   ├── srv.go fx服务
│   └── var.go fx变量
├── LICENSE 项目开源声明
├── main.go 主函数
├── cmd 项目启动命令行
│   └── config.go
├── Makefile 
├── README.md 
├── go.mod
├── go.sum
```

<br>

## master分支提供的功能
* **HTTP Server Router** gorilla/muxl路由支持,可自定义middleware。
* **Logger** UBER/zap日志系统。
* **Flag** go原生的flag支持。
* **Config File** 支持读取配置文件 
* 将来会开辟不同分支来将常用的三方中间件与Fx框架融合。

<br>

## 编译
若需要二开可以进行编译
```shell
# 使用makefile可以对linux和windows进行编译
make build-linux
make build-windows
```

<br>

## 分支说明
项目中的`init_project`, `http_server`, `register_handler`, `many_handlers`, `logger`, `decouple_registration`, `many_handlers分支`为Fx的教学示例。    

`master分支`最fx-tool脚手架所搭建的基础版本项目。   

Fx的基础使用方法还可以参考官网docs(https://uber-go.github.io/fx/get-started/)。