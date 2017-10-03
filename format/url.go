package format

import (
	URL "net/url"
	"errors"
)

func UrlDecode(url string) (string, error) {
	if url == "" {
		return "", errors.New("miss argument url")
	}

	url, err := URL.QueryUnescape(url)
	return url, err
}
