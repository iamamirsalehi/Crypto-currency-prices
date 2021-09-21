package main

import (
	"crypto/tls"
	"net/http"
)

func GetClient() *http.Client {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: config,
	}

	netClient := &http.Client{
		Transport: transport,
	}

	return netClient
}