package igprotocol

// SessionKeys ...
type SessionKeys struct {
	SecurityToken string
	CST           string
}

// HTTPResponseHeader ...
type HTTPResponseHeader struct {
	Status           string
	StatusCode       int
	Success          bool
	ErrorCode        string
	ErrorDescription string
	SecurityToken    string
}
