package format

import (
	URL "net/url"
)

func UrlDecode(uri string) (string, error) {
	u, err := URL.Parse(uri)
	if err != nil {
		return "", err
	}
	return u.Path, err
}
