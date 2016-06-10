package main

import (
	"github.com/dbcdk/go-smaug/smaug"
	"net/http"
	"net/url"
)

func main() {
	smaugUrl := url.URL{Scheme: "https", Host: "platform-i01.dbc.dk:3001"}
	token := "qwerty"
	smaug.Authenticate(smaugUrl, &token)

	urlWithToken, _ := url.Parse("https://localhost/?access_token=abc")

	token2, err := smaug.TokenFromRequest(&http.Request{URL: urlWithToken})

	if err != nil {
		println(err.Error())
	} else {
		println(*token2)
	}
}
