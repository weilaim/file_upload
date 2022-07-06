package middleware

import (

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	// config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type","Access-Control-Allow-Origin", "Cookie"}
	config.AllowHeaders = []string{"*"}
	config.AllowOrigins = []string{"*"}
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}
	// if gin.Mode() == gin.ReleaseMode {
	// 	//生产环境需要配置跨域域名，否则403
	// 	config.AllowOrigins = []string{"http://www.example.com"}
	// } else {
	// 	//测试环境下模糊匹配本地开头的请求
	// 	config.AllowOriginFunc = func(origin string) bool {

	// 		if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
	// 			return true
	// 		}

	// 		if regexp.MustCompile(`^http://172\.22\.206\.19:\d+$`).MatchString(origin) {
	// 			return true
	// 		}

	// 		if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
	// 			return true
	// 		}

	// 		if regexp.MustCompile(`^http://172.18.240.1:\d+$`).MatchString(origin) {
	// 			return true
	// 		}
			
			

	// 		return true
	// 	}
	// }
	config.AllowCredentials = true
	return cors.New(config)
}
