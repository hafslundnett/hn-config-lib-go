package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type loginToken struct {
	Token string `json:"token"`
}

type authToken struct {
	Token string `json:"auth"`
}

func main() {
	configFile := "C:/secret/config.yaml"
	cfg, err := getConfig(configFile)
	if err != nil {
		log.Println("while getting configuration: ", err)
	}

	gitToken := loginToken{cfg.GitToken}
	address := cfg.VaultAddr
	path := "auth/github/login"
	version := "/v1/"

	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	jsonStr, err := json.Marshal(gitToken)
	if err != nil {
		log.Println("while marshaling token: ", err)
	}

	body := bytes.NewBuffer(jsonStr)
	url := address + version + path

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
