package entry

import (
	"bytes"
	"encoding/json"
	"net/http"

	"e.coding.net/zklytech/serverlesss/user-model-sdk/mongo-sdk/domain"
)

const (
	applicationId = "Application-Id"
	modelId       = "Model-Id"
	instanceId    = "Instance-Id"
)

func CallSdkService(httpReq *http.Request, req domain.CallSdkReq) (resp domain.CallSdkResp, err error) {
	bData, err := json.Marshal(req)
	if err != nil {
		return
	}
	request, err := http.NewRequest("POST", "http://mongo-sdk-v1.default.local.com", bytes.NewReader(bData))
	if err != nil {
		return
	}

	request.Header.Add(applicationId, httpReq.Header.Get(applicationId))
	request.Header.Add(modelId, httpReq.Header.Get(modelId))
	request.Header.Add(instanceId, httpReq.Header.Get(instanceId))

	var client = http.Client{}

	httpResp, err := client.Do(request)
	if err != nil {
		return
	}

	var buff = make([]byte, httpResp.ContentLength)
	_, err = httpResp.Body.Read(buff)

	if err != nil {
		return
	}

	err = json.Unmarshal(buff, &resp)
	return
}
