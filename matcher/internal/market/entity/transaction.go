package entity

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
		DateTime:     time.Now(),
	}
}

func (t *Transaction) CalculateTotal() {
	t.Total = float64(t.Shares) * t.BuyingOrder.Price
}

func (t *Transaction) MoveAssetPosition(shares int) {
	t.SellingOrder.DecreaseAssetPosition(shares)
	t.BuyingOrder.IncreaseAssetPosition(shares)
}

func (t *Transaction) CloseOrders() {
	t.BuyingOrder.CloseOrder()
	t.SellingOrder.CloseOrder()
}
