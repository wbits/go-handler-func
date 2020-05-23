package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func Routes() {
	http.HandleFunc("/users", SendJson)
	http.HandleFunc("/sendUsers", ReceiveJson)
}

func SendJson(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name, Email string
	} {
		Name: "Dick",
		Email: "dickbrouwers@chello.nl",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

func ReceiveJson(rw http.ResponseWriter, r *http.Request) {
	var u struct {
		Name string `json:"Name"`
		Email string `json:"Email"`
	}

	f, err := os.OpenFile("application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error create file: %v", err)
	}
	defer f.Close()

	l := log.New(f, "INFO: ", log.Llongfile)

	if r.Method != http.MethodPost {
		log.Fatal("receive json only accepts post requests")
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	rw.WriteHeader(204)
	l.Printf("%v", u)
}
