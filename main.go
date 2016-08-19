package main

import (
	//"fmt"
	//"github.com/franela/goreq"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type ApiResponse struct {
	Status  string  `json:"status"`
	Data    ApiData `json:"data"`
	Message string  `json:"message,omitempty"`
	Code    int     `json:"code,omitempty"`
}

type ApiData map[string]interface{}

type MarketRecord struct {
	volume          int     `json:"volume"`
	latest_trade    int     `json:"latest_trade"`
	bid             int     `json:"bid"`
	high            float64 `json:"high"`
	currency        string  `json:"currency"`
	currency_volume float64 `json:"currency_volume"`
	ask             float64 `json:"ask"`
	close           float64 `json:"close"`
	avg             float64 `json:"avg"`
	symbol          string  `json:"symbol"`
	low             float64 `json:"low"`
}

func main() {
	r := gin.Default()
	r.GET("/bitcoin", bitcoinGET)
	r.Run()
}

func bitcoinGET(c *gin.Context) {
	url := "http://api.bitcoincharts.com/v1/markets.json"

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data interface{} // Market Records
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	var rsp = ApiResponse{
		Status: "success",
		Data: ApiData{
			"records": data,
		},
	}

	c.JSON(http.StatusOK, rsp)
}
