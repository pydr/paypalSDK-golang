package paypal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pydr/tools-golang"
)

// 创建订单
func (c *Client) CreateOrder(intent, currencyCode, value string) (*CreatedOrderData, error) {
	var ret CreatedOrderData

	url := c.APIBase + "/v2/checkout/orders"
	valueFloat, _ := strconv.ParseFloat(value, 64)
	value = fmt.Sprintf("%.2f", valueFloat/c.Curreny)

	amount := &Amount{
		CurrencyCode: currencyCode,
		Value:        value,
	}

	amountInfo := &AmountInfo{
		Amount: amount,
	}

	appInfo := &ApplicationContext{
		BrandName:  c.Brand,
		UserAction: "PAY_NOW",
		ReturnUrl:  c.ReturnUrl,
		CancelUrl:  c.CancelUrl,
	}

	payee := &Payee{
		Email: c.Account,
	}

	var amountList []*AmountInfo
	amountList = append(amountList, amountInfo)

	params := CreateOrderReqData{
		Intent:             intent,
		PurchaseUnits:      amountList,
		ApplicationContext: appInfo,
		Payee:              payee,
	}
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.Token.TokenType+" "+c.Token.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resq, result := tools.Request(c.HttpClient, req, 3)
	if !result {
		return nil, errors.New("request failed")
	}

	defer resq.Body.Close()
	body, err := ioutil.ReadAll(resq.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// 确认支付
func (c *Client) ConfirmOrder(orderId string) (*OrderCaptureData, error) {
	var ret OrderCaptureData

	url := c.APIBase + "/v2/checkout/orders/" + orderId + "/capture"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.Token.TokenType+" "+c.Token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resq, result := tools.Request(c.HttpClient, req, 3)
	if !result {
		return nil, errors.New("request failed")
	}

	defer resq.Body.Close()
	body, err := ioutil.ReadAll(resq.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
