package logout

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

var ErrInvalidRequest = errors.New("invalid request")

func Logout(username string) error {
	var client http.Client
	data := []byte(`{"username":"` + username + `"}`)
	reqBody := bytes.NewReader(data)
	request, err := http.NewRequest(http.MethodPost, "http://micrologin:8080/updateLog", reqBody)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if response.StatusCode != 202 {
		fmt.Println(response)
		return ErrInvalidRequest
	}

	return nil
}
