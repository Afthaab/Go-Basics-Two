package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Hobbies struct {
	Userid   string
	LHobbies []string
}

func HomeFunction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusBadRequest)

	h := Hobbies{
		Userid: "101",
		LHobbies: []string{
			"eating",
			"sleeping",
			"playing",
		},
	}
	data, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
		return
	}
	// w.Write([]byte("Hello this is home function"))
	w.Write(data)
}
