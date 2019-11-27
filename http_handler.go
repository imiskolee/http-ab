package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	jd "github.com/josephburnett/jd/lib"
)

func DoHttpRequest(requestConfig *RequestConfig) (*http.Response,error) {
	buf := bytes.NewBufferString(requestConfig.Payload)
	req,err := http.NewRequest(requestConfig.Method,requestConfig.URL + requestConfig.URI,buf)
	if err != nil {
		return nil,err
	}
	for k,v := range requestConfig.Headers {
		req.Header.Set(k,v)
	}
	resp,err := http.DefaultClient.Do(req)
	return resp,err
}

func CompareResponse(config *Config,mainlyResp *http.Response,secondlyResp *http.Response) {
	if mainlyResp.StatusCode != secondlyResp.StatusCode {
		fmt.Printf("[CompareResponse] diff status code: %d != %d\n",mainlyResp.StatusCode,secondlyResp.StatusCode)
		os.Exit(1)
	}
	c1,err := ioutil.ReadAll(mainlyResp.Body)
	defer mainlyResp.Body.Close()
	if err != nil {
		fmt.Printf("[CompareResponse] Can't read mainly response:%s\n",err.Error())
		os.Exit(1)
	}

	c2,err := ioutil.ReadAll(secondlyResp.Body)
	defer mainlyResp.Body.Close()
	if err != nil {
		fmt.Printf("[CompareResponse] Can't read secondly response:%s\n",err.Error())
		os.Exit(1)
	}
	switch config.ResponseParser {
	case ResponseParserJSON:
		a, _ := jd.ReadJsonString(string(c1))
		b, _ := jd.ReadJsonString(string(c2))
		ret := a.Diff(b)
		diff := ret.Render()
		if diff == "" {
			fmt.Println("no any diffs found...")
			os.Exit(0)
		}
		fmt.Print(diff)
		os.Exit(1)
	case ResponseParserTEXT:
		panic("not yet supported ")
	}
}
