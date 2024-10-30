package main

import (
	"github.com/gin-gonic/gin"
	"wtf-credential/configs"
	"wtf-credential/daos"
	"wtf-credential/handle"
	"wtf-credential/middleware"
	"wtf-credential/tasks"
)

func main() {
	configs.Config()
	configs.ParseConfig("./configs/config.yaml") // 加载 configs 目录中的配置文件
	configs.NewRedis()
	daos.InitPostgres()
	go tasks.GetContributorsJob()
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(middleware.CORSMiddleware())
	route(r)
	err := r.Run(":" + configs.Config().Port)
	if err != nil {
		return
	}
}

func route(r *gin.Engine) {
	public := r.Group("/api/v1")
	{
		public.GET("/ping", handle.GetPing)                                                   //不鉴权的测试接口 ✅
		public.POST("/auth/nonce", handle.GenerateNonce)                                      //获取nonce ✅❌
		public.POST("/auth/github_login", handle.GithubLogin)                                 //github登陆✅❌
		public.POST("/auth/login", handle.Login)                                              //钱包登陆✅❌
		public.POST("/contributors", handle.GetContributorsList)                              //全部贡献者列表✅❌
		public.GET("/courses", handle.GetAllCourse)                                           //获取课程列表✅❌
		public.GET("/courses/type", handle.GetCoursesByType)                                  //获取课程列表按type分类✅
		public.GET("/courses/stats", handle.GetStatistics)                                    //获取数值统计信息✅
		public.GET("/courses/:course_id", handle.GetCourseInfo)                               //根据课程id获取课程信息❌❌
		public.GET("/courses/:course_id/quizzes", handle.GetCourseQuizzes)                    //根据课程 ID 获取课程检测（quiz）列表❌❌
		public.GET("/courses/:course_id/user_lessons/:lesson_id", handle.GetUserCourseLesson) //根据课程 ID 和单元 ID 获取用户的课程单元信息❌❌

	}

	private := r.Group("/api/v1")
	private.Use(middleware.JWTAuthMiddleware())
	{
		private.GET("/user/wallet/", handle.GetUserWallet)             //获取钱包✅❌
		private.POST("/user/wallet/bind", handle.BindWallet)           //绑定钱包✅❌
		private.POST("/user/wallet/change", handle.ChangeWallet)       //改变钱包✅❌
		private.POST("/user/wallet/unbindWallet", handle.UnbindWallet) //取消绑定钱包✅                         //取消绑定钱包
		private.GET("/user/profile", handle.GetProfileByUserID)        //获取用户信息✅
	}
}
