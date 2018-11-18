package apistuff

import "net/http"

// GetCookieByName scans the cookies and returns one with a matching key, if
// there is one.
func GetCookieByName(cookies []*http.Cookie, name string) string {
	var result string
	for i := 0; i < len(cookies); i++ {
		if cookies[i].Name == name {
			result = cookies[i].Value
		}
	}
	return result
}
