package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project5/models"
	"strconv"
)

var ErrMap = make(map[string]string)

func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userId := r.URL.Query().Get("userid")

	// uid, err := strconv.Atoi(userId)

	uid, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		appErr := map[string]string{"Status": http.StatusText(http.StatusBadRequest), "Message": "Error while converting the string"}
		data, err := json.Marshal(appErr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, http.StatusText(http.StatusBadRequest))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)

		return
	}

	userData, err := models.FetchUser(uid)
	if err != nil {
		appErr := map[string]string{"Status": http.StatusText(http.StatusNotFound), "Error": err.Error()}
		data, err := json.Marshal(appErr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, http.StatusText(http.StatusBadRequest))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)

		return
	}

	data, err := json.Marshal(userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, http.StatusText(http.StatusBadRequest))
		return

	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
