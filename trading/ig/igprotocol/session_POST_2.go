package igprotocol

// Copyright 2016 The golangtradingig Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is an implementation of the IG Trading platform
// Command: session
// Method: POST
// Version: 2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SendAuthenticationRequest ...
func SendAuthenticationRequest(
	conn IGConnection,
	inputData AuthenticationRequest,
	headerKeys map[string]string) (*AuthenticationResponseResult, error) {

	conn.GetLogger().Log("SendAuthenticationRequest start...")
	defer conn.GetLogger().Log("SendAuthenticationRequest finished.")

	request, err := CreateAuthenticationRequest(conn, inputData, headerKeys)
	if err != nil {
		return nil, err
	}

	response, err := conn.GetConnection().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	SessionKeys := SessionKeys{
		SecurityToken: response.Header.Get(XSecurityTokenConst),
		CST:           response.Header.Get(CstConst)}

	ResultData := AuthenticationResponse{}

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

	result := &AuthenticationResponseResult{
		Data:        ResultData,
		SessionKeys: SessionKeys,
		Header:      ResultHeader}

	return result, nil
}

// CreateAuthenticationRequest ...
func CreateAuthenticationRequest(conn IGConnection,
	inputData AuthenticationRequest,
	headerKeys map[string]string) (*http.Request, error) {

	conn.GetLogger().Log("Authentication start...")
	defer conn.GetLogger().Log("Authentication finished.")

	header := http.Header{}
	header.Add(ContentTypeConst, "application/json; charset=UTF-8")
	header.Add(AcceptConst, "application/json; charset=UTF-8")
	header.Add(VersionConst, "2")

	for key, value := range headerKeys {
		header.Add(key, value)
	}
	sURL := fmt.Sprintf("%s/%s%s", conn.GetURL(), "session", "")
	conn.GetLogger().Log(fmt.Sprintf("URL: %s, Command: %s", sURL, "POST"))

	URL, err := url.Parse(sURL)
	if err != nil {
		return nil, err
	}
	arrayOfBytes, _ := json.Marshal(inputData)

	// Create web request
	request := &http.Request{
		Method:        "POST",
		URL:           URL,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        header,
		ContentLength: int64(len(arrayOfBytes)),
		Body:          ioutil.NopCloser(bytes.NewBuffer(arrayOfBytes)),
		Host:          URL.Host}

	return request, nil
}
