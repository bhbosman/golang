package igprotocol

import (
	"fmt"
	"testing"
	
	testvariables "bitbucket.org/bhbosman/golangtradingig/igprotocol/test"
)

func TestSessionPutVersion1TryingToSwitchPreferredAccountNumber(t *testing.T) {
	test := MyTestingT{t}

	ctx := NewIGContextForTesting(
		testvariables.TestAccountIdentifier, 
		testvariables.TestAccountPassword, 
		testvariables.TestAccountAPIKey, t)

	_, err := ctx.Login()
	test.CheckErrorWithMessage(err, "IG Context did not connect")

	defer func() {
		_, err = ctx.Logout()
		test.CheckErrorWithMessage(err, "IG Context did not disconnect")
	}()

	acc := ctx.GetDefaultAccountAccount()

	if acc == nil {
		test.Fatal("No account number found")
	}

	result, err := ctx.AccountSwitch(acc.AccountID, true)
	test.CheckBool(!result.Header.Success, "Must not be Successfull")
	test.CheckIntWithMessage(412, result.Header.StatusCode, "Status code must be 412")
	test.CheckBool(err != nil, "Error must be assigned")
}


func TestSessionPutVersion1TryingToSwitchOtherAccountNumber(t *testing.T) {
	test := MyTestingT{t}

	ctx := NewIGContextForTesting(
		testvariables.TestAccountIdentifier, 
		testvariables.TestAccountPassword, 
		testvariables.TestAccountAPIKey, 
		t)
	

	_, err := ctx.Login()
	defer func() {
		_, err = ctx.Logout()
		test.CheckErrorWithMessage(err, "IG Context did not disconnect")
	}()
	test.CheckErrorWithMessage(err, "IG Context did not connect")
	accounts := ctx.GetNonDefaultAccountAccounts()

	test.CheckBool(len(accounts) > 0, "Must have a non default account")
	
	acc := accounts[0]
	result, err := ctx.AccountSwitch(acc.AccountID, true)
	test.CheckBool(result.Header.Success, fmt.Sprintf("Must be Successfull(%s, %s)", result.Header.ErrorCode, result.Header.ErrorDescription))
	test.CheckIntWithMessage(200, result.Header.StatusCode, "Status code must be 200")
	test.CheckBool(err == nil, "Error must be assigned")
}