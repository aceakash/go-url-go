package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"reflect"
	"time"
)

func TestFetchsOneUrl(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "bla")
	}))

	resp, _ := fetchUrl(ts.URL)
	if resp != "bla" {
		t.Error("Did not get expected response", resp)
	}
}

func TestFetchsManyUrls(t *testing.T) {

	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response for first")
	}))
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response for second")
	}))
	ts3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response for third")
	}))

	expectedResponse := []urlAndRespBody{
		urlAndRespBody{
			url: ts1.URL,
			respBody: "Response for first",
		},
		urlAndRespBody{
			url: ts2.URL,
			respBody: "Response for second",
		},
		urlAndRespBody{
			url: ts3.URL,
			respBody: "Response for third",
		},
	}

		actualResponse, _ := fetchUrls([]string{ts1.URL, ts2.URL, ts3.URL})


	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Error("Expected to get", expectedResponse, "but got", actualResponse)
	}
}


func TestConcatenatingUrls(t *testing.T) {

	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response for first")
	}))
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response for second")
	}))
	ts3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Response for third")
	}))

	expectedResponse := "Response for firstResponse for secondResponse for third"

	actualResponse, _ := Concatenator(ts1.URL, ts2.URL, ts3.URL)


	if actualResponse != expectedResponse {
		t.Error("Expected to get", expectedResponse, "but got", actualResponse)
	}
}

func BenchmarkConcatenator(b *testing.B) {

	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 200)
		fmt.Fprint(w, "Response for first")
	}))
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 200)
		fmt.Fprint(w, "Response for second")
	}))
	ts3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 200)
		fmt.Fprint(w, "Response for third")
	}))



	for i := 0; i < b.N; i++ {
		Concatenator(ts1.URL, ts2.URL, ts3.URL)
	}
}
