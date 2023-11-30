package router

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
	"server/plugin/cloudcmdb/api"
)

type CloudRegionRouter struct {
}

func (c *CloudRegionRouter) InitCloudRegionRouter(Router *gin.RouterGroup) {
	cloudRegionRouter := Router.Group("cloud_region").Use(middleware.OperationRecord())
	cloudRegionApi := api.ApiGroupApp.CloudRegionApi
	{
		cloudRegionRouter.POST("syncRegion", cloudRegionApi.CloudPlatformSyncRegion) // 同步区域信息
	}
}
