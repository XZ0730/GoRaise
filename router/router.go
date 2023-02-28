package router

import (
	"Raising/api"
	"Raising/middleware"

	_ "Raising/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("api/v1")
	{
		//测试
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("/register", api.UserRegister)
		v1.POST("/login", api.UserLogin)
		v1.POST("/phone", api.SendPhoneNum)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{

			authed.PUT("user", api.UserUpdate)
			authed.GET("user", api.UserGet)
			authed.POST("user/upload_ava", api.UploadAvatar)
			authed.POST("user/email", api.SendEmail)        //发送邮件
			authed.GET("user/email/:token", api.ValidEmail) //验证邮箱 ：绑定；解绑；修改密码
			authed.POST("user/pay", api.PayDown)            //支付订单并删除
			authed.GET("user/audit", api.GetUserAudit)      //获取用户正在审核的项目

			authed.POST("project", api.Create_Project)      //创建筹资项目
			authed.GET("project/:pid", api.GetProject_Pid)  //根据pid获取详细信息
			authed.GET("project", api.GetProject)           //获取projectlist 积分推送机制
			authed.GET("project/search", api.SearchProject) //获取name来搜索某一项目
			//用户先提交订单，然后选择是否支付
			//如果支付则直接删除订单，或者选择取消订单来删除订单
			authed.POST("order", api.Create_Order)  //提交订单
			authed.GET("order", api.GetOderList)    //得到当前用户的所有订单
			authed.DELETE("order", api.DeleteOrder) //取消订单

			Admin := authed.Group("/admin")
			Admin.Use(api.IsAdmin)
			{ //管理员操作
				//获取审核列表：?page=1默认搜索全部分页 (简要信息：title 作者 创建时间)按照时间排序
				//"/pid" 获取pid 跳转单个页面 展示info title author img
				//"?name=" 搜索项目并分页
				// Admin.GET("/audit")
				Admin.GET("/audit/:pid", api.Get_AuditByID) //项目pid获取具体项目信息
				Admin.GET("/audit", api.Get_Audit)          //得到项目审核列表
				Admin.POST("/audit", api.Audit_Project)     //pid+ispass 审核项目
				Admin.DELETE("", api.Delete_Project)        //传入pid query删除某项目  / 删除所有项目

			}
		}

	}

	return r
}
