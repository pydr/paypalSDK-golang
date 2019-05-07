package paypal

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"turing/utils/tools"
)

// 验证webhook 有效性
func (c *Client) VerifyWebhook(transmissionId, transmissionTime, certUrl, authAlgo, transmissionSig, webhookId string, webhookEvent *WebhookEvent) bool {
	url := c.APIBase + "/v1/notifications/verify-webhook-signature"

	verifyData := WebhookVerifyData{
		TransmissionId:   transmissionId,
		TransmissionTime: transmissionTime,
		CertUrl:          certUrl,
		AuthAlgo:         authAlgo,
		TransmissionSig:  transmissionSig,
		WebhookId:        webhookId,
		WebhookEvent:     webhookEvent,
	}

	data, err := json.Marshal(verifyData)
	if err != nil {
		logs.Error("params error")
		return false
	}

	logs.Warn(string(data))
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		logs.Error("make request failed: ", err)
		return false
	}

	req.Header.Set("Authorization", c.Token.TokenType+" "+c.Token.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resq, result := tools.Request(c.HttpClient, req, 3)
	if !result {
		logs.Error("call api failed.")
		return false
	}

	defer resq.Body.Close()

	body, err := ioutil.ReadAll(resq.Body)
	logs.Warn(string(body))

	type retData struct {
		VerificationStatus string `json:"verification_status"`
	}

	var ret retData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		logs.Error("parse json data failed: ", err)
		return false
	}

	if ret.VerificationStatus != "SUCCESS" {
		return false
	}

	return true

}
