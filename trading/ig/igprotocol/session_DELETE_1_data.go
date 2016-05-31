package igprotocol

// Data Definition
// http://labs.ig.com/rest-trading-api-reference/service-detail?id=261

// LogoutResponse ...
type LogoutResponse struct {
	ErrorCode string `json:"errorCode"`
}

// LogoutResponseResult ...
type LogoutResponseResult struct {
	Header HTTPResponseHeader
	Data   LogoutResponse
}


