package handlertest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BogdanMelchenko/gorest-task/application/util"
)

func handleMock(w http.ResponseWriter, r *http.Request) {
	util.RespondWithoutError(w, http.StatusOK, `{"nothin":12}`, r.Header.Get("Content-Type"))
}

func errorMock(w http.ResponseWriter, r *http.Request) {
	util.RespondWithError(w, http.StatusBadRequest, "test error message")
}

func TestXMLResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/xml")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleMock)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := "application/xml"

	if rr.Header().Get("Content-Type") != expected {
		t.Errorf("handler returned unexpected content-type: got %v wanted %v",
			rr.Header().Get("Content-Type"), expected)
	}
}

func TestJSONResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleMock)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := "application/json"

	if rr.Header().Get("Content-Type") != expected {
		t.Errorf("handler returned unexpected content-type: got %v wanted %v",
			rr.Header().Get("Content-Type"), expected)
	}
}

func TestErrorResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(errorMock)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusBadGateway)
	}

	expected := `{"error":"test error message"}`

	if actual, _ := rr.Body.ReadString(';'); actual != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v",
			actual, expected)
	}
}
