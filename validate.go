package paypal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pydr/tools-golang"
)

// 验证webhook 有效性
func (c *Client) VerifyWebhook(transmissionId, transmissionTime, certUrl, authAlgo, transmissionSig, webhookId string, webhookEvent *WebhookEvent) (bool, error) {
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
		return false, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", c.Token.TokenType+" "+c.Token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resq, result := tools.Request(c.HttpClient, req, 3)
	if !result {
		return false, err
	}

	defer resq.Body.Close()
	body, err := ioutil.ReadAll(resq.Body)

	type retData struct {
		VerificationStatus string `json:"verification_status"`
	}

	var ret retData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return false, err
	}

	if ret.VerificationStatus != "SUCCESS" {
		return false, nil
	}

	return true, nil

}
