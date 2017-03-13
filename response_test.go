package respond_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	resp "github.com/nicklaw5/go-respond"
)

func TestIsJSON(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w).IsJSON(true)
		res.Created(nil)
	})
	handler.ServeHTTP(rr, req)

	expected := "application/json; charset=utf-8"
	if rr.Header().Get("Content-Type") != expected {
		t.Errorf("Handler returned unexpected header: got %v wanted %v",
			rr.Header().Get("Content-Type"), expected)
	}
}

func TestSetHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		headers := map[string]string{
			"X-One-1": "headers-one",
			"X-One-2": "headers-one",
		}
		res.SetHeaders(headers)

		headers = map[string]string{
			"X-Two-1": "headers-two",
			"X-Two-2": "headers-two",
		}
		res.SetHeaders(headers)
		res.Created(nil)
	})
	handler.ServeHTTP(rr, req)

	expected := "headers-two"
	if rr.Header().Get("X-Two-1") != expected {
		t.Errorf("Handler returned unexpected header: got %v wanted %v",
			rr.Header().Get("X-Two-1"), expected)
	}

	if rr.Header().Get("X-Two-2") != expected {
		t.Errorf("Handler returned unexpected header: got %v wanted %v",
			rr.Header().Get("X-Two-2"), expected)
	}
}

func TestAddHeader(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.AddHeader("Content-Type", "application/json")
		res.Created(nil)
	})
	handler.ServeHTTP(rr, req)

	expected := "application/json"
	if rr.Header().Get("Content-Type") != expected {
		t.Errorf("Handler returned unexpected header: got %v wanted %v",
			rr.Header().Get("Content-Type"), expected)
	}
}

func TestDeleteHeader(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.AddHeader("Content-Type", "application/json")
		res.DeleteHeader("Content-Type")
		res.Created(nil)
	})
	handler.ServeHTTP(rr, req)

	expected := ""
	if rr.Header().Get("Content-Type") != expected {
		t.Errorf("Handler returned unexpected header: got %v wanted %v",
			rr.Header().Get("Content-Type"), expected)
	}
}