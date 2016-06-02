package igprotocol

import (
	"fmt"
	"testing"

	bhbosmanTesting "github.com/bhbosman/golang/testing"
	testvariables "github.com/bhbosman/golang/trading/ig/igprotocol/test"
)

func TestSessionPutVersion1TryingToSwitchPreferredAccountNumber(t *testing.T) {
	test := bhbosmanTesting.MyTestingT{T: t}

	ctx := NewIGContextForTesting(
		t,
		testvariables.TestProxy,
		testvariables.TestAccountIdentifier,
		testvariables.TestAccountPassword,
		testvariables.TestAccountAPIKey)

	_, err := ctx.Login()
	if err != nil {
		t.Fatalf("IG Context did not connect. Error: %s", err)
	}

	defer func() {
		_, err = ctx.Logout()
		if err != nil {
			t.Fatalf("IG Context did not disconnect. Error: %s", err)
		}
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
	test := bhbosmanTesting.MyTestingT{T: t}

	ctx := NewIGContextForTesting(
		t,
		testvariables.TestProxy,
		testvariables.TestAccountIdentifier,
		testvariables.TestAccountPassword,
		testvariables.TestAccountAPIKey)

	_, err := ctx.Login()
	if err != nil {
		t.Fatalf("IG Context did not connect. Error: %s", err)
	}
	defer func() {
		_, err = ctx.Logout()
		if err != nil {
			t.Fatalf("IG Context did not disconnect. Error: %s", err)
		}
	}()
	accounts := ctx.GetNonDefaultAccountAccounts()

	test.CheckBool(len(accounts) > 0, "Must have a non default account")

	acc := accounts[0]
	result, err := ctx.AccountSwitch(acc.AccountID, true)
	test.CheckBool(result.Header.Success, fmt.Sprintf("Must be Successfull(%s, %s)", result.Header.ErrorCode, result.Header.ErrorDescription))
	test.CheckIntWithMessage(200, result.Header.StatusCode, "Status code must be 200")
	test.CheckBool(err == nil, "Error must be assigned")
}
