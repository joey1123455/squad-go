package utils

import "net/url"

func IsValidURL(input string) bool {
	parsedURL, err := url.ParseRequestURI(input)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}
