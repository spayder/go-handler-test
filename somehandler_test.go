package somehandler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetSomeHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetSomeHandler))
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	expected := "Some handler"
	b, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	if string(b) != expected {
		t.Errorf("expected %sbut got %s", expected, string(b))
	}
}

func TestHandleGetSomeHandlerForRR(t *testing.T) {
	// response recorder
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)

	if err != nil {
		t.Error(err)
	}

	handleGetSomeHandler(rr, req)

	result := rr.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", result.StatusCode)
	}

	defer result.Body.Close()

	expected := "Some handler"
	b, err := io.ReadAll(result.Body)

	if err != nil {
		t.Error(err)
	}

	if string(b) != expected {
		t.Errorf("expected %sbut got %s", expected, string(b))
	}
}
