package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/cloudcmdb/service"

type ApiGroup struct {
	CloudPlatformApi
	CloudRegionApi
	CloudVirtualMachineApi
	CloudLoadBalancerApi
	CloudRDSApi
}

var ApiGroupApp = new(ApiGroup)

var (
	cloudPlatformService       = service.ServiceGroupApp.CloudPlatformService
	cloudVirtualMachineService = service.ServiceGroupApp.CloudVirtualMachineService
	cloudRegionService         = service.ServiceGroupApp.CloudRegionService
	cloudLoadBalancerService   = service.ServiceGroupApp.CloudLoadBalancerService
	cloudRDSService            = service.ServiceGroupApp.CloudRDSService
)
