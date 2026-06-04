package toolbox

import "net/url"

func IsValidURL(urlToCheck string) bool {
	u, err := url.Parse(urlToCheck)
	if err != nil {
		return false
	}
	return u.Scheme != "" && u.Host != ""
}
