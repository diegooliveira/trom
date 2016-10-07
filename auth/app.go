package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AppAuth struct {
}

func App() *AppAuth {

	return new(AppAuth)
}

func (a *AppAuth) PreHandle(rw http.ResponseWriter, req *http.Request) bool {

	req.Header.Del("X-User-Id")

	email := req.URL.Query().Get("email")
	token := req.URL.Query().Get("token")

	resp, errGet := http.Get("http://sample.com.br/user?email=" + email + "&token=" + token)
	if errGet != nil {
		return false
	}

	defer resp.Body.Close()
	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {

	}

	appResp := new(appResponse)
	if err := json.Unmarshal(body, &appResp); err != nil {

		return false
	}

	return true
}

type appResponse struct {
}
