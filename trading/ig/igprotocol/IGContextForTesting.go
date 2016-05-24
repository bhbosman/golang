package igprotocol

import (
	"net/http"
	"testing"
	
)

// NewIGContextForTesting ...
func NewIGContextForTesting(UserName, Password, APIKey string, t *testing.T) *IGData {
	return &IGData{
		ConnectionInformation: AuthenticationResponse{},
		ConnectionContext: IGContext{
			UserName:              UserName,
			Password:              Password,
			APIKey:                APIKey,
			CST:                   "",
			SecurityToken:         "",
			Connection:            &http.Client{},
			Log: &TestingLogging{
				LogEnabled: true,
				TestUnit:   t}}}
}
