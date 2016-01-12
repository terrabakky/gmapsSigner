package main

import (
	"crypto/sha1"
	"crypto/hmac"
    "encoding/base64"
	"fmt"
	"log"
	"net/url"
)


func SignUrl(secretKey string, fullUrl string) string {

	u, err := url.Parse(fullUrl)
	if err != nil {
		log.Fatal(err)
	}

	urlToSign := u.Path + "?" + u.RawQuery

	fmt.Println("URL TO SIGN: " + urlToSign)

	key, _ := base64.URLEncoding.DecodeString(secretKey)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(urlToSign))


	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	// Sig should be: vBayVIo1sb7_5LJ-uEddsadsL0g=

	secretKey := "chaRF2hTJKOScPr-RQCEhZbSzIE="
	fullUrl := "http://maps.googleapis.com/maps/api/geocode/json?client=gme-test123"

	fmt.Println("Signature: " + SignUrl(secretKey, fullUrl))
}