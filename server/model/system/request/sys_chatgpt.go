package request

import (
	"server/model/common/request"
	"server/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
