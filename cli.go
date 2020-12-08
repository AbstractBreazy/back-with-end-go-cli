package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"mime/multipart"
	"net/http"
	"strconv"
)

const (
	addr         = "http://localhost:8090/"    // BASIC_URL
	process_addr = addr + "api/status/process" // [POST]
	check_addr   = addr + "api/status/check"   // [GET]
)

type Entity struct {
	ID     string `json:"id"`
	Value  string `json:"value"`
	Data   int64  `json:"data"`
	Result string `json:"result"`
}

func main() {
	requestType := flag.String("type", "", "Type of Request: prepare, start processing or clear, "+
		"default = undefined type request")
	value := flag.String("value", "", "Text Value")
	data := flag.String("data", "", "Int Value")

	flag.Parse()

	if len(*requestType) == 0 {
		panic(errors.New("please provide correct requestType"))
	}

	switch *requestType {
	case "prepare":
		{
			if len(*value) < 1 || len(*data) < 1 {
				panic(errors.New("please provide correct value, data"))
			}
			if IsNumeric(*data) == false {
				panic(errors.New("data isn't num"))
			}
			dataNum, _ := strconv.ParseFloat(*data, 64)
			if math.Signbit(dataNum) == true {
				panic(errors.New("data is negative"))
			}
			client := &http.Client{}
			regBody := &bytes.Buffer{}
			writer := multipart.NewWriter(regBody)
			_ = writer.WriteField("value", *value)
			_ = writer.WriteField("data", *data)
			_ = writer.WriteField("type", *requestType)
			err := writer.Close()
			if err != nil {
				panic(errors.New("failed to close writer"))
			}
			req, err := http.NewRequest("POST", process_addr, regBody)
			if err != nil {
				panic(errors.New("failed to create http request"))
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())
			res, err := client.Do(req)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)

			fmt.Println(string(body))

		}
	case "start processing":
		{
			if len(*value) < 1 || len(*data) < 1 {
				panic(errors.New("please provide correct value, data"))
			}
			if IsNumeric(*data) == false {
				panic(errors.New("data isn't num"))
			}
			dataNum, _ := strconv.ParseFloat(*data, 64)
			if math.Signbit(dataNum) == true {
				panic(errors.New("data is negative"))
			}
			client := &http.Client{}
			regBody := &bytes.Buffer{}
			writer := multipart.NewWriter(regBody)
			_ = writer.WriteField("value", *value)
			_ = writer.WriteField("data", *data)
			_ = writer.WriteField("type", *requestType)
			err := writer.Close()
			if err != nil {
				panic(errors.New("failed to close writer"))
			}
			req, err := http.NewRequest("POST", process_addr, regBody)
			if err != nil {
				panic(errors.New("failed to create http request"))
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())
			res, err := client.Do(req)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)

			fmt.Println(string(body))

		}
	case "clear":
		{
			client := &http.Client{}
			regBody := &bytes.Buffer{}
			writer := multipart.NewWriter(regBody)
			_ = writer.WriteField("type", *requestType)
			err := writer.Close()
			if err != nil {
				panic(errors.New("failed to close writer"))
			}
			req, err := http.NewRequest("POST", process_addr, regBody)
			if err != nil {
				panic(errors.New("failed to create http request"))
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())
			res, err := client.Do(req)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)

			fmt.Println(string(body))
		}
	case "list":
		{
			client := &http.Client{}
			reg, err := http.NewRequest("GET", check_addr, nil)
			if err != nil {
				panic(errors.New("failed to create http request"))
			}
			resp, err := client.Do(reg)
			if err != nil {
				panic(errors.New("failed to do http request"))
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(errors.New("failed to ready request body"))
			}
			fmt.Println(string(body))
		}
	default:
		{
			panic(errors.New("undefined type request!"))
		}
	}
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
