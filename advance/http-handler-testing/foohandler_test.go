package foohandler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go run test execute go test -v ./...

// method 2
func TestHandleGetFooRR(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}
	handleGetFoo(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but go %d", rr.Result().StatusCode)
	}
	defer rr.Result().Body.Close()

	expected := "Foo"
	b, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but we got %s ", expected, string(b))
	}

}

// method 1
func TestHandleGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but go %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expected := "Foo"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but we got %s ", expected, string(b))
	}

}
