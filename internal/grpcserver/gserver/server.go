package gserver

import (
	"Testovoe/internal/converter"
	"Testovoe/internal/service"
	proto "Testovoe/protos/gen"
	"context"
	"go.uber.org/zap"
)

type Rate struct {
	proto.UnimplementedGetRatesServer
	serv   *service.RateService
	logger *zap.Logger
}

func NewServer(serv *service.RateService, logger *zap.Logger) *Rate {
	return &Rate{serv: serv, logger: logger}
}

func (r *Rate) Get(ctx context.Context, req *proto.Req) (*proto.Response, error) {
	resp, err := r.serv.Get(ctx, req.Market)
	if err != nil {
		r.logger.Error("Failed get rate: ", zap.Error(err))
		return nil, err
	}

	askOrders, bidOrders := converter.FromResponseToProto(resp)
	return &proto.Response{
		Timestamp: resp.Timestamp,
		Ask:       askOrders,
		Bid:       bidOrders,
	}, nil
}
