package gmapsSigner

import (
    "crypto/sha1"
    "crypto/hmac"
    "encoding/base64"
    "log"
    "net/url"
)

func SignUrl(secretKey string, fullUrl string) string {

    u, err := url.Parse(fullUrl)
    if err != nil {
        log.Fatal(err)
    }

    urlToSign := u.Path + "?" + u.RawQuery

    key, _ := base64.URLEncoding.DecodeString(secretKey)
    h := hmac.New(sha1.New, key)
    h.Write([]byte(urlToSign))


    signedUrl := fullUrl + "&signature=" + base64.URLEncoding.EncodeToString(h.Sum(nil))

    return signedUrl
}
