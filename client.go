package paypal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pydr/tools-golang"
)

func NewPaypalClient(clientId, secret, apiBase, account, brand, returnUrl, cancelUrl string) *Client {

	client := &Client{
		ClientId:   clientId,
		Secret:     secret,
		APIBase:    apiBase,
		HttpClient: &http.Client{},
		Account:    account,
		Brand:      brand,
		ReturnUrl:  returnUrl,
		CancelUrl:  cancelUrl,
	}

	return client
}

func (c *Client) GetAccessToken() error {

	var ret TokenInfo

	url := c.APIBase + "/v1/oauth2/token"
	data := bytes.NewBuffer([]byte("grant_type=client_credentials"))
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.ClientId, c.Secret)
	resp, result := tools.Request(c.HttpClient, req, 3)
	if !result {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return err
	}

	c.Token = &ret
	c.TokenExpires = time.Now().Add(time.Duration(ret.ExpiresIn) * time.Second)

	return nil
}

func (c *Client) UpdateAccessToken() {
	for {
		timeout := c.TokenExpires.Sub(time.Now()) - time.Duration(10)
		<-time.After(timeout)
		c.GetAccessToken()
	}
}
