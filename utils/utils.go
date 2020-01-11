package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

const (
	success        = "0"
	defaultMessage = "Success"
)
type ResponseBean struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *ResponseBean {
	return &ResponseBean{Data: data,Code:success,Message: defaultMessage}
}

func ToJson(d interface{}) string {
	jsonByte, _ := json.Marshal(d)
	return string(jsonByte)
}

func FromJson(jsonStr string, s interface{}) {
	err := json.Unmarshal([]byte(jsonStr), s)
	if nil != err {
		logs.Error("解析json出错")
		return
	}
}

func ConvertBytesToJson(jsonBytes []byte, s interface{}) {
	err := json.Unmarshal(jsonBytes, s)
	if nil != err {
		logs.Error("解析json出错")
		return
	}
}