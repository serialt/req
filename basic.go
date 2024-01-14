package req

import "encoding/base64"

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func BasicAuthHeaderValue(username, password string) string {
	return "Basic " + basicAuth(username, password)
}
