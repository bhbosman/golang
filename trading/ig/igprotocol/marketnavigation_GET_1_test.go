package igprotocol

import (
	"encoding/json"
	"fmt"
	"testing"

	bhbosmanTesting "github.com/bhbosman/golang/testing"
	testvariables "github.com/bhbosman/golang/trading/ig/igprotocol/test"
)

func TestMarketNavigationAll(t *testing.T) {
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

	keys := ctx.ConnectionContext.CreateConnectionContext()
	response, err := SendMarketNavigationRequest(
		&ctx.ConnectionContext,
		MarketNavigationRequest{},
		keys)
	if err != nil {
		t.Fatalf("A. Error: %s", err)
	}

	test.CheckBool(response.Header.Success, response.Header.ErrorCode)
}

func TestMarketNavigationPerInstance(t *testing.T) {
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

	result, err := ctx.SendMarketNavigationRequest("")
	test.CheckBool(result.Header.Success, fmt.Sprintf("Must be Successfull(%s, %s)", result.Header.ErrorCode, result.Header.ErrorDescription))
	test.CheckIntWithMessage(200, result.Header.StatusCode, "Status code must be 200")
	test.CheckBool(err == nil, "Error must be assigned")

	test.CheckBool(len(result.Data.Nodes) > 0, "Node count > 0")
	for i, node := range result.Data.Nodes {
		test.Logf("Instance %d: NodeID: %s, Node Name: %s", i, node.ID, node.Name)
		//
		result, err = ctx.SendMarketNavigationRequest(node.ID)
		//
		test.CheckBool(result.Header.Success, fmt.Sprintf("Must be Successfull(%s, %s)", result.Header.ErrorCode, result.Header.ErrorDescription))
		test.CheckIntWithMessage(200, result.Header.StatusCode, "Status code must be 200")
		test.CheckBool(err == nil, "Error must be assigned")
	}
}

func TestGetMarkets(t *testing.T) {

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

	result, err := ctx.GetMarkets("", true)
	t.Log(json.Marshal(result))
}
