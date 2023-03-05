package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/4IDTest/SingIn/history"
	"github.com/4IDTest/SingIn/user"
	"github.com/4IDTest/SingIn/utils"
)

func AuthHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var userData user.User
		err = json.Unmarshal(data, &userData)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		if userData.Username == "" || userData.Password == "" {
			http.Error(rw, "Missing credentials", http.StatusBadRequest)
			return
		}
		err = user.AuthUser(userData.Username, userData.Password)
		if err != nil {
			if errors.Is(err, utils.ErrInvalidCredentials) {
				fmt.Println(utils.ErrInvalidCredentials.Error())
				http.Error(rw, "Invalid Credentials", http.StatusUnauthorized)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}
}

func UpdateLogRecordHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		input := make(map[string]any)
		err = json.Unmarshal(data, &input)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = history.UpdateLogRecord(input["username"].(string))
		if err != nil {
			if errors.Is(err, utils.ErrNotLogged) {
				fmt.Println(utils.ErrNotLogged.Error())
				rw.WriteHeader(http.StatusNotFound)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		rw.WriteHeader(http.StatusAccepted)
	}
}

func RegisterUserHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		input := make(map[string]any)
		err = json.Unmarshal(data, &input)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = user.RegisterUser(input["username"].(string), input["password"].(string))
		if err != nil {
			if errors.Is(err, utils.ErrAlreadyExist) {
				fmt.Println(utils.ErrAlreadyExist.Error())
				rw.WriteHeader(http.StatusConflict)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		rw.WriteHeader(http.StatusAccepted)
	}
}
