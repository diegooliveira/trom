package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserAuth struct {
}

func User() *UserAuth {

	return new(UserAuth)
}

func (a *UserAuth) PreHandle(rw http.ResponseWriter, req *http.Request) bool {

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

	userResp := new(userResponse)
	if err := json.Unmarshal(body, &userResp); err != nil {

		return false
	}

	return true
}

type userResponse struct {
}
