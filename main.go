package main

import (
	"github.com/gin-gonic/gin"
	"wtf-credential/configs"
	"wtf-credential/handle"
	"wtf-credential/middleware"
)

func main() {
	configs.Config()
	configs.ParseConfig("./configs/config.yaml") // 加载 configs 目录中的配置文件
	configs.NewRedis()
	//daos.InitPostgres()
	//go tasks.GetContributorsJob()
	r := gin.Default()
	route(r)
	r.Use(middleware.Core())
	err := r.Run(":" + configs.Config().Port)
	if err != nil {
		return
	}
}

func route(r *gin.Engine) {
	public := r.Group("/api/v1")
	{
		public.GET("/ping", handle.GetPing)                      // 不鉴权的测试接口 ✅
		public.POST("/auth/nonce", handle.GenerateNonce)         //获取nonce ✅❌
		public.POST("/auth/github_login", handle.GithubLogin)    //github登陆✅❌
		public.POST("/auth/login", handle.Login)                 //钱包登陆✅❌
		public.POST("/contributors", handle.GetContributorsList) //全部贡献者列表❌❌github列表没拿到信息
	}

	private := r.Group("/api/v1")
	private.Use(middleware.JWTAuthMiddleware())
	{
		private.GET("/user/wallet/", handle.GetUserWallet)       //获取钱包✅❌
		private.POST("/user/wallet/bind", handle.BindWallet)     //绑定钱包✅❌
		private.POST("/user/wallet/change", handle.ChangeWallet) //改变钱包✅❌
	}
}
