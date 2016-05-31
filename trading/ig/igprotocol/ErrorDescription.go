package igprotocol

// Data Definition
// http://labs.ig.com/rest-trading-api-reference/service-detail?id=261

// ErrorDescription ...
func ErrorDescription(key string) string {
	data := make(map[string]string)
	data["error.public-api.failure.encryption.required"] = "A login has been attempted to the login V1 service by a client from the IG Singapore company. They need to use the v2 version as they need to send their passwords encrypyted."
	data["invalid.input"] = "A generic input data error has occurred"
	data["authentication.failure.not-a-client-account"] = "The account is not a valid client account"
	data["error.public-api.failure.kyc.required"] = "The account is not allowed to log into public API. Please use the web platform."
	data["error.public-api.failure.missing.credentials"] = "The user has not provided all required security credentials."
	data["error.public-api.failure.pending.agreements.required"] = "The account is not allowed to log into public API. Please use the web platform."
	data["error.public-api.failure.preferred.account.disabled"] = "The user's preferred account is disabled."
	data["error.public-api.failure.preferred.account.not.set"] = "The user has not set a preferred account."
	data["error.security.account-access-denied"] = "The account has been denied login privileges"
	data["error.security.account-migrated"] = "The account has been migrated to the client-account model, please authenticate with the client credentials"
	data["error.security.account-not-yet-activated"] = "The account has not been activated yet"
	data["error.security.account-suspended"] = "The account has been suspended"
	data["error.security.account-token-invalid"] = "The service requires an account token and the one provided was not valid"
	data["error.security.account-token-missing"] = "The service requires an account token and it was not provided"
	data["error.security.all-accounts-pending"] = "All of the accounts are in a pending state"
	data["error.security.all-accounts-suspended"] = "All of the clients accounts have been suspended"
	data["error.security.client-suspended"] = "The client has been suspended from using the platform"
	data["error.security.client-token-invalid"] = "The service requires a client token and the one provided was not valid"
	data["error.security.client-token-missing"] = "The service requires a client token and it was not provided"
	data["error.security.generic"] = "An unexpected error has been encountered on the server side, cannot proceed. Please contact the support."
	data["error.security.invalid-application"] = "The provided user agent string is not valid"
	data["error.security.invalid-details"] = "The credentials used to authenticate the users are not valid, login is rejected"
	data["error.security.invalid-website"] = "This site is not accessible via the API services"
	data["error.security.too-many-failed-attempts"] = "Maximum number of failed login attempts have been reached"
	data["endpoint.unavailable.for.api-key"] = "The provided api key was not accepted"
	data["error.public-api.exceeded-account-allowance"] = "The account traffic allowance has been exceeded"
	data["error.public-api.exceeded-account-historical-data-allowance"] = "The account historical data traffic allowance has been exceeded"
	data["error.public-api.exceeded-account-trading-allowance"] = "The account trading traffic allowance has been exceeded"
	data["error.public-api.exceeded-api-key-allowance"] = "The api key traffic allowance has been exceeded"
	data["error.public-api.failure.stockbroking-not-supported"] = "Stockbroking not supported for Public API users."
	data["error.security.api-key-disabled"] = "The provided api key was not accepted because it is not currently enabled"
	data["error.security.api-key-invalid"] = "The provided api key was not accepted"
	data["error.security.api-key-missing"] = "The api key was not provided"
	data["error.security.api-key-restricted"] = "The provided api key was not valid for the requesting account"
	data["error.security.api-key-revoked"] = "The provided api key was not accepted because it has been revoked"
	data["invalid.url"] = ""
	data["system.error"] = ""

	result, found := data[key]
	if found {
		return result
	}
	return "(Not Found)"
}
