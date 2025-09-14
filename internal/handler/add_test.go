package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"plus-number/internal/model"
)

// Mock DB
type mockDB struct {
	result int
	err    error
}

func (m *mockDB) AddNumbers(ctx context.Context, a, b int) (int, error) {
	return m.result, m.err
}

func TestAddHandler_Success(t *testing.T) {
	h := &AddHandler{DB: &mockDB{result: 42}}

	body := []byte(`{"a": 20, "b": 22}`)
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp model.AddResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Result != 42 {
		t.Fatalf("expected result=42, got %d", resp.Result)
	}
}

func TestAddHandler_InvalidJSON(t *testing.T) {
	h := &AddHandler{DB: &mockDB{result: 0}}

	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewReader([]byte(`invalid`)))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestAddHandler_MethodNotAllowed(t *testing.T) {
	h := &AddHandler{DB: &mockDB{result: 0}}

	req := httptest.NewRequest(http.MethodGet, "/api/add", nil)

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rr.Code)
	}
}
