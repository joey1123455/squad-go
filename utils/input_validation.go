package utils

import (
	"net/mail"
	"net/url"
	"regexp"
	"time"
)

func IsValidURL(input string) bool {
	parsedURL, err := url.ParseRequestURI(input)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}

func IsValidNigerianPhoneNumber(number string) bool {
	// Define the regex pattern
	pattern := `^(0)[789][01]\d{8}$`

	// Compile the regex
	regex := regexp.MustCompile(pattern)

	// Match the string against the regex
	return regex.MatchString(number)
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidNigerianBVN(bvn string) bool {
	// Define the regex pattern for BVN
	pattern := `^\d{11}$`

	// Compile the regex
	regex := regexp.MustCompile(pattern)

	// Match the string against the regex
	return regex.MatchString(bvn)
}

func IsValidNigerianAccountNumber(accountNumber string) bool {
	// Define the regex pattern for a generic Nigerian bank account number
	pattern := `^\d{10}$`

	// Compile the regex
	regex := regexp.MustCompile(pattern)

	// Match the string against the regex
	return regex.MatchString(accountNumber)
}

func IsValidDateOfBirth(dateOfBirth string) bool {
	// Define the expected date layout
	layout := "01/02/2006"

	// Parse the date string
	parsedDate, err := time.Parse(layout, dateOfBirth)

	// Check for parsing errors and validate the date
	return err == nil && parsedDate.Before(time.Now())
}

func IsValidSesId(bvn string) bool {
	// Define the regex pattern for BVN
	pattern := `^\d{30}$`

	// Compile the regex
	regex := regexp.MustCompile(pattern)

	// Match the string against the regex
	return regex.MatchString(bvn)
}
