package main

import (
	"Testovoe/internal/domen"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var market = "BTC/USDT"
	resp, err := http.Get(fmt.Sprintf("https://garantex.org/api/v2/depth?market=%s", market))
	if err != nil {
		fmt.Println(err, "bad zapros")
	}

	fmt.Println(resp.Status)
	var RG domen.Response
	err = json.NewDecoder(resp.Body).Decode(&RG)
	if err != nil {
		fmt.Println(err, "jsonerrdec")
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Println("err closed body")
		}

	}()
	fmt.Println(RG)

}
