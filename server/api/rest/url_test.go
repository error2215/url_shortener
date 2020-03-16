package rest

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestCreateURLHandler(t *testing.T) {
	r, err := recorder.New("../../../cassettes/create_url")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	client := &http.Client{
		Transport: r,
	}

	urlStr := "http://localhost:3034/api/v1/url_shortener"
	val := url.Values{}
	val.Add("url", "google.com")
	resp, err := client.PostForm(urlStr, val)
	if err != nil {
		t.Fatalf("Failed to get url %s: %s", urlStr, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}

	wantHeading := `"data":{"alias"`
	bodyContent := string(body)

	if !strings.Contains(bodyContent, wantHeading) {
		t.Errorf("Heading %s not found in response", wantHeading)
	}
}

func TestCreateURLHandlerNoUrlValues(t *testing.T) {
	r, err := recorder.New("../../../cassettes/create_url_no_values")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	client := &http.Client{
		Transport: r,
	}

	urlStr := "http://localhost:3034/api/v1/url_shortener"
	val := url.Values{}
	resp, err := client.PostForm(urlStr, val)
	if err != nil {
		t.Fatalf("Failed to get url %s: %s", urlStr, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}

	wantHeading := `"message":"Missing required params"`
	bodyContent := string(body)

	if !strings.Contains(bodyContent, wantHeading) {
		t.Errorf("Heading %s not found in response", wantHeading)
	}
}

func TestRedirectHandler(t *testing.T) {
	r, err := recorder.New("../../../cassettes/redirect")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	client := &http.Client{
		Transport: r,
	}

	urlStr := "http://localhost:3034/rd/5e6e67a2accd204c905bedbd"
	resp, err := client.Get(urlStr)
	if err != nil {
		t.Fatalf("Failed to get url %s: %s", urlStr, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}
	if len(body) == 0 {
		t.Errorf("Response body length is 0")
	}

}

func TestRedirectHandlerBadLink(t *testing.T) {
	r, err := recorder.New("../../../cassettes/redirect_bad_link")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	client := &http.Client{
		Transport: r,
	}

	urlStr := "http://localhost:3034/rd/asd"
	resp, err := client.Get(urlStr)
	if err != nil {
		t.Fatalf("Failed to get url %s: %s", urlStr, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}
	wantHeading := `"success":false`
	bodyContent := string(body)

	if !strings.Contains(bodyContent, wantHeading) {
		t.Errorf("Heading %s not found in response", wantHeading)
	}
}
