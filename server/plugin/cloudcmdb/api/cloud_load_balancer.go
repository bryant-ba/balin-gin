package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	cloudcmdbreq "server/plugin/cloudcmdb/model/cloudcmdb"
	"server/utils"
)

type CloudLoadBalancerApi struct{}

// @Tags CloudLoadBalancerApi
// @Summary 同步负载均衡
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /cloudcmdb/loadBalancer/sync [post]

func (l *CloudLoadBalancerApi) CloudLoadBalancerSync(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := cloudLoadBalancerService.SyncLoadBalancer(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("同步操作失败!", zap.Error(err))
		response.FailWithMessage("同步操作失败", c)
	} else {
		response.OkWithMessage("同步操作成功, 数据异步处理中, 请稍后!", c)
	}
}

// @Tags CloudLoadBalancerApi
// @Summary 负载均衡列表
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /cloudcmdb/loadBalancer/list [post]

func (l *CloudLoadBalancerApi) CloudLoadBalancerList(c *gin.Context) {
	var pageInfo cloudcmdbreq.SearchLoadBalancerParams
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	List, total, err := cloudLoadBalancerService.List(pageInfo.LoadBalancer, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     List,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags CloudLoadBalancerApi
// @Summary 负载均衡Tree数据
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /cloudcmdb/loadBalancer/tree [post]

func (l *CloudLoadBalancerApi) CloudLoadBalancerTree(c *gin.Context) {
	var pageInfo cloudcmdbreq.SearchCloudPlatformParams
	_ = c.ShouldBindJSON(&pageInfo)
	list, err := cloudLoadBalancerService.LoadBalancerTree(pageInfo.CloudPlatform, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("获取目录树失败!", zap.Error(err))
		response.FailWithMessage("获取目录树失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
