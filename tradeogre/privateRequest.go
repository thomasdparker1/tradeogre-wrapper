package tradeogre

import (
	"bytes"
	b64 "encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (to *Keys) PrivateRequest(u string, params map[string]string, query string) interface{} {
	data := url.Values{}
	if params != nil {
		x := 0
		for k, v := range params {
			if x == 0 {
				data.Set(k, v)
			}
			if x != 0 {
				data.Add(k, v)
			}
			x++
		}
	}

	req, err := http.NewRequest("POST", u+query, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Println(err.Error())
	}
	if params == nil {
		req, err = http.NewRequest("GET", u+query, nil)
	}
	key := "Basic " + b64.StdEncoding.EncodeToString([]byte(to.PublicKey+":"+to.PrivateKey))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("Authorization", key)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	return response
}
