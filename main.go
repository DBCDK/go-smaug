package main

import (
	"github.com/dbcdk/go-smaug/smaug"
	"net/http"
	"net/url"
)

func main() {
	smaugUrl := url.URL{Scheme: "https", Host: "platform-i01.dbc.dk:3001"}
	smaug.Authenticate(smaugUrl, "qwerty")

	token, err := smaug.TokenFromRequest(&http.Request{URL: &smaugUrl})
	println(token)
	println(err)
}
