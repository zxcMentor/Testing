package converter

import (
	"Testovoe/internal/domen"
	pbrate "Testovoe/protos/gen"
)

func FromResponseToProto(resp *domen.ResponseDTO) (*pbrate.Order, *pbrate.Order) {
	orderAsk := &pbrate.Order{
		Price:  resp.Asks.Price,
		Volume: resp.Asks.Volume,
		Amount: resp.Asks.Amount,
		Factor: resp.Asks.Factor,
		Type:   resp.Asks.Type,
	}

	orderBids := &pbrate.Order{
		Price:  resp.Bids.Price,
		Volume: resp.Bids.Volume,
		Amount: resp.Bids.Amount,
		Factor: resp.Bids.Factor,
		Type:   resp.Bids.Type,
	}
	return orderAsk, orderBids
}

func FromResponseToResponseDTO(resp *domen.Response) *domen.ResponseDTO {
	return &domen.ResponseDTO{
		Asks: domen.Order{
			Price:  resp.Asks[0].Price,
			Volume: resp.Asks[0].Volume,
			Amount: resp.Asks[0].Amount,
			Factor: resp.Asks[0].Factor,
			Type:   resp.Asks[0].Type,
		},
		Bids: domen.Order{
			Price:  resp.Bids[0].Price,
			Volume: resp.Bids[0].Volume,
			Amount: resp.Bids[0].Amount,
			Factor: resp.Bids[0].Factor,
			Type:   resp.Bids[0].Type,
		},
		Timestamp: resp.Timestamp,
	}
}
