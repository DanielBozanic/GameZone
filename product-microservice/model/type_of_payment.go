package model

type TypeOfPayment int

const (
	CASH_ON_DELIVERY Type = iota + 1
	PAYMENT_SLIP
)

func (t TypeOfPayment) String() string {
	return [...]string{
		"Cash on delivery",
		"Payment slip"}[t-1]
}