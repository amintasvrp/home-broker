package entity

import (
	"github.com/amintasvrp/prosperity/trader/internal/market/entity/enums"
)

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     string
	Status        string
	Transactions  []*Transaction
}

func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	return &Order{
		ID:            orderID,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        enums.OPEN,
		Transactions:  []*Transaction{},
	}
}

func (o *Order) IncreaseAssetPosition(shares int) {
	o.Investor.IncreaseAssetPosition(o.Asset.ID, shares)
	o.PendingShares -= shares
}

func (o *Order) DecreaseAssetPosition(shares int) {
	o.Investor.DecreaseAssetPosition(o.Asset.ID, shares)
	o.PendingShares -= shares
}

func (o *Order) CloseOrder() {
	if o.PendingShares == 0 {
		o.Status = enums.CLOSED
	}
}
