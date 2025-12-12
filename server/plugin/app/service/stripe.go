package service

import (
	"errors"
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/client"
)

type StripeSvc struct{ c *client.API }

var Stripe = new(StripeSvc)

func (s *StripeSvc) getClient() (*client.API, error) {
	if s.c != nil {
		return s.c, nil
	}
	key := os.Getenv("STRIPE_SECRET_KEY")
	if key == "" {
		return nil, errors.New("STRIPE_SECRET_KEY 未配置")
	}
	c := &client.API{}
	c.Init(key, nil)
	s.c = c
	return s.c, nil
}

// EnsureCustomer 在Stripe创建客户并返回ID
func (s *StripeSvc) EnsureCustomer(email string, userID uint) (string, error) {
	c, err := s.getClient()
	if err != nil {
		return "", err
	}
	params := &stripe.CustomerParams{Email: stripe.String(email)}
	params.AddMetadata("user_id", fmt.Sprintf("%d", userID))
	cust, err := c.Customers.New(params)
	if err != nil {
		return "", err
	}
	return cust.ID, nil
}

// CreatePaymentIntent 创建支付意图，返回 id 与 client_secret
func (s *StripeSvc) CreatePaymentIntent(amount int64, currency string, method string, customerID string, metadata map[string]string) (id string, clientSecret string, err error) {
	c, err := s.getClient()
	if err != nil {
		return "", "", err
	}
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
	}
	if customerID != "" {
		params.Customer = stripe.String(customerID)
	}
	switch method {
	case "alipay":
		params.PaymentMethodTypes = stripe.StringSlice([]string{"alipay"})
	case "wechat":
		params.PaymentMethodTypes = stripe.StringSlice([]string{"wechat_pay"})
		params.PaymentMethodOptions = &stripe.PaymentIntentPaymentMethodOptionsParams{
			WeChatPay: &stripe.PaymentIntentPaymentMethodOptionsWeChatPayParams{Client: stripe.String("android")},
		}
	default:
		params.PaymentMethodTypes = stripe.StringSlice([]string{"card"})
	}
	for k, v := range metadata {
		params.AddMetadata(k, v)
	}
	pi, err := c.PaymentIntents.New(params)
	if err != nil {
		return "", "", err
	}
	return pi.ID, pi.ClientSecret, nil
}
