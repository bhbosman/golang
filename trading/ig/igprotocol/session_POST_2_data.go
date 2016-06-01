package igprotocol

// Data Definition
// http://labs.ig.com/rest-trading-api-reference/service-detail?id=261

// AuthenticationRequest ...
type AuthenticationRequest struct {
	EncryptedPassword bool   `json:"encryptedPassword"`
	Identifier        string `json:"identifier"`
	Password          string `json:"password"`
}

// AccountInfo ...
type AccountInfo struct {
	Balance    float64 `json:"balance"`
	Deposit    float64 `json:"deposit"`
	ProfitLoss float64 `json:"profitLoss"`
	Available  float64 `json:"available"`
}

// Account ...
type Account struct {
	AccountID   string `json:"accountId"`
	AccountName string `json:"accountName"`
	Preferred   bool   `json:"preferred"`
	AccountType string `json:"accountType"`
}

// AuthenticationResponse ...
type AuthenticationResponse struct {
	ErrorCode             string      `json:"errorCode"`
	AccountInfo           AccountInfo `json:"accountInfo"`
	AccountType           string      `json:"accountType"`
	Account               []Account   `json:"accounts"`
	AuthenticationStatus  string      `json:"authenticationStatus"`
	ClientID              string      `json:"clientId"`
	CurrencyIsoCode       string      `json:"currencyIsoCode"`
	CurrentAccountID      string      `json:"currentAccountId"`
	LightStreamerEndpoint string      `json:"lightstreamerEndpoint"`
	CurrencySymbol        string      `json:"currencySymbol"`
	DealingEnabled        bool        `json:"dealingEnabled"`
	Encrypted             bool        `json:"encrypted"`
	HasActiveDemoAccounts bool        `json:"hasActiveDemoAccounts"`
	HasActiveLiveAccounts bool        `json:"hasActiveLiveAccounts"`
	IgCompany             string      `json:"igCompany"`
	ReroutingEnvironment  string      `json:"reroutingEnvironment"`
	TimezoneOffset        int         `json:"timezoneOffset"`
	TrailingStopsEnabled  bool        `json:"trailingStopsEnabled"`
}

// AuthenticationResponseResult ...
type AuthenticationResponseResult struct {
	Header      HTTPResponseHeader
	Data        AuthenticationResponse
	SessionKeys SessionKeys
}
