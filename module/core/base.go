package core

// StockTao ==> Global Response Object
type StockTao struct {
	Code uint        `json:"code"` // 响应状态码
	Msg  string      `json:"msg"`  // 响应消息
	Data interface{} `json:"data"` // 响应数据
}

type code uint32

const (
	// Success => success code
	Success code = 200
)
