package paypal

import (
	"net/http"
	"time"
)

const (
	APIBaseSandBox = "https://api.sandbox.paypal.com"
	APIBaseLive    = "https://api.paypal.com"
)

type (
	Client struct {
		ClientId     string
		Secret       string
		APIBase      string
		Token        *TokenInfo
		HttpClient   *http.Client
		TokenExpires time.Time
		Curreny      float64
		Account      string // 收款方paypal帐号
		Brand        string // 收款方名称
		ReturnUrl    string // 支付页面url
		CancelUrl    string // 取消支付按钮url
	}

	TokenInfo struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		AppId       string `json:"app_id"`
		ExpiresIn   int    `json:"expires_in"`
	}

	Link struct {
		Href    string `json:"href"`
		Rel     string `json:"rel"`
		Method  string `json:"method"`
		EncType string `json:"encType,omitempty"`
	}

	CreatedOrderData struct {
		Id     string  `json:"id"`
		Status string  `json:"status"`
		Links  []*Link `json:"links"`
	}

	Details struct {
		Subtotal string `json:"subtotal"`
	}

	Amount struct {
		Total        string   `json:"total,omitempty"`
		Details      *Details `json:"details,omitempty"`
		CurrencyCode string   `json:"currency_code,omitempty"`
		Value        string   `json:"value,omitempty"`
	}

	AmountInfo struct {
		Amount *Amount `json:"amount"`
	}

	Payee struct {
		Email string `json:"email"`
	}

	ApplicationContext struct {
		BrandName  string `json:"brand_name"`
		UserAction string `json:"user_action"`
		ReturnUrl  string `json:"return_url"`
		CancelUrl  string `json:"cancel_url"`
	}

	CreateOrderReqData struct {
		Intent             string              `json:"intent"`
		PurchaseUnits      []*AmountInfo       `json:"purchase_units"`
		ApplicationContext *ApplicationContext `json:"application_context"`
		Payee              *Payee              `json:"payee"`
	}

	PayerName struct {
		GivenName string `json:"given_name"`
		Surname   string `json:"surname"`
	}

	Payer struct {
		Name         *PayerName `json:"name"`
		EmailAddress string     `json:"email_address"`
		PayerId      string     `json:"payer_id"`
	}

	Address struct {
		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		AdminArea1   string `json:"admin_area_1"`
		AdminArea2   string `json:"admin_area_2"`
		PostalCode   string `json:"postal_code"`
		CountryCode  string `json:"country_code"`
	}

	Shipping struct {
		Address *Address `json:"address"`
	}

	SellerReceivableBreakdown struct {
		GrossAmount *Amount `json:"gross_amount"`
		PaypalFee   *Amount `json:"paypal_fee"`
		NetAmount   *Amount `json:"net_amount"`
	}

	Capture struct {
		Id                        string                     `json:"id"`
		Status                    string                     `json:"status"`
		Amount                    *Amount                    `json:"amount"`
		SellerReceivableBreakdown *SellerReceivableBreakdown `json:"seller_receivable_breakdown"`
		CreateTime                time.Time                  `json:"create_time"`
		UpdateTime                time.Time                  `json:"update_time"`
	}

	Payments struct {
		Captures []*Capture
	}

	PurchaseUnits struct {
		ReferenceId string    `json:"reference_id"`
		Shipping    *Shipping `json:"shipping"`
		Payments    *Payments `json:"payments"`
	}

	OrderCaptureData struct {
		Id            string           `json:"id"`
		Status        string           `json:"status"`
		Payer         *Payer           `json:"payer"`
		PurchaseUnits []*PurchaseUnits `json:"purchase_units"`
	}

	SellerProtection struct {
		Status            string   `json:"status"`
		DisputeCategories []string `json:"dispute_categories"`
	}

	Resource struct {
		Id         string `json:"id"`
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
		//State         string    `json:"state"`
		Amount                    *Amount                    `json:"amount"`
		FinalCapture              bool                       `json:"final_capture"`
		SellerProtection          *SellerProtection          `json:"seller_protection"`
		SellerReceivableBreakdown *SellerReceivableBreakdown `json:"seller_receivable_breakdown"`
		Status                    string                     `json:"status"`
		//ParentPayment string    `json:"parent_payment"`
		//ValidUntil    time.Time `json:"valid_until"`
		Links []*Link `json:"links"`
	}

	WebhookEvent struct {
		Id              string    `json:"id"`
		CreateTime      string    `json:"create_time"`
		ResourceType    string    `json:"resource_type"`
		EventVersion    string    `json:"event_version"`
		EventType       string    `json:"event_type"`
		Summary         string    `json:"summary"`
		ResourceVersion string    `json:"resource_version"`
		Resource        *Resource `json:"resource"`
		Links           []*Link   `json:"links"`
	}

	WebhookVerifyData struct {
		TransmissionId   string        `json:"transmission_id"`
		TransmissionTime string        `json:"transmission_time"`
		CertUrl          string        `json:"cert_url"`
		AuthAlgo         string        `json:"auth_algo"`
		TransmissionSig  string        `json:"transmission_sig"`
		WebhookId        string        `json:"webhook_id"`
		WebhookEvent     *WebhookEvent `json:"webhook_event"`
	}
)
