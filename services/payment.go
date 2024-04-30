package services

import "github.com/viadroid/xpanel-go/services/gateway"

type Payment struct{}

func NewPayment() *Payment {
	return &Payment{}
}
func (p *Payment) GetAllPaymentMap() []gateway.Gateway {
	return gateway.Gateways
}

func (p *Payment) GetPaymentsEnabled() []gateway.Gateway {
	var enabledList []gateway.Gateway

	for _, v := range p.GetAllPaymentMap() {
		if v.Enabled() {
			enabledList = append(enabledList, v)
		}
	}
	return enabledList
}

func (p *Payment) GetPaymentMap() map[string]gateway.Gateway {
	m := make(map[string]gateway.Gateway)
	for _, v := range p.GetPaymentsEnabled() {
		m[v.Name()] = v
	}
	return m
}

func (p *Payment) GetPaymentByName(name string) gateway.Gateway {
	return p.GetPaymentMap()[name]
}
