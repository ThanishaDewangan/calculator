package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Person born in 1990",
			dob:      time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Now().Year() - 1990,
		},
		{
			name:     "Person born this year",
			dob:      time.Now().AddDate(0, 0, -1),
			expected: 0,
		},
		{
			name:     "Person born exactly one year ago",
			dob:      time.Now().AddDate(-1, 0, 0),
			expected: 1,
		},
		{
			name:     "Person born 25 years ago",
			dob:      time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Now().Year() - 1999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Adjust expected age based on whether birthday has occurred this year
			now := time.Now()
			expected := now.Year() - tt.dob.Year()
			nowMonth, nowDay := now.Month(), now.Day()
			dobMonth, dobDay := tt.dob.Month(), tt.dob.Day()
			if nowMonth < dobMonth || (nowMonth == dobMonth && nowDay < dobDay) {
				expected--
			}

			result := CalculateAge(tt.dob)
			if result != expected {
				t.Errorf("CalculateAge() = %v, want %v", result, expected)
			}
		})
	}
}

func TestCalculateAgeEdgeCases(t *testing.T) {
	now := time.Now()
	
	// Test birthday today
	dobToday := time.Date(now.Year()-25, now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	age := CalculateAge(dobToday)
	if age != 25 {
		t.Errorf("CalculateAge() for birthday today = %v, want 25", age)
	}

	// Test birthday tomorrow (handle month boundaries)
	nextDay := now.AddDate(0, 0, 1)
	dobTomorrow := time.Date(nextDay.Year()-25, nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, time.UTC)
	age = CalculateAge(dobTomorrow)
	if age != 24 {
		t.Errorf("CalculateAge() for birthday tomorrow = %v, want 24", age)
	}

	// Test birthday yesterday
	prevDay := now.AddDate(0, 0, -1)
	dobYesterday := time.Date(prevDay.Year()-25, prevDay.Month(), prevDay.Day(), 0, 0, 0, 0, time.UTC)
	age = CalculateAge(dobYesterday)
	if age != 25 {
		t.Errorf("CalculateAge() for birthday yesterday = %v, want 25", age)
	}
}


