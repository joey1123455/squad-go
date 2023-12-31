package utils

import (
	"testing"
)

func Test_IsValidURL(t *testing.T) {
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

func Test_IsValidNigerianPhoneNumber(t *testing.T) {
	testCases := []struct {
		number        string
		expectedValid bool
	}{
		{"+2347012345678", false},
		{"08123456789", true},
		{"07012345678", true},
		{"0123456789", false},
		{"0801234567", false},
		{"+2340123456789", false},
		{"+234701234567890", false},
		{"+23470123456", false},
		{"+234 7012345678", false},
		{"0812-345-6789", false},
		{"+234+7012345678", false},
		{"+234+70123a45678", false},
		{"+2347a012345678", false},
		{"2347012345678", false},
		{"0092347012345678", false},
		{"08123456789\n", false},
		{"08123456789\r\n", false},
	}

	for _, tc := range testCases {
		t.Run(tc.number, func(t *testing.T) {
			result := IsValidNigerianPhoneNumber(tc.number)
			if result != tc.expectedValid {
				t.Errorf("Expected IsValidNigerianPhoneNumber(%s) to be %t, but got %t", tc.number, tc.expectedValid, result)
			}
		})
	}
}

func Test_IsValidEmail(t *testing.T) {
	testCases := []struct {
		email string
		value bool
	}{
		{"john.doe@example.com", true},
		{"alice_smith123@gmail.com", true},
		{"invalid-email", false},
		{"missing_at_symbol.com", false},
		{"no@tld", true},
		{"special!chars@example.com", true},
		{"uppercase@CASE.com", true},
		{"john.doe@example..com", false},
		{"@example.com", false},
		{"john.doe@.com", false},
		{"john.doe@com", true},
		{"john.doe@.com", false},
	}
	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {
			result := IsValidEmail(tc.email)
			if result != tc.value {
				t.Errorf("Expected IsValidEmail(%s) to be %t, but got %t", tc.email, tc.value, result)
			}
		})
	}
}

func Test_IsValidNigerianBVN(t *testing.T) {
	testCases := []struct {
		bvn            string
		expectedResult bool
	}{
		{"12345678901", true},
		{"123456789", false},
		{"1234567890a", false},
		{"1234567890", false},
		{"98765432109", true},
		{"9876543210", false},
		{"9876543210a", false},
		{"9876543210", false},
	}

	for _, tc := range testCases {
		t.Run(tc.bvn, func(t *testing.T) {
			result := IsValidNigerianBVN(tc.bvn)
			if result != tc.expectedResult {
				t.Errorf("Expected isValidNigerianBVN(%s) to be %t, but got %t", tc.bvn, tc.expectedResult, result)
			}
		})
	}
}

func Test_IsValidNigerianAccountNumber(t *testing.T) {
	testCases := []struct {
		accountNumber  string
		expectedResult bool
	}{
		{"1234567890", true},   // Valid
		{"9876543210", true},   // Valid
		{"ABC123456", false},   // Invalid: Contains non-digit character
		{"123-456-789", false}, // Invalid: Contains non-digit character
		{"", false},            // Invalid: Empty string
	}

	for _, tc := range testCases {
		t.Run(tc.accountNumber, func(t *testing.T) {
			result := IsValidNigerianAccountNumber(tc.accountNumber)
			if result != tc.expectedResult {
				t.Errorf("Expected isValidNigerianAccountNumber(%s) to be %t, but got %t", tc.accountNumber, tc.expectedResult, result)
			}
		})
	}
}

func Test_IsValidDateOfBirth(t *testing.T) {
	testCases := []struct {
		dateOfBirth    string
		expectedResult bool
	}{
		{"05/15/1990", true},    // Valid
		{"12/31/2005", true},    // Valid
		{"02/30/1980", false},   // Invalid: February 30th doesn't exist
		{"invalid-date", false}, // Invalid: Not a valid date format
		{"01/01/2025", false},   // Invalid: Future date
		{"13/25/1990", false},   // Invalid: Month greater than 12
		{"05/32/1990", false},   // Invalid: Day greater than 31
		{"00/15/1990", false},   // Invalid: Month zero
		{"05/00/1990", false},   // Invalid: Day zero
	}

	for _, tc := range testCases {
		t.Run(tc.dateOfBirth, func(t *testing.T) {
			result := IsValidDateOfBirth(tc.dateOfBirth)
			if result != tc.expectedResult {
				t.Errorf("Expected isValidDateOfBirth(%s) to be %t, but got %t", tc.dateOfBirth, tc.expectedResult, result)
			}
		})
	}
}

func Test_IsValidSesId(t *testing.T) {
	testCases := []struct {
		bvn            string
		expectedResult bool
	}{
		{"891234567890123456789012345678", true},
		{"2451879320654187032918475032198745032", false},
		{"K7bFg2pXqL4sU1rM3wZ5yD8vA9cO6hJ0", false},
		{"4819260357142098563725189402", false},
		{"543210987654321098765432109876", true},
		{"9876543210987654321098765", false},
		{"9876543210a", false},
		{"8923456712034567890123456789012", false},
	}

	for _, tc := range testCases {
		t.Run(tc.bvn, func(t *testing.T) {
			result := IsValidSesId(tc.bvn)
			if result != tc.expectedResult {
				t.Errorf("Expected isValidSesId(%s) to be %t, but got %t", tc.bvn, tc.expectedResult, result)
			}
		})
	}
}
