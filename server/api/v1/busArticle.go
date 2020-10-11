package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags BusArticle
// @Summary 创建BusArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusArticle true "创建BusArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busArticle/createBusArticle [post]
func CreateBusArticle(c *gin.Context) {
	var busArticle model.BusArticle
	_ = c.ShouldBindJSON(&busArticle)
	err := service.CreateBusArticle(busArticle)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags BusArticle
// @Summary 删除BusArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusArticle true "删除BusArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busArticle/deleteBusArticle [delete]
func DeleteBusArticle(c *gin.Context) {
	var busArticle model.BusArticle
	_ = c.ShouldBindJSON(&busArticle)
	err := service.DeleteBusArticle(busArticle)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BusArticle
// @Summary 批量删除BusArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busArticle/deleteBusArticleByIds [delete]
func DeleteBusArticleByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteBusArticleByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BusArticle
// @Summary 更新BusArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusArticle true "更新BusArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busArticle/updateBusArticle [put]
func UpdateBusArticle(c *gin.Context) {
	var busArticle model.BusArticle
	_ = c.ShouldBindJSON(&busArticle)
	err := service.UpdateBusArticle(&busArticle)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags BusArticle
// @Summary 用id查询BusArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusArticle true "用id查询BusArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busArticle/findBusArticle [get]
func FindBusArticle(c *gin.Context) {
	var busArticle model.BusArticle
	_ = c.ShouldBindQuery(&busArticle)
	err, rebusArticle := service.GetBusArticle(busArticle.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rebusArticle": rebusArticle}, c)
	}
}

// @Tags BusArticle
// @Summary 分页获取BusArticle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BusArticleSearch true "分页获取BusArticle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busArticle/getBusArticleList [get]
func GetBusArticleList(c *gin.Context) {
	var pageInfo request.BusArticleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetBusArticleInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
