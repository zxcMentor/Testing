package service

import (
	"Testovoe/internal/domen"
	"Testovoe/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const urlAPI = "https://garantex.org/api/v2/depth"

type GetRates struct {
	Repos repository.Repository
}

func (gr *GetRates) Get(ctx context.Context, market string) (*domen.Response, error) {
	resp, err := http.Get(fmt.Sprintf("%s?market=%s", urlAPI, market))
	if err != nil {
		fmt.Println(err, "bad zapros")
	}
	var Rp domen.Response
	err = json.NewDecoder(resp.Body).Decode(&Rp)
	if err != nil {
		fmt.Println("err decode ")
	}
	err = gr.Repos.Create(&Rp)
	if err != nil {
		fmt.Println("err create ")
	}
	return &Rp, nil
}
