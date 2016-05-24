package igprotocol


import (
    "net/http"
)

// IGConnection ... 
type IGConnection interface {
    GetConnection() *http.Client
    GetURL() string
    GetLogger() Logger
    
}


// Logger ...
type Logger interface {
    GetLogEnabled() bool
    Logf(format string, args... interface{})
    Log(s string)
}




