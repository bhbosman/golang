package main

import (
	"fmt"

	"github.com/bhbosman/golang/trading/ig/igprotocol"
	testvariables "github.com/bhbosman/golang/trading/ig/igprotocol/test"
)

func main() {

	// Send logon
	credentials := igprotocol.AuthenticationRequest{
		Identifier: testvariables.TestAccountIdentifier,
		Password:   testvariables.TestAccountPassword}
	keys := make(map[string]string)
	keys[igprotocol.XIGApiKeyConst] = testvariables.TestAccountAPIKey

	response, err := igprotocol.SendAuthenticationRequest()

	request, err := igprotocol.CreateAuthenticationRequest("", credentials, keys)
	if err != nil {
		fmt.Println(err)
		return
	}

}
