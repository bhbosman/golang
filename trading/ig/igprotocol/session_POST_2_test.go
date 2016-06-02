package igprotocol

import (
	"testing"

	bhbosmanTesting "github.com/bhbosman/golang/testing"
	testvariables "github.com/bhbosman/golang/trading/ig/igprotocol/test"
)

func TestSessionPostVersion2(t *testing.T) {
	test := bhbosmanTesting.MyTestingT{T: t}

	IGConnection := NewIGContextForTesting(
		t,
		testvariables.TestProxy,
		"",
		"",
		"")

	// Send logon
	credentials := AuthenticationRequest{
		Identifier: testvariables.TestAccountIdentifier,
		Password:   testvariables.TestAccountPassword}

	keys := make(map[string]string)
	keys[XIGApiKeyConst] = testvariables.TestAccountAPIKey

	responseLogin, err := SendAuthenticationRequest(&IGConnection.ConnectionContext, credentials, keys)
	if err != nil {
		t.Fatalf("Error with SendAuthenticationRequest. Error: %s", err)
	}

	defer func() {
		responseLogout, err := SendLogoutRequest(&IGConnection.ConnectionContext, keys)

		if err != nil {
			t.Fatalf("Error with SendLogoutRequest. Error: %s", err)
		}
		test.CheckIntWithMessage(responseLogout.Header.StatusCode, 204, "StatusCode wrong")
		test.CheckBool(responseLogout.Header.Success, "Not success")
	}()

	test.CheckIntWithMessage(responseLogin.Header.StatusCode, 200, "StatusCode wrong")
	test.CheckBool(responseLogin.Header.Success, "Not success")

	keys = make(map[string]string)
	keys[XIGApiKeyConst] = testvariables.TestAccountAPIKey
	keys[XSecurityTokenConst] = responseLogin.SessionKeys.SecurityToken
	keys[CstConst] = responseLogin.SessionKeys.CST
}
