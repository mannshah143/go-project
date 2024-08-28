package main

import (
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func additionHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	aStr := values.Get("a")
	bStr := values.Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "Parameter 'a' must be an integer", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "Parameter 'b' must be an integer", http.StatusBadRequest)
		return
	}

	result := Add(a, b)

	fmt.Fprintf(w, "%d", result)
}

func TestAdditionHandler(t *testing.T) {
	// Test case 1
	req1 := httptest.NewRequest("GET", "/addition?a=3&b=5", nil)
	w1 := httptest.NewRecorder()
	additionHandler(w1, req1)
	resp1 := w1.Result()
	defer resp1.Body.Close()
	body1, _ := ioutil.ReadAll(resp1.Body)
	result1, _ := strconv.Atoi(string(body1))
	expected1 := 8
	if result1 != expected1 {
		t.Errorf("Test case 1 failed, expected %d but got %d", expected1, result1)
	}
}

func main() {
	http.HandleFunc("/addition", additionHandler)

	fmt.Println("Server listening on port 8080....")
	http.ListenAndServe(":8080", nil)
}