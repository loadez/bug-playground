package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// --- Passing tests ---

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var resp map[string]string
	json.NewDecoder(w.Body).Decode(&resp)
	if resp["status"] != "ok" {
		t.Errorf("expected ok, got %s", resp["status"])
	}
}

func TestItemsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/items", nil)
	w := httptest.NewRecorder()
	itemsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var items []Item
	json.NewDecoder(w.Body).Decode(&items)
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
}

func TestItemHandler_Found(t *testing.T) {
	req := httptest.NewRequest("GET", "/items/1", nil)
	w := httptest.NewRecorder()
	itemHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestItemHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/items/999", nil)
	w := httptest.NewRecorder()
	itemHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

func TestItemHandler_InvalidID(t *testing.T) {
	req := httptest.NewRequest("GET", "/items/abc", nil)
	w := httptest.NewRecorder()
	itemHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestFormatPrice(t *testing.T) {
	tests := []struct {
		price    float64
		expected string
	}{
		{9.99, "9.99"},
		{0, "0.00"},
		{100, "100.00"},
		{19.999, "20.00"},
	}
	for _, tc := range tests {
		got := FormatPrice(tc.price)
		if got != tc.expected {
			t.Errorf("FormatPrice(%.3f) = %s, want %s", tc.price, got, tc.expected)
		}
	}
}

func TestValidatePrice(t *testing.T) {
	if err := ValidatePrice(10); err != nil {
		t.Errorf("expected no error for positive price, got %v", err)
	}
	if err := ValidatePrice(0); err != nil {
		t.Errorf("expected no error for zero price, got %v", err)
	}
	if err := ValidatePrice(-5); err == nil {
		t.Error("expected error for negative price")
	}
}

// --- Failing test: catches the discount bug ---

func TestApplyDiscount(t *testing.T) {
	// 20% off $100 should be $80
	got := ApplyDiscount(100, 20)
	if got != 80 {
		t.Errorf("ApplyDiscount(100, 20) = %.2f, want 80.00", got)
	}
}

func TestApplyDiscount_50Percent(t *testing.T) {
	got := ApplyDiscount(200, 50)
	if got != 100 {
		t.Errorf("ApplyDiscount(200, 50) = %.2f, want 100.00", got)
	}
}

// --- Failing test: catches the status code bug ---

func TestTimeHandler_CreateReturns201(t *testing.T) {
	req := httptest.NewRequest("GET", "/time?create=true", nil)
	w := httptest.NewRecorder()
	timeHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}

// --- Flaky test: depends on timing ---

func TestFlakyTiming(t *testing.T) {
	start := time.Now()
	// Simulate work that sometimes takes too long
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	elapsed := time.Since(start)

	if elapsed > 30*time.Millisecond {
		t.Errorf("too slow: %v (threshold 30ms)", elapsed)
	}
}

// --- Flaky test: random failure ---

func TestFlakyRandom(t *testing.T) {
	// Fails ~20% of the time
	if rand.Intn(5) == 0 {
		t.Error("random failure — this test is flaky")
	}
}
