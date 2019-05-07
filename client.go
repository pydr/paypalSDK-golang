package paypal

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"time"
	"turing/utils/tools"
)

func NewPaypalClient(clientId, secret, apiBase string) *Client {

	client := &Client{
		ClientId:   clientId,
		Secret:     secret,
		APIBase:    apiBase,
		HttpClient: &http.Client{},
	}

	return client
}

func (c *Client) GetAccessToken() {

	var ret TokenInfo

	url := c.APIBase + "/v1/oauth2/token"
	data := bytes.NewBuffer([]byte("grant_type=client_credentials"))

	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		logs.Error("make request failed: ", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.ClientId, c.Secret)

	resp, result := tools.Request(c.HttpClient, req, 3)
	if !result {
		logs.Error("call api failed: ", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		logs.Error("parse json data failed: ", err)
		return
	}

	c.Token = &ret
	c.TokenExpires = time.Now().Add(time.Duration(ret.ExpiresIn) * time.Second)

	return
}

func (c *Client) UpdateAccessToken() {
	for {
		timeout := c.TokenExpires.Sub(time.Now()) - time.Duration(10)
		<-time.After(timeout)
		logs.Info("更新paypal token")
		c.GetAccessToken()
	}
}
