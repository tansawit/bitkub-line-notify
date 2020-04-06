package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

type walletCoins struct {
	Btc  float64 `json:"BTC"`
	Eth  float64 `json:"ETH"`
	Wan  float64 `json:"WAN"`
	Ada  float64 `json:"ADA"`
	Omg  float64 `json:"OMG"`
	Bch  float64 `json:"BCH"`
	Usdt float64 `json:"USDT"`
	Ltc  float64 `json:"LTC"`
	Xrp  float64 `json:"XRP"`
	Bsv  float64 `json:"BSV"`
	Zil  float64 `json:"ZIL"`
	Snt  float64 `json:"SNT"`
	Cvc  float64 `json:"CVC"`
	Link float64 `json:"LINK"`
	Gnt  float64 `json:"GNT"`
	Iost float64 `json:"IOST"`
	Zrx  float64 `json:"ZRX"`
	Knc  float64 `json:"KNC"`
	Eng  float64 `json:"ENG"`
	Rdn  float64 `json:"RDN"`
	Abt  float64 `json:"ABT"`
	Mana float64 `json:"MANA"`
	Inf  float64 `json:"INF"`
	Ctcx float64 `json:"CTCX"`
	Xlm  float64 `json:"XLM"`
	Six  float64 `json:"SIX"`
	Jfin float64 `json:"JFIN"`
	Evx  float64 `json:"EVX"`
	Bnb  float64 `json:"BNB"`
	Pow  float64 `json:"POW"`
	Doge float64 `json:"DOGE"`
	Thb  float64 `json:"THB"`
	Dai  float64 `json:"DAI"`
}

type tickerRequest struct {
	coinStat coinStat
}

type coinStat struct {
	ID            int     `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	Last          float64 `json:"last,omitempty"`
	LowestAsk     float64 `json:"lowest_ask,omitempty"`
	HighestBid    float64 `json:"highest_bid,omitempty"`
	PercentChange float64 `json:"percent_change,omitempty"`
	BaseVolume    float64 `json:"base_volume,omitempty"`
	QuoteVolume   float64 `json:"quote_volume,omitempty"`
	IsFrozen      int     `json:"is_frozen,omitempty"`
	Figh24Hr      float64 `json:"high_24_hr,omitempty"`
	Low24Hr       float64 `json:"low_24_hr,omitempty"`
	Change        float64 `json:"change,omitempty"`
	PrevClose     float64 `json:"prev_close,omitempty"`
	PrevOpen      float64 `json:"prev_open,omitempty"`
}

type wallet struct {
	Err    int         `json:"error"`
	Result walletCoins `json:"result"`
}

type port struct {
	coins []coinStat
}

func getWallet(apiKey, apiSecret string) map[string]interface{} {
	// Balance (POST https://api.bitkub.com/api/market/wallet)
	now := time.Now()
	time := now.Unix()

	payload := fmt.Sprintf(`{"ts":%d}`, time)
	signedPayload := getSig(apiSecret, payload)
	jsonBody := []byte(signedPayload)
	body := bytes.NewBuffer(jsonBody)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://api.bitkub.com/api/market/wallet", body)

	// Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-BTK-APIKEY", apiKey)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}
	defer resp.Body.Close()

	var wallet wallet
	json.NewDecoder(resp.Body).Decode(&wallet)

	v := reflect.ValueOf(wallet.Result)
	port := make(map[string]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() != 0.0 {
			port[v.Type().Field(i).Name] = v.Field(i).Interface()
		}
	}
	return port
}

func getCurrentPrice(apiKey, ticker string) coinStat {
	// Market Ticker (GET https://api.bitkub.com/api/market/ticker?sym=THB_BTC)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bitkub.com/api/market/ticker?sym=THB_%s", ticker), nil)

	// Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-BTK-APIKEY", apiKey)
	req.Header.Add("Cookie", "__cfduid=d382fb39d4eae82331a4267439298f7341585974284; XSRF-TOKEN=eyJpdiI6IjdSXC9MNG9aa1p2Qk1WZEJjV0FnNkFBPT0iLCJ2YWx1ZSI6Ik9CbzlzNWt2Mkc1WGtNZVpNaXMwWjIwVkVLckxVdTdVN1dsRW5JTERqakJNaXNOVDNPSWRTM0p1ZkpKWGpkZVNXc0R2XC9WaUNpZFhsWlpFNXNQN2hRZz09IiwibWFjIjoiOGM3N2EzYjRkNDAxZmE2ZGI0ZDZlMGFlMTRlZDI4N2FiMmFiZDJhOThiMzAyMmVhNzRkOWZhY2E5ZjhiNjI5MyJ9; laravel_session=eyJpdiI6Im1qbThwMXM5STlUS0p5VjhYSTA4M0E9PSIsInZhbHVlIjoiUEJyZVBuYnRQUGlHd0RxMzhjbisrQ25OOTlncENvd292bVFKY2xJQm1SYm9aaDNxTTZMRG1mWlwvOVJtMnpsaXJZVkVrK1M3ekVrQ1BTSHVQZjhwcXdBPT0iLCJtYWMiOiI2Y2E4MGVjNjlmYTUxMDI1Y2MzYTBkMTQyNGU3NzM2NmEzYTNiZDRjN2JhZTRmMWFiODYwMmRlODE3N2Q4Yzk4In0%3D")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)

	var coinStat coinStat
	mapstructure.Decode(data[fmt.Sprintf("THB_%s", ticker)], &coinStat)
	return coinStat
}
