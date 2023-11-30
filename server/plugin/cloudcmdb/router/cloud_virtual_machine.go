package router

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
	"server/plugin/cloudcmdb/api"
)

type CloudVirtualMachineRouter struct {
}

func (c *CloudVirtualMachineRouter) InitVirtualMachineRouter(Router *gin.RouterGroup) {
	cloudvirtualMachineRouter := Router.Group("virtualMachine").Use(middleware.OperationRecord())
	cloudvirtualMachineRouterWithoutRecord := Router.Group("virtualMachine")
	cloudvirtualMachineApi := api.ApiGroupApp.CloudVirtualMachineApi
	{
		cloudvirtualMachineRouter.POST("sync", cloudvirtualMachineApi.CloudVirtualMachineSync) // 同步云主机
		cloudvirtualMachineRouter.POST("tree", cloudvirtualMachineApi.CloudVirtualMachineTree) // 目录树
	}

	{
		cloudvirtualMachineRouterWithoutRecord.POST("list", cloudvirtualMachineApi.CloudVirtualMachineList) // 分页获取列表
	}
}
