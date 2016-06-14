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

	cookie := http.Cookie{Name: "access_token", Value: "qwerty"}
	request, err := http.NewRequest("GET", "https://localhost/?access_token=abc", nil)
	request.AddCookie(&cookie)

	token2, err := smaug.TokenFromRequest(request)

	if err != nil {
		println(err.Error())
	} else {
		println(*token2)
	}
}
