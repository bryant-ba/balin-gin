package utils

import "server/utils"

var (
	CloudVerify = utils.Rules{"Name": {utils.NotEmpty()}, "AccessKeyId": {utils.NotEmpty()}, "AccessKeySecret": {utils.NotEmpty()}, "Platform": {utils.NotEmpty()}}
)
