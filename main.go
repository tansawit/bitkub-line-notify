package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gocron.Every(1).Day().At("9:00").Do(sendNotification)
	<-gocron.Start()
}

func sendNotification() {
	apiKey := os.Getenv("BITKUB_API_KEY")
	apiSecret := os.Getenv("BITKUB_API_SECRET")

	// Get port info
	port := getWallet(apiKey, apiSecret)
	var portValue float64
	for ticker, amount := range port {
		coinStat := getCurrentPrice(apiKey, strings.ToUpper(ticker))
		amountFloat, err := getFloat(amount)
		if err != nil {
			panic("")
		}
		portValue += amountFloat * coinStat.Last
	}

	// Create the message to be sent
	message := fmt.Sprintf("\nTotal portfolio value: %.2f THB", portValue)
	for ticker, amount := range port {
		coinStat := getCurrentPrice(apiKey, strings.ToUpper(ticker))
		message += fmt.Sprintf("\n%s: %.2f coins (%.2f THB, %.2f%%)", ticker, amount, coinStat.Last, coinStat.PercentChange)
	}

	//Send the message
	apiURL := "https://notify-api.line.me/api/notify"
	data := url.Values{}
	data.Set("message", message)
	u, _ := url.ParseRequestURI(apiURL)
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", "Bearer cyECqwAUl8vOFmeEMrEIQFIjUHtl0iz8uKc7odC6cIA")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
}
