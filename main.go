package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation/v2"
	gh "github.com/google/go-github/v32/github"
)

var APP_ID = int64(123456)
var ENCODED_CERT = ""

func main() {
	fmt.Println("main")
	VerifyAppInstallation()
}

func VerifyAppInstallation() {
	appID := APP_ID
	encodedCert := ENCODED_CERT
	certificate, err := base64.StdEncoding.DecodeString(encodedCert)
	if err != nil {
		return
	}
	tr := http.DefaultTransport
	itr, err := ghinstallation.NewAppsTransport(tr, appID, certificate)
	if err != nil {
		return
	}
	// Use installation transport with github.com/google/go-github
	client := gh.NewClient(&http.Client{Transport: itr})

	ins, _, err := client.Apps.ListInstallations(context.Background(), &gh.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(ins))

	for _, installation := range ins {
		fmt.Println(*installation.Account.HTMLURL)
	}
}
