package paypal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/pydr/tools-golang"
	"io/ioutil"
	"net/http"
	"strconv"
)

// 创建订单
func (c *Client) CreateOrder(intent, currencyCode, value, paypalAccount, brandName, returnUrl, cancelUrl string) *CreatedOrderData {
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
		BrandName:  brandName,
		UserAction: "PAY_NOW",
		ReturnUrl:  returnUrl,
		CancelUrl:  cancelUrl,
	}

	payee := &Payee{
		Email: paypalAccount,
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
		logs.Error("params error")
		return nil
	}

	logs.Warn(string(data))
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		logs.Error("make request failed: ", err)
		return nil
	}

	req.Header.Set("Authorization", c.Token.TokenType+" "+c.Token.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resq, result := tools.Request(c.HttpClient, req, 3)
	if !result {
		logs.Error("call api failed.")
		return nil
	}

	defer resq.Body.Close()

	body, err := ioutil.ReadAll(resq.Body)

	logs.Warn(string(body))
	err = json.Unmarshal(body, &ret)
	if err != nil {
		logs.Error("parse json data failed: ", err)
		return nil
	}

	return &ret
}

// 确认支付
func (c *Client) ConfirmOrder(orderId string) *OrderCaptureData {
	var ret OrderCaptureData

	url := c.APIBase + "/v2/checkout/orders/" + orderId + "/capture"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		logs.Error("make request failed: ", err)
		return nil
	}

	req.Header.Set("Authorization", c.Token.TokenType+" "+c.Token.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resq, result := tools.Request(c.HttpClient, req, 3)
	if !result {
		logs.Error("call api failed.")
		return nil
	}

	defer resq.Body.Close()

	body, err := ioutil.ReadAll(resq.Body)

	logs.Warn(string(body))
	err = json.Unmarshal(body, &ret)
	if err != nil {
		logs.Error("parse json data failed: ", err)
		return nil
	}

	return &ret
}
