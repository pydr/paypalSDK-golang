## Paypal SDK for golang

[![Build Status](https://travis-ci.org/pydr/paypalSDK-golang.svg?branch=master)](https://travis-ci.org/pydr/paypalSDK-golang)



### Descript

---

This sdk is base on paypal restful api.



### Available interface

---

| InterfaceName  | Descript                  | Api                                        |
| -------------- | ------------------------- | ------------------------------------------ |
| GetAccessToken | Get api access token      | /v1/oauth2/token                           |
| CreateOrder    | Creat a new payment order | /v2/checkout/orders                        |
| ConfirmOrder   | Confirm payment           | /v2/checkout/orders/<orderid>/capture      |
| VerifyWebhook  | Verify Webhook event data | /v1/notifications/verify-webhook-signature |



### Before you start

---

- <a href="https://www.paypal.com/us/webapps/mpp/account-selection">Get a Paypal Account</a>
- <a href="https://developer.paypal.com/developer/applications">Create a restful APP</a>



### Document

- <a href="https://developer.paypal.com/docs/api/overview/">English Version</a>

### Quick Start

---

#### installation

```bash
go get -u github.com/pydr/paypalSDK-golang
```



#### Import the SDK package

```go
import paypal "github.com/pydr/paypalSDK-golang"
```



#### Create a new sdk client

```go
// apiBase: enum("APIBaseSandBox" | "APIBaseLive") -> devlopement | live
// account: receiver paypal account(email address)
// brand:   receiver name, display on payment page
// returnUrl: payment page url
// cancelUrl: cancel payment
client := paypal.NewPaypalClient(clientId, secret, apiBase, account, brand, returnUrl, cancelUrl)
```



#### Create a new payment order

```go
// intent: 
// currencyCode: E.g "USD", "CNY" etc.
// value: payment value
// newOrderData include orderId, approve url, capture url etc.
newOrderData := client.CreateOrder(intent, currencyCode, value)
```



> when customer approved this order, you can capture this paymet.



### Capture payment

```
newCaptureData := client.ConfirmOrder(orderId)
```

> Ok, this payment maybe success. but you must to verify. Paypal use Webhook to notify server events. So the server should provide an interface(restful api) to recive paypal events.



#### Verify webhook data

```go
result := client.VerifyWebhook(transmissionId, transmissionTime, certUrl, authAlgo, transmissionSig, webhookId, webhookEvent)
```

>About parameters, you can visit <a href="https://developer.paypal.com/docs/api/webhooks/v1/#verify-webhook-signature_post">Doc</a>.

