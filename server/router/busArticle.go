package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBusArticleRouter(Router *gin.RouterGroup) {
	BusArticleRouter := Router.Group("busArticle").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		BusArticleRouter.POST("createBusArticle", v1.CreateBusArticle)   // 新建BusArticle
		BusArticleRouter.DELETE("deleteBusArticle", v1.DeleteBusArticle) // 删除BusArticle
		BusArticleRouter.DELETE("deleteBusArticleByIds", v1.DeleteBusArticleByIds) // 批量删除BusArticle
		BusArticleRouter.PUT("updateBusArticle", v1.UpdateBusArticle)    // 更新BusArticle
		BusArticleRouter.GET("findBusArticle", v1.FindBusArticle)        // 根据ID获取BusArticle
		BusArticleRouter.GET("getBusArticleList", v1.GetBusArticleList)  // 获取BusArticle列表
	}
}
