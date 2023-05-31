package handler

import (
	"net/http"

	"go.uber.org/zap"
)

// 定义结构体,可以包含任何需要的其他实例(fx中注册的其他任何provider实例)
type HelloHandler struct {
	log *zap.Logger //这里我们举例使用了*zap.Logger实例
}

// fx的provider的构造器需要新的New构造方法，返回的*HelloHandler将会注册到fx中，参数为所需的任何实例（fx中注册的实例）。
func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}

// 路由url配置
func (*HelloHandler) Pattern() string {
	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("这里可以使用h中的任何属性，这是使用了log属性记录日志")
	w.Write([]byte("hello"))
}
