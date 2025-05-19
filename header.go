package go_payoucard

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/jason",
	}
}

// notice:  input params is fixed
func getAuthHeaders(auth string) map[string]string {
	return map[string]string{
		"Content-Type":      "application/jason",
		"charset":           "utf-8",
		"Gt-Authentication": auth,
	}
}
