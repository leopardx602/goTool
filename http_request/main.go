package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func Get() error {
	resp, err := http.Get("http://localhost:5000/json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func Post() error {
	postData := url.Values{"key1": {"value1"}, "key2": {"value2"}}
	body := strings.NewReader(postData.Encode())

	resp, err := http.Post("http://localhost:5000/json", "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func Request(method, uri string) (respBody []byte, err error) {
	// postData := map[string]string{"key1": "value1"}
	// dataByte, err := json.Marshal(postData)
	// if err != nil {
	// 	return nil, err
	// }

	req, err := http.NewRequest(method, uri, nil) // bytes.NewBuffer(dataByte)
	if err != nil {
		return nil, errors.Wrap(err, "error in NewRequest")
	}

	// req.Header.Add("Content-type", "application/json")

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConnsPerHost = 1000
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: t,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do a request")
	}
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error in reading body")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected status code:%v, body:%v", resp.StatusCode, string(respBody))
	}
	return respBody, nil
}

func main() {
	if err := Get(); err != nil {
		panic(err)
	}
}
