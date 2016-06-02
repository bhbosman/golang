package igprotocol

import (
	"net/http"
	"net/url"
	"testing"
)

// NewIGContextForTesting ...
func NewIGContextForTesting(
	t *testing.T,
	proxy string,
	UserName string,
	Password string,
	APIKey string) *IGData {

	var connection *http.Client

	if proxy != "" {
		proxyURL, _ := url.Parse(proxy)
		connection = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL)}}

	} else {
		connection = &http.Client{}
	}

	return &IGData{
		ConnectionInformation: AuthenticationResponse{},
		ConnectionContext: IGContext{
			UserName:      UserName,
			Password:      Password,
			APIKey:        APIKey,
			CST:           "",
			SecurityToken: "",
			Connection:    connection,
			Log: &TestingLogging{
				LogEnabled: true,
				TestUnit:   t}}}
}
