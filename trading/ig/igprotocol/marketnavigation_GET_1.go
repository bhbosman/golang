package igprotocol

// Copyright 2016 The golangtradingig Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is an implementation of the IG Trading platform
// Command: marketnavigation
// Method: GET
// Version: 1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SendMarketNavigationRequest ...
func SendMarketNavigationRequest(
	conn IGConnection,
	inputData MarketNavigationRequest,
	headerKeys map[string]string) (*MarketNavigationResponseResult, error) {

	conn.GetLogger().Log("SendMarketNavigationRequest start...")
	defer conn.GetLogger().Log("SendMarketNavigationRequest finished.")

	request, err := CreateMarketNavigationRequest(conn.GetURL(), inputData, headerKeys)
	if err != nil {
		return nil, err
	}

	response, err := conn.GetConnection().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	ResultData := MarketNavigationResponse{}

	var bodyBytes []byte
	var jsonData []byte
	if response.ContentLength != 0 {
		if conn.GetLogger().GetLogEnabled() {
			bodyBytes, err = ioutil.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}
			byteReader := bytes.NewReader(bodyBytes)
			jsonEncoder := json.NewDecoder(byteReader)
			err = jsonEncoder.Decode(&ResultData)
			if err != nil {
				return nil, err
			}
			jsonData, _ = json.Marshal(ResultData)
		} else {
			jsonEncoder := json.NewDecoder(response.Body)
			err = jsonEncoder.Decode(&ResultData)
			if err != nil {
				return nil, err
			}
		}
	}
	conn.GetLogger().Logf("Raw data received: %s", bodyBytes)
	conn.GetLogger().Logf("Translated data received: %s", jsonData)

	ResultHeader := HTTPResponseHeader{
		SecurityToken:    response.Header.Get("X-SECURITY-TOKEN"),
		Success:          response.StatusCode == 200,
		ErrorCode:        ResultData.ErrorCode,
		ErrorDescription: ErrorDescription(ResultData.ErrorCode),
		Status:           response.Status,
		StatusCode:       response.StatusCode}

	result := &MarketNavigationResponseResult{
		Data: ResultData,

		Header: ResultHeader}

	return result, nil
}

// CreateMarketNavigationRequest ...
func CreateMarketNavigationRequest(BaseURL string,
	inputData MarketNavigationRequest,
	headerKeys map[string]string) (*http.Request, error) {

	// conn.GetLogger().Log("MarketNavigation start...")
	// defer conn.GetLogger().Log("MarketNavigation finished.")

	header := http.Header{}
	header.Add(ContentTypeConst, "application/json; charset=UTF-8")
	header.Add(AcceptConst, "application/json; charset=UTF-8")
	header.Add(VersionConst, "1")

	for key, value := range headerKeys {
		header.Add(key, value)
	}

	sURL := fmt.Sprintf("%s/%s%s", BaseURL, "marketnavigation", inputData.URLPostFix())
	// conn.GetLogger().Log(fmt.Sprintf("URL: %s, Command: %s", sURL, "GET"))

	URL, err := url.Parse(sURL)
	if err != nil {
		return nil, err
	}

	// Create web request
	request := &http.Request{
		Method:     "GET",
		URL:        URL,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,

		Host: URL.Host}

	return request, nil
}
