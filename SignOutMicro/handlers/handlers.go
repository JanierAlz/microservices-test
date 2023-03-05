package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/4IDTest/SingOut/logout"
)

func LogoutHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userData := make(map[string]any)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		err = json.Unmarshal(data, &userData)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		if userData["username"].(string) == "" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		err = logout.Logout(userData["username"].(string))
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
		}
	}
}
