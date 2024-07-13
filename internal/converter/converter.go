package converter

import (
	"Testovoe/internal/domen"
	pbrate "Testovoe/protos/gen"
)

func ConvertToProto(get *domen.Response) ([]*pbrate.Order, []*pbrate.Order) {
	askOrders := make([]*pbrate.Order, 0, len(get.Ask))
	for _, ask := range get.Ask {
		askOrder := &pbrate.Order{
			Price:  ask.Price,
			Volume: ask.Volume,
			Amount: ask.Amount,
			Factor: ask.Factor,
			Type:   ask.Type,
		}
		askOrders = append(askOrders, askOrder)
	}
	bidOrders := make([]*pbrate.Order, 0, len(get.Bids))
	for _, bid := range get.Bids {
		askOrder := &pbrate.Order{
			Price:  bid.Price,
			Volume: bid.Volume,
			Amount: bid.Amount,
			Factor: bid.Factor,
			Type:   bid.Type,
		}
		bidOrders = append(bidOrders, askOrder)
	}
	return askOrders, bidOrders
}
