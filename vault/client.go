package vault

import (
	"crypto/tls"
	"hafslundnett/hn-config-lib-go/certificates"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

//MakeClient explanation
func MakeClient(cfg Config) (*http.Client, error) {
	pool, err := certificates.MakePool(cfg.CaCert)
	if err != nil {
		log.Println("while getting CA Certs")
	}

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    pool,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	if err := http2.ConfigureTransport(transport); err != nil {
		log.Println("while configuring http2: ", err)
	}

	client := &http.Client{
		Transport: transport,
	}

	return client, nil
}
