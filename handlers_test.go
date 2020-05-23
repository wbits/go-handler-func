package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
    "os"
    "testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
    Routes()
}

// TestSendJSON testing the users internal endpoint.
func TestSendJSON(t *testing.T) {
    t.Log("Given the need to test the SendJSON endpoint.")
    {
        req, err := http.NewRequest("GET", "/users", nil)
        if err != nil {
            t.Fatal("\tShould be able to create a request.",
                ballotX, err)
        }
        t.Log("\tShould be able to create a request.",
            checkMark)

        rw := httptest.NewRecorder()
        http.DefaultServeMux.ServeHTTP(rw, req)

        if rw.Code != 200 {
            t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
        }
        t.Log("\tShould receive \"200\"", checkMark)


        u := struct {
            Name  string
            Email string
        }{}

        if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
            t.Fatal("\tShould decode the response.", ballotX)
        }
        t.Log("\tShould decode the response.", checkMark)

        if u.Name == "Dick" {
          t.Log("\tShould have a Name.", checkMark)
        } else {
          t.Error("\tShould have a Name.", ballotX, u.Name)
        }

        if u.Email == "dickbrouwers@chello.nl" {
            t.Log("\tShould have an Email.", checkMark)
        } else {
            t.Error("\tShould have an Email.", ballotX, u.Email)
        }
    }
}

func TestReceiveJSON(t *testing.T) {
    t.Log("Given the need to test the ReceiveJSON endpoint.")
    {
        jr, err := os.Open("../user.json")
        if err != nil {
            t.Fatal("\tCould not load the json file.",
                ballotX, err)
        }
        req, err := http.NewRequest("POST", "/sendUsers", jr)
        if err != nil {
            t.Fatal("\tShould be able to create a request.",
                ballotX, err)
        }
        t.Log("\tShould be able to create a request.",
            checkMark)

        rw := httptest.NewRecorder()
        http.DefaultServeMux.ServeHTTP(rw, req)

        if rw.Code != 204 {
            t.Fatal("\tShould receive \"204\"", ballotX, rw.Code)
        }
        t.Log("\tShould receive \"204\"", checkMark)
    }
}
