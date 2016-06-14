package smaug

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

func TokenFromRequest(request *http.Request) (*string, error) {
	token_from_cookie, err := request.Cookie("access_token")
	if err == nil && token_from_cookie != nil {
		return &token_from_cookie.Value, nil
	}

	token_from_query := request.URL.Query().Get("access_token")
	if len(token_from_query) > 0 {
		return &token_from_query, nil
	}

	return nil, errors.New("no access token")
}

func Authenticate(u url.URL, token *string) (*Identity, error) {
	q := make(url.Values)
	q.Add("token", *token)

	u.Path = "/configuration"
	u.RawQuery = q.Encode()

	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("forbidden")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var context Context
	err = json.Unmarshal(body, &context)
	if err != nil {
		return nil, err
	}

	return &context.User, nil
}
