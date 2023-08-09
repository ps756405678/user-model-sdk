package domain

type CallSdkReq struct {
	Method     string `json:"method"`
	Schema     string `json:"schema"`
	Collection string `json:"collection"`
	Data       any    `json:"data"`
}

type CallSdkResp struct {
	ErrCode    int    `json:"err_code"`
	ErrMessage string `json:"err_message"`
	Result     any    `json:"result"`
}
