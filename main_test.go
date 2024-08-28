package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strconv"
	"testing"
)

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