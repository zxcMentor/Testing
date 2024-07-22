package gserver

import (
	"Testovoe/internal/converter"
	"Testovoe/internal/service"
	proto "Testovoe/protos/gen"
	"context"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Rate struct {
	proto.UnimplementedGetRatesServer
	serv   *service.RateService
	logger *zap.Logger
	tp     trace.Tracer
}

func NewServer(serv *service.RateService, logger *zap.Logger, tp trace.Tracer) *Rate {
	return &Rate{serv: serv, logger: logger, tp: tp}
}

func (r *Rate) Get(ctx context.Context, req *proto.Req) (*proto.Response, error) {
	ctx, span := r.tp.Start(ctx, "server grpc")
	defer span.End()
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
