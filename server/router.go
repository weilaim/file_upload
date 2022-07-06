package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/weilaim/blog-api/api"
	"github.com/weilaim/blog-api/middleware"
)

//NewRouter 路由配置
func NewRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Loggers())

	//路由中间件
	// r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))

	r.Use(middleware.Cors())
	// r.Use(middleware.CurrentUser())

	//路由
	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		//TODO

		auth.DELETE("user/logout", api.UserLogout)
		//视频操作
		//添加视频
		auth.POST("videos", api.CreateVideo)
		//获取视频
		auth.PUT("video/:id", api.UpdateVideo)
		auth.DELETE("video/:id", api.DeleteVideo)

		//创建文件
		auth.POST("files",api.CreateFile)
		auth.PUT("file/:id",api.UpdateFile)
		auth.DELETE("file/:id",api.DeleteFile)

		//创建FileAcc
		auth.POST("fileacc",api.CreateFileAcc)
		auth.PUT("fileacc",api.UpdateFileAcc)
		auth.DELETE("fileacc/:id",api.DeleteFileAcc)
		//创建小情书
		auth.POST("love", api.CreateLove)

		//获取用户列表
		auth.GET("users", api.ShowUser)
		//获取用户信息
		auth.GET("user/info", api.UserInfo)
		// // oss toke
		// auth.POST("upload/token", api.UploadToken)

		//用户操作

		//修改用户
		//删除用户

		//创建行程
	}
	v1 := r.Group("api/v1")
	{
		// oss toke
		v1.POST("upload/token", api.UploadToken)
		v1.POST("upload/qiniu", api.QiniuToken)
		v1.DELETE("upload/qiniu", api.QiniuDelet)

		//like
		v1.POST("like", api.CreateLike)
		//collection 收藏
		v1.POST("collection",api.Collection)

		v1.GET("ping", api.Ping)
		// // User Routing
		v1.GET("user/me", api.UserMe)
		//获取视频
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		//排行榜
		v1.GET("rank/daily", api.DailyRank)

		//

		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// // 用户登录
		v1.POST("user/login", api.UserLogin)

		//get openid && login
		v1.POST("wxlogin", api.GetOpenId)

		//获取小情书列表
		v1.GET("loves", api.ListLove)
		//查询单个图集的数据
		v1.GET("love", api.GetLove)
		//file
		v1.GET("files",api.ListFiles)
		v1.GET("file",api.GetFile)
		//fileacc
		v1.GET("fileaccs",api.ListFileAccs)
		v1.GET("fileacc/:id",api.GetFileAcc)
		v1.POST("fileauth",api.FileAccAuth)
	}
	r.GET("/", api.Ping)
	return r
}
