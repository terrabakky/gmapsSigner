package gmapsSigner

import (
    "crypto/sha1"
    "crypto/hmac"
    "encoding/base64"
    "log"
    "net/url"
)

// SignUrl takes a secret key and a url to be signed and returns the request
// url with a signature parameter appended to the end
func SignUrl(secretKey string, fullUrl string) string {

    u, err := url.Parse(fullUrl)
    if err != nil {
        log.Fatal(err)
    }

    // We only sign the path and query portion of the URL
    urlToSign := u.Path + "?" + u.RawQuery

    // Decode the secret key using the bas64 URL encoding standard
    key, _ := base64.URLEncoding.DecodeString(secretKey)

    // Generate the SHA1-HMAC for the request
    h := hmac.New(sha1.New, key)
    h.Write([]byte(urlToSign))

    // Build up the request URL and append the generated request signature
    signedUrl := fullUrl + "&signature=" + base64.URLEncoding.EncodeToString(h.Sum(nil))
    
    return signedUrl
}
