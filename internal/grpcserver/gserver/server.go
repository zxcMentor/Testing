package gserver

import (
	"Testovoe/internal/converter"
	"Testovoe/internal/service"
	pbrate "Testovoe/protos/gen"
	"context"
	"log/slog"
)

type Rate struct {
	pbrate.UnimplementedGetRatesServer
	serv service.GetRates
}

func (r *Rate) Get(ctx context.Context, req *pbrate.Req) (*pbrate.Response, error) {
	get, err := r.serv.Get(ctx, req.Market)
	if err != nil {
		slog.Info("err in service Get")
		return nil, err
	}

	askOrders, bidOrders := converter.ConvertToProto(get)
	return &pbrate.Response{
		Timestamp: get.Timestamp,
		Ask:       askOrders,
		Bid:       bidOrders,
	}, nil
}
