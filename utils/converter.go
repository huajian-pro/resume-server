package utils

import "net/http"

// Handler 转换程序：用于某些无非使用 Fasthttp 的场景
// net/http 与 Fiber 请求的相互转换适配器
// 将 net/http.HandlerFunc 转换为 Fiber.Handler
// 使用："github.com/gofiber/adaptor/v2" 中的 adaptor.HTTPHandler()
// 传入 Handler() 并在 Handler() 传入 http 的请求
// 详细内容请查看其项目文档
func Handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}
