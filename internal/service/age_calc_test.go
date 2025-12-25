package service

import (
	"testing"
	"time"
)

func mustDate(t *testing.T, s string) time.Time {
	t.Helper()
	d, err := time.Parse("2006-01-02", s)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return d
}

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name string
		dob  string
		now  string
		want int
	}{
		{"exact birthday", "1990-05-10", "2025-05-10", 35},
		{"before birthday", "1990-12-31", "2025-01-01", 34},
		{"after birthday", "1990-01-01", "2025-12-31", 35},
		{"future dob", "2030-01-01", "2025-01-01", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dob := mustDate(t, tt.dob)
			now := mustDate(t, tt.now)
			got := CalculateAge(dob, now)
			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}


