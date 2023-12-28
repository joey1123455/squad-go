package utils

import "testing"

func TestIsValidURL(t *testing.T) {
	testCases := []struct {
		url      string
		expected bool
	}{
		{"https://www.example.com", true},
		{"http://localhost:8080/path", true},
		{"ftp://ftp.example.com", true},
		{"invalid-url", false},
		{"not_a_scheme://example.com", false},
	}

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			result := IsValidURL(tc.url)
			if result != tc.expected {
				t.Errorf("Expected isValidURL(%s) to be %t, but got %t", tc.url, tc.expected, result)
			}
		})
	}
}
