package utils

func CompleteUrl(endPoint string, live bool) string {
	const (
		testBase = "https://sandbox-api-d.squadco.com/"
		liveBase = "https://api-d.squadco.com/"
	)

	if live {
		return liveBase + endPoint
	}
	return testBase + endPoint
}
