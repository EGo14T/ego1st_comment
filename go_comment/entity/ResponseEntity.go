package entity

import "net/http"

type Code struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ReturnMsg --------对需要返回的信息进行封装，方便对数据进行进一步处理
type ReturnMsg struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//--------对需要返回的信息进行赋值，并以结构体返回
func responseFunc(code Code, data interface{}) *ReturnMsg {
	rm := new(ReturnMsg)
	rm.Code = code.Code
	rm.Message = code.Message
	rm.Data = data
	return rm
}

// ResOk 返回默认成功信息200
func ResOk(data interface{}) *ReturnMsg {
	code := Code{Code: http.StatusOK, Message: "Success"}
	rm := new(ReturnMsg)
	rm.Code = code.Code
	rm.Message = code.Message
	rm.Data = data
	return rm
}

// ResError 返回默认错误500
func ResError(data interface{}) *ReturnMsg {
	code := Code{Code: http.StatusInternalServerError, Message: "系统错误"}
	rm := new(ReturnMsg)
	rm.Code = code.Code
	rm.Message = code.Message
	rm.Data = data
	return rm
}
