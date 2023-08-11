package entry

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	modeldbDomain "github.com/ps756405678/modeldb-sdk/domain"
	"github.com/ps756405678/user-model-sdk/consts"
	"github.com/ps756405678/user-model-sdk/domain"
)

const (
	applicationIdHK = "Application-Id"
	modelIdHK       = "Model-Id"
	instanceIdHK    = "Instance-Id"

	callModelFuncApi = "http://192.168.0.68:8082/api/model/func/url"
)

func Instantiate(httpReq *http.Request, req modeldbDomain.ModelDBDescribe) (resp modeldbDomain.ModelDBDescribe, err error) {
	respData, err := callModelFunc(httpReq, req, consts.Instantiate)
	if err != nil {
		return
	}
	err = json.Unmarshal(respData, &resp)
	return
}

func Login(httpReq *http.Request, req domain.User) (resp domain.User, err error) {
	respData, err := callModelFunc(httpReq, req, consts.Login)
	if err != nil {
		return
	}
	err = json.Unmarshal(respData, &resp)
	return
}

func Register(httpReq *http.Request, req domain.User) (resp domain.User, err error) {
	respData, err := callModelFunc(httpReq, req, consts.Register)
	if err != nil {
		return
	}
	err = json.Unmarshal(respData, &resp)
	return
}

func callModelFunc(httpReq *http.Request, req any, methodName string) (resp []byte, err error) {
	url, err := getModelFuncUrl(httpReq.Header.Get(modelIdHK), methodName)
	if err != nil {
		return
	}
	bData, err := json.Marshal(req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bData))
	if err != nil {
		return
	}

	var client = http.Client{}

	httpResp, err := client.Do(request)
	if err != nil {
		return
	}

	resp = make([]byte, httpResp.ContentLength)
	_, err = httpResp.Body.Read(resp)

	return
}

func getModelFuncUrl(modelId string, methodName string) (url string, err error) {
	var callFuncReq = map[string]any{
		"model_db_id": modelId,
		"func_name":   methodName,
	}
	bData, err := json.Marshal(callFuncReq)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, callModelFuncApi, bytes.NewReader(bData))
	if err != nil {
		return
	}

	var client = http.Client{}

	httpResp, err := client.Do(request)
	if err != nil {
		return
	}

	buff := make([]byte, httpResp.ContentLength)
	httpResp.Body.Read(buff)

	var result map[string]any
	err = json.Unmarshal(buff, &result)
	if err != nil {
		return
	}

	if result["errcode"].(int) != 0 {
		err = errors.New(result["msg"].(string))
		return
	}

	url = "http://" + result["data"].(string)
	return
}
