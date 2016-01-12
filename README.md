# gmapsSigner
Go package for signing Google Maps for Work requests use the client secret key

Example usage:

**Fetch the package**

`$ go get github.com/terrabakky/gmapsSigner `

**test.go**
```
package main

import "fmt"
import "github.com/terrabakky/gmapsSigner"

func main() {
    // The request URL you want to sign
	url := "http://maps.googleapis.com/maps/api/geocode/json?client=gme-test123"
	// Your client signature - this one is borrowed from the Google Maps Docs
	secret := "chaRF2hTJKOScPr-RQCEhZbSzIE="
	signedUrl := gmapsSigner.SignUrl(secret, url)
	fmt.Println(signedUrl)
}
```
**Run**

`go run test.go`

The output for this test should be:

`http://maps.googleapis.com/maps/api/geocode/json?client=gme-test123&signature=vBayVIo1sb7_5LJ-uEddsadsL0g=`

You can check the output against the [Maps for Work Signing Debugger](https://m4b-url-signer.appspot.com/) provided by Google.



