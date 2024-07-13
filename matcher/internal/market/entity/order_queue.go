package entity

type OrderQueue struct {
	Orders []*Order
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}

func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

func (oq *OrderQueue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

func (oq *OrderQueue) Pop() interface{} {
	oldOrders := oq.Orders
	n := len(oldOrders)
	item := oldOrders[0]
	oq.Orders = oldOrders[1:n]
	return item
}
