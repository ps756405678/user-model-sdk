package entry

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ps756405678/user-model-sdk/consts"
	"github.com/ps756405678/user-model-sdk/domain"
)

const (
	applicationId = "Application-Id"
	modelId       = "Model-Id"
	instanceId    = "Instance-Id"
	method        = "method"
)

func Register(httpReq *http.Request, req domain.User) (resp domain.User, err error) {
	respData, err := callModelFunc(httpReq, req, consts.Register)
	if err != nil {
		return
	}
	err = json.Unmarshal(respData, &resp)
	return
}

func callModelFunc(httpReq *http.Request, req any, methodName string) (resp []byte, err error) {
	bData, err := json.Marshal(req)
	if err != nil {
		return
	}
	request, err := http.NewRequest("POST", "http://user-model-v1.default.scv.cluster.local", bytes.NewReader(bData))
	if err != nil {
		return
	}

	request.Header.Add(applicationId, httpReq.Header.Get(applicationId))
	request.Header.Add(modelId, httpReq.Header.Get(modelId))
	request.Header.Add(instanceId, httpReq.Header.Get(instanceId))
	request.Header.Add(method, methodName)

	var client = http.Client{}

	httpResp, err := client.Do(request)
	if err != nil {
		return
	}

	resp = make([]byte, httpResp.ContentLength)
	_, err = httpResp.Body.Read(resp)

	return
}
