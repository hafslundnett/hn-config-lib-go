package vault

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

type loginToken struct {
	Token string `json:"token"`
}

type authToken struct {
	Token string `json:"auth"`
}

func main2() {
	configFile := "C:/secret/config.yaml"

	cfg, err := getConfig(configFile)
	if err != nil {
		log.Println("while getting configuration: ", err)
	}

	address := cfg.VaultAddr
	loginPath := "auth/github/login"
	version := "/v1/"

	rootCAs, err := RootCAs(cfg.CaCert)
	if err != nil {
		log.Println("while getting CA Cert")
	}

	tlsClientConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    rootCAs,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsClientConfig,
	}

	if err := http2.ConfigureTransport(transport); err != nil {
		log.Println("while configuring http2: ", err)
	}

	client := &http.Client{
		Transport: transport,
	}

	jsonStr, err := json.Marshal(cfg.LoginToken)
	if err != nil {
		log.Println("while marshaling token: ", err)
	}

	body := bytes.NewBuffer(jsonStr)
	url := address + version + loginPath

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Println("while building http request: ", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("while do-ing http request: ", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("while reading response body: ", err)
	}

	var auth interface{}

	err = json.Unmarshal(respBody, &auth)
	if err != nil {
		log.Println("while unmarhaling token: ", err)
	}

	m := auth.(map[string]interface{})
	n := m["auth"].(map[string]interface{})

	fmt.Println(n["client_token"])
}
