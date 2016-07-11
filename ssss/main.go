package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/bhbosman/golang/trading/ig/igprotocol"
	"github.com/bhbosman/golang/trading/ig/igprotocol/test"
)

func main() {

	authentication := igprotocol.AuthenticationRequest{
		Identifier: test.TestAccountIdentifier,
		Password:   test.TestAccountPassword}

	keys := make(map[string]string)
	keys[igprotocol.XIGApiKeyConst] = test.TestAccountAPIKey

	request, _ := igprotocol.CreateAuthenticationRequest(
		test.TestDemoAPI,
		authentication,
		keys)

	buffer := bytes.NewBuffer(nil)
	// request.Write(buffer)

	dd, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	AuthResponse, _ := igprotocol.CreateAuthenticationResponse(dd)

}
