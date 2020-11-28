package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Peace-for-ever/peace-api/models"
	"github.com/stretchr/testify/assert"
)

func TestGetPersons(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/persons", nil)

	q := req.URL.Query()
	q.Add("maxY", "0")
	q.Add("maxX", "0")
	q.Add("maxPersons", "1")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)

	var persons []models.Person
	err := json.Unmarshal(w.Body.Bytes(), &persons)
	if err != nil {
		t.Error("Unable to unmarshal response")
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, IfThenElse(len(persons) == 1, true, false), true, "Max person is incorect")
	assert.Equal(t, models.Location{X: 0, Y: 0}, persons[0].Location, req.URL.String())
	assert.Equal(t, IfThenElse(persons[0].LastName != "" && persons[0].FirstName != "", true, false), true, "Max person is incorect")
}

func TestGetPersonsFlood(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/persons", nil)

	q := req.URL.Query()
	q.Add("maxY", "0")
	q.Add("maxX", "0")
	q.Add("maxPersons", "900000")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	expected := json.RawMessage(`{"message":"reste calme"}`)
	expectedBytes, _ := json.Marshal(expected)

	assert.Equal(t, 406, w.Code)
	assert.Equal(t, string(expectedBytes), w.Body.String())
}

func TestGetSentence(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sentence", nil)
	router.ServeHTTP(w, req)
	var sentence string
	err := json.Unmarshal(w.Body.Bytes(), &sentence)
	if err != nil {
		t.Error("Unable to unmarshal response")
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, IfThenElse(sentence != "", true, false), true, "Sentence is empty")
}

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
