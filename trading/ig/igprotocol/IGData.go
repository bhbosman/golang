package igprotocol

// IGData ...
type IGData struct {
	ConnectionContext     IGContext
	ConnectionInformation AuthenticationResponse
	Markets               []Market
}

// Login ...
func (ctx *IGData) Login() (*AuthenticationResponseResult, error) {
	result, err := ctx.ConnectionContext.login()
	if err != nil {
		ctx.ConnectionContext.Log.Logf("Error while logging in. Error: %s", err)
		return nil, err
	}
	ctx.ConnectionInformation = result.Data
	//
	return result, nil
}

// Logout ...
func (ctx *IGData) Logout() (*LogoutResponseResult, error) {
	result, err := ctx.ConnectionContext.logout()
	if err != nil {
		ctx.ConnectionContext.Log.Logf("Error while logging out. Error: %s", err)
		return nil, err
	}

	ctx.ConnectionInformation = AuthenticationResponse{}
	//
	return result, nil
}

// GetDefaultAccountAccount ...
func (ctx *IGData) GetDefaultAccountAccount() *Account {
	for _, acc := range ctx.ConnectionInformation.Account {
		if acc.Preferred {
			return &acc
		}
	}
	//
	return nil
}

// GetNonDefaultAccountAccounts ...
func (ctx *IGData) GetNonDefaultAccountAccounts() []Account {
	result := make([]Account, 0, len(ctx.ConnectionInformation.Account))
	for _, acc := range ctx.ConnectionInformation.Account {
		if !acc.Preferred {
			result = append(result, acc)
		}
	}
	//
	return result
}

// AccountSwitch ...
func (ctx *IGData) AccountSwitch(AccountID string, switchToDefault bool) (*AccountSwitchResponseResult, error) {
	result, err := ctx.ConnectionContext.accountSwitch(AccountID, switchToDefault)
	//
	return result, err
}

// GetMarkets ...
func (ctx *IGData) GetMarkets(nodeID string, rescursive bool) ([]Market, error) {
	resultData := make([]Market, 0, 16)
	result, err := ctx.SendMarketNavigationRequest("")
	if err != nil {
		return make([]Market, 0, 0), err
	}
	if rescursive {
		for _, node := range result.Data.Nodes {
			resurData, err := ctx.GetMarkets(node.ID, rescursive)
			if err != nil {
				return make([]Market, 0, 0), err
			}
			for _, market := range resurData {
				resultData = append(resultData, market)
			}
		}
	}
	for _, market := range result.Data.Markets {
		resultData = append(resultData, market)
	}
	//
	return resultData, nil
}

// SendMarketNavigationRequest ...
func (ctx *IGData) SendMarketNavigationRequest(nodeID string) (*MarketNavigationResponseResult, error) {
	return ctx.ConnectionContext.sendMarketNavigationRequest(nodeID)
}
