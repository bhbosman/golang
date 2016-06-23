package main

import (
	"bufio"
	"bytes"
	"fmt"

	"net"

	"github.com/bhbosman/golang/trading/ig/igprotocol"
	testvariables "github.com/bhbosman/golang/trading/ig/igprotocol/test"
)

func main() {

	// Send logon
	credentials := igprotocol.AuthenticationRequest{
		Identifier: testvariables.TestAccountIdentifier,
		Password:   testvariables.TestAccountPassword}
	keys := make(map[string]string)
	keys[igprotocol.XIGApiKeyConst] = testvariables.TestAccountAPIKey
	request, _ := igprotocol.CreateAuthenticationRequest("", credentials, keys)

	buffer := &bytes.Buffer{}

	request.Write(buffer)

	fmt.Println(buffer.Bytes())
	fmt.Println(buffer)

	// dd = "https://demo-apd.marketdatasystems.com"

	conn, err := net.Dial("tcp", "https://demo-apd.marketdatasystems.com")
	if err != nil {
		// handle error
	}
	conn.Write(buffer.Bytes())
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...
    fmt.Println()

}
