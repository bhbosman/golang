package igprotocol

// Data Definition
// http://labs.ig.com/rest-trading-api-reference/service-detail?id=261

// AccountSwitchRequest ...
type AccountSwitchRequest struct {
	AccountID      string `json:"accountId"`
	DefaultAccount bool   `json:"defaultAccount"`
}

// AccountSwitchResponse ...
type AccountSwitchResponse struct {
	ErrorCode             string `json:"errorCode"`
	DealingEnabled        bool   `json:"dealingEnabled"`
	HasActiveDemoAccounts bool   `json:"hasActiveDemoAccounts"`
	HasActiveLiveAccounts bool   `json:"hasActiveLiveAccounts"`
	TrailingStopsEnabled  bool   `json:"trailingStopsEnabled"`
}

// AccountSwitchResponseResult ...
type AccountSwitchResponseResult struct {
	Header HTTPResponseHeader
	Data   AccountSwitchResponse
}
