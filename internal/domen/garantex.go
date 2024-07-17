package domen

type Response struct {
	Asks      []Order `json:"asks"`
	Bids      []Order `json:"bids"`
	Timestamp int64   `json:"timestamp"`
}

type Order struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}
type Request struct {
	Message string `json:"message"`
}

type ResponseDTO struct {
	Asks      Order `json:"asks"`
	Bids      Order `json:"bids"`
	Timestamp int64 `json:"timestamp"`
}
