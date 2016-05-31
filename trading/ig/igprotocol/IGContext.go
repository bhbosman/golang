package igprotocol

import (
	"fmt"
	"net/http"
	"os"
	
)

// IGContext ...
type IGContext struct {
	UserName              string
	Password              string
	APIKey                string
	SecurityToken         string
	CST                   string
	
	Connection            *http.Client
	Log                   Logger
}

// GetLogger ...
func (ctx *IGContext) GetLogger() Logger {
	return ctx.Log
}

// GetURL ...
func (ctx *IGContext) GetURL() string {
	return "https://demo-api.ig.com/gateway/deal"
}

// GetConnection ...
func (ctx *IGContext) GetConnection() *http.Client {
	return ctx.Connection
}

// LogToDevice ...
func (ctx *IGContext) LogToDevice(s string) {
	os.Stdout.WriteString(s)
}

// DealWithReturningHeader ...
func (ctx *IGContext) DealWithReturningHeader(header HTTPResponseHeader) {
	if header.SecurityToken != ctx.SecurityToken  &&  header.SecurityToken != "" {
		ctx.SecurityToken = header.SecurityToken
		ctx.GetLogger().Logf("Change SecurityToken to: %s", ctx.SecurityToken)
	}
}


// CreateConnectionContext ...
func (ctx *IGContext) CreateConnectionContext() map[string]string {
	result := make(map[string]string)
	result[XIGApiKeyConst] = ctx.APIKey
	result[XSecurityTokenConst] = ctx.SecurityToken
	result[CstConst] = ctx.CST
	//
	return result
}

// Login ...
func (ctx *IGContext) login() (*AuthenticationResponseResult, error) {
	ctx.Log.Log("Login start...")
	defer ctx.Log.Log("Login finished.")
	//
	credentials := AuthenticationRequest{
		Identifier: ctx.UserName,
		Password:   ctx.Password}

	keys := make(map[string]string)
	keys[XIGApiKeyConst] = ctx.APIKey

	result, err := SendAuthenticationRequest(ctx, credentials, keys)
	if err != nil {
		return result, err
	}
	if !result.Header.Success {
		err := fmt.Errorf(
			"Login was not successfull. Code: %d, Error Message: %s, Error Description: %s",
			result.Header.StatusCode,
			result.Header.ErrorCode,
			result.Header.ErrorDescription)
		return result, err
	}
	ctx.SecurityToken = result.SessionKeys.SecurityToken
	ctx.CST = result.SessionKeys.CST
	//
	return result, nil
}

// Logout ...
func (ctx *IGContext) logout() (*LogoutResponseResult, error) {
	ctx.Log.Log("Logout start...")
	defer ctx.Log.Log("Logout finished.")
	//
	keys := ctx.CreateConnectionContext()

	result, err := SendLogoutRequest(ctx, keys)
	if err != nil {
		return result, err
	}
	if !result.Header.Success {
		err := fmt.Errorf(
			"Logout was not successfull. Code: %d, Error Message: %s, Error Description: %s",
			result.Header.StatusCode,
			result.Header.ErrorCode,
			result.Header.ErrorDescription)
		return result, err
	}
	ctx.SecurityToken = ""
	ctx.CST = ""
	//
	return result, nil
}

// AccountSwitch ...
func (ctx *IGContext) accountSwitch(AccountID string, switchToDefault bool) (*AccountSwitchResponseResult, error) {
	ctx.Log.Log("AccountSwitch start...")
	defer ctx.Log.Log("AccountSwitch finished.")

	inputData := AccountSwitchRequest{
		AccountID:      AccountID,
		DefaultAccount: switchToDefault}

	keys := ctx.CreateConnectionContext()
	result, err := SendAccountSwitchRequest(ctx, inputData, keys)
	if err != nil {
		return result, err
	}
	if !result.Header.Success {
		err := fmt.Errorf(
			"AccountSwitch was not successfull. Code: %d, Error Message: %s, Error Description: %s",
			result.Header.StatusCode,
			result.Header.ErrorCode,
			result.Header.ErrorDescription)
		return result, err
	}
	ctx.DealWithReturningHeader(result.Header)
	//
	return result, nil
}

// SendMarketNavigationRequest ...
func (ctx *IGContext) sendMarketNavigationRequest(nodeID string) (*MarketNavigationResponseResult, error) {

	ctx.Log.Log("SendMarketNavigationRequest start...")
	defer ctx.Log.Log("SendMarketNavigationRequest finished.")

	inputData := MarketNavigationRequest{
		NodeID: nodeID}

	keys := ctx.CreateConnectionContext()
	result, err := SendMarketNavigationRequest(
		ctx,
		inputData,
		keys)
	if err != nil {
		return result, err
	}
	if !result.Header.Success {
		err := fmt.Errorf(
			"SendMarketNavigationRequest was not successfull. Code: %d, Error Message: %s, Error Description: %s",
			result.Header.StatusCode,
			result.Header.ErrorCode,
			result.Header.ErrorDescription)
		return result, err
	}
	ctx.DealWithReturningHeader(result.Header)
	//
	return result, nil
}
