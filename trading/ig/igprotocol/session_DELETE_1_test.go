package igprotocol

import (
	"testing"

	bhbosmanTesting "github.com/bhbosman/golang/testing"
	testvariables "github.com/bhbosman/golang/trading/ig/igprotocol/test"
)

func TestSessionDeleteVersion2(t *testing.T) {
	test := bhbosmanTesting.MyTestingT{T: t}

	IGConnection := NewIGContextForTesting("", "", "", t)

	// Send logon
	credentials := AuthenticationRequest{
		Identifier: testvariables.TestAccountIdentifier,
		Password:   testvariables.TestAccountPassword}

	keys := make(map[string]string)
	keys[XIGApiKeyConst] = testvariables.TestAccountAPIKey

	responseLogin, err := SendAuthenticationRequest(&IGConnection.ConnectionContext, credentials, keys)
	test.CheckErrorWithMessage(err, "Error with SendAuthenticationRequest")
	test.CheckIntWithMessage(responseLogin.Header.StatusCode, 200, "StatusCode wrong")
	test.CheckBool(responseLogin.Header.Success, "Not success")

	keys = make(map[string]string)
	keys[XIGApiKeyConst] = testvariables.TestAccountAPIKey
	keys[XSecurityTokenConst] = responseLogin.SessionKeys.SecurityToken
	keys[CstConst] = responseLogin.SessionKeys.CST

	responseLogout, err := SendLogoutRequest(&IGConnection.ConnectionContext, keys)

	test.CheckErrorWithMessage(err, "Error with SendLogoutRequest")
	test.CheckIntWithMessage(responseLogout.Header.StatusCode, 204, "StatusCode wrong")
	test.CheckBool(responseLogout.Header.Success, "Not success")

}
