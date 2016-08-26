package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"bytes"

	"github.com/bhbosman/golang/trading/ig/igprotocol"
	test_env "github.com/bhbosman/golang/trading/ig/igprotocol/test"
)

type LSContext struct {
	UserName               string
	Password               string
	APIKey                 string
	URL                    string
	Conn                   *http.Client
	HeaderKeys             map[string]string
	AuthenticationResponse igprotocol.AuthenticationResponse
	SessionKeys            igprotocol.SessionKeys
}

func createLSConnection(proxy string) *http.Client {
	if proxy != "" {
		proxyURL, _ := url.Parse(proxy)
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL)}}

	}
	return &http.Client{}
}

func NewLSConText(username, password, APIKey, URL, proxy string) *LSContext {

	ctx := &LSContext{
		UserName:   username,
		Password:   password,
		APIKey:     APIKey,
		URL:        URL,
		Conn:       createLSConnection(proxy),
		HeaderKeys: make(map[string]string)}
	//
	ctx.HeaderKeys[igprotocol.XIGApiKeyConst] = ctx.APIKey
	//
	return ctx
}

type SendFunc func(response *http.Response) (igprotocol.HTTPResponseHeader, error)

func (ctx *LSContext) Send2(request *http.Request, f SendFunc) (igprotocol.HTTPResponseHeader, error) {
	response, err := ctx.Send(request)
	if err != nil {
		return igprotocol.HTTPResponseHeader{}, err
	}
	defer response.Body.Close()
	if f != nil {
		return f(response)
	}
	return igprotocol.HTTPResponseHeader{}, err
}

func (ctx *LSContext) Send(request *http.Request) (*http.Response, error) {
	return ctx.Conn.Do(request)
}

func (ctx *LSContext) Connect() (igprotocol.HTTPResponseHeader, error) {
	credentials := igprotocol.AuthenticationRequest{
		Identifier: ctx.UserName,
		Password:   ctx.Password}
	request, err := igprotocol.CreateAuthenticationRequest(
		ctx.URL,
		credentials,
		ctx.HeaderKeys)
	if err != nil {
		return igprotocol.HTTPResponseHeader{}, err
	}
	return ctx.Send2(
		request,
		func(response *http.Response) (igprotocol.HTTPResponseHeader, error) {
			AuthResponse, err := igprotocol.CreateAuthenticationResponse(response)
			if err != nil {
				return igprotocol.HTTPResponseHeader{}, err
			}
			ctx.AuthenticationResponse = AuthResponse.Data
			ctx.SessionKeys = AuthResponse.SessionKeys
			ctx.HeaderKeys[igprotocol.XSecurityTokenConst] = AuthResponse.SessionKeys.SecurityToken
			ctx.HeaderKeys[igprotocol.CstConst] = AuthResponse.SessionKeys.CST
			return AuthResponse.Header, nil
		})
}

func (ctx *LSContext) DisConnect() (igprotocol.HTTPResponseHeader, error) {
	request, err := igprotocol.CreateLogoutRequest(
		ctx.URL,
		ctx.HeaderKeys)
	if err != nil {
		return igprotocol.HTTPResponseHeader{}, err
	}
	return ctx.Send2(
		request,
		func(response *http.Response) (igprotocol.HTTPResponseHeader, error) {
			if err != nil {
				return igprotocol.HTTPResponseHeader{}, err
			}
			defer response.Body.Close()

			LogoutResponse, err := igprotocol.CreateLogoutResponse(response)
			if err != nil {
				return igprotocol.HTTPResponseHeader{}, err
			}
			return LogoutResponse.Header, nil
		})
}

func (ctx *LSContext) ActiveUser() string {
	for _, account := range ctx.AuthenticationResponse.Account {
		if account.Preferred {
			return account.AccountID
		}
	}
	return ""
}

func (ctx *LSContext) ConnectToLS() (igprotocol.HTTPResponseHeader, error) {
	sURL := fmt.Sprintf("%s/%s", ctx.AuthenticationResponse.LightStreamerEndpoint, "lightstreamer/create_session.txt")

	URL, _ := url.Parse(sURL)
	q := url.Values{}
	q.Set("LS_op2", "create")
	q.Set("LS_adapter_set", "DEFAULT")
	q.Set("LS_cid", "mgQkwtwdysogQz2BJ4Ji kOj2Bg")
	// q.Set("LS_user", ctx.ActiveUser())
	q.Set("LS_password", fmt.Sprintf("CST-%s|XST-%s", ctx.SessionKeys.CST, ctx.SessionKeys.SecurityToken))

	byteBuffer := bytes.NewBufferString(q.Encode())
	fmt.Println(q)
	fmt.Println(byteBuffer)

	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	request := &http.Request{
		Method:        "POST",
		URL:           URL,
		Header:        header,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(byteBuffer.Len()),
		Body:          ioutil.NopCloser(bytes.NewBuffer(byteBuffer.Bytes())),
		Host:          URL.Host}

	return ctx.Send2(
		request,
		func(response *http.Response) (igprotocol.HTTPResponseHeader, error) {
			fmt.Println(response)
			fmt.Println(response.ContentLength)
			if response.ContentLength != 0 {

				b, _ := ioutil.ReadAll(response.Body)
				s := string(b)

				fmt.Println(s)
			}

			ResultHeader := igprotocol.HTTPResponseHeader{
				SecurityToken: response.Header.Get("X-SECURITY-TOKEN"),
				Success:       response.StatusCode == 200,
				Status:        response.Status,
				StatusCode:    response.StatusCode}
			return ResultHeader, nil
		})
}

func main() {
	ctx := NewLSConText(
		test_env.TestAccountIdentifier,
		test_env.TestAccountPassword,
		test_env.TestAccountAPIKey,
		test_env.TestDemoAPI,
		"")

	fmt.Println("Connect...")
	result, err := ctx.Connect()
	if err != nil {
		fmt.Printf("Error with connect. Error: %s\n", err)
		return
	}
	if !result.Success {
		fmt.Printf("Result failed with connect. Error: %s\n", result.ErrorCode)
		return
	}

	fmt.Println("Connect to LS")

	ctx.ConnectToLS()

	fmt.Println("DisConnect...")
	result, err = ctx.DisConnect()
	if err != nil {
		fmt.Printf("Error with connect. Error: %s\n", err)
	}

}
