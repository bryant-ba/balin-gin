package utils

import "github.com/flipped-aurora/gin-vue-admin/server/utils"

var (
	CloudVerify = utils.Rules{"Name": {utils.NotEmpty()}, "AccessKeyId": {utils.NotEmpty()}, "AccessKeySecret": {utils.NotEmpty()}, "Platform": {utils.NotEmpty()}}
)
