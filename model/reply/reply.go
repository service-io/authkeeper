// Package reply
// @author tabuyos
// @since 2023/7/19
// @description model
package reply

import (
	"deepsea/helper/runerror"
)

type Reply struct {
	// 错误码
	Code *int `json:"code"`
	// 执行状态
	State *bool `json:"state"`
	// 响应消息
	Message *string `json:"message"`
	// 响应负载
	Payload *any `json:"payload"`
}

func New() *Reply {
	return &Reply{}
}

func of(code int, state bool, message string, payload any) *Reply {
	return &Reply{
		Code:    &code,
		State:   &state,
		Message: &message,
		Payload: &payload,
	}
}

func Ok() *Reply {
	return of(runerror.OK, true, "成功", nil)
}

func OkPayload(payload any) *Reply {
	return of(runerror.OK, true, "成功", payload)
}

func OkMessage(message string) *Reply {
	return of(runerror.OK, true, message, nil)
}

func Failed() *Reply {
	return of(runerror.Failed, false, "失败", nil)
}

func FailedPayload(payload any) *Reply {
	return of(runerror.Failed, false, "失败", payload)
}

func FailedMessage(message string) *Reply {
	return of(runerror.Failed, false, message, nil)
}

func (receiver *Reply) WithCode(code int) *Reply {
	receiver.Code = &code
	return receiver
}

func (receiver *Reply) WithState(state bool) *Reply {
	receiver.State = &state
	return receiver
}

func (receiver *Reply) WithMessage(message string) *Reply {
	receiver.Message = &message
	return receiver
}

func (receiver *Reply) WithPayload(payload any) *Reply {
	receiver.Payload = &payload
	return receiver
}
