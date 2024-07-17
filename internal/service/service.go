package service

import (
	"Testovoe/internal/converter"
	"Testovoe/internal/domen"
	"Testovoe/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

const urlAPI = "https://garantex.org/api/v2/depth"

type RateServicer interface {
	Get(ctx context.Context, market string) (*domen.Response, error)
}

type RateService struct {
	repo   *repository.Repo
	logger *zap.Logger
}

func NewService(repos *repository.Repo, logger *zap.Logger) *RateService {
	return &RateService{repo: repos, logger: logger}
}

func (gr *RateService) Get(ctx context.Context, market string) (*domen.ResponseDTO, error) {
	resp, err := http.Get(fmt.Sprintf("%s?market=%s", urlAPI, market))
	if err != nil {
		gr.logger.Error("Failed get response", zap.Error(err))
		return nil, err
	}
	Rp := &domen.Response{}
	err = json.NewDecoder(resp.Body).Decode(Rp)
	if err != nil {
		gr.logger.Error("Failed decode", zap.Error(err))
		return nil, err
	}
	responseDTO := converter.FromResponseToResponseDTO(Rp)
	err = gr.repo.Create(ctx, responseDTO)
	if err != nil {
		gr.logger.Error("Failed saved DB", zap.Error(err))
		return nil, err
	}
	gr.logger.Info("Get rate and save to DB")
	return responseDTO, nil
}
