package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const infoFile = "info.json"

type Stock struct {
	Name          string
	PurchasePrice float64
	SellingPrice  float64
	Result        float64
	Period        struct {
		Begin time.Time
		End   time.Time
	}
}

func calcStockReturns(stock *Stock) float64 {
	x := (stock.SellingPrice - stock.PurchasePrice) / (stock.PurchasePrice) //рыночная доходность акций
	y := (countDays(stock.Period.Begin, stock.Period.End) / 390)
	stock.Result = x * y * 100
	return stock.Result
}

func countDays(t1 time.Time, t2 time.Time) float64 {
	return t2.Sub(t1).Hours() / 24 //count days between two dates
}

func toJson(stock *Stock) {
	byteArray, err := json.Marshal(stock)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile(infoFile, byteArray, 0)
	fmt.Println(string(byteArray))
}

func main() {
	stock := &Stock{
		Name:          "Sber",
		PurchasePrice: 65,
		SellingPrice:  98,
		Period: struct {
			Begin time.Time
			End   time.Time
		}{
			Begin: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	fmt.Println(calcStockReturns(stock))
	toJson(stock)
}
