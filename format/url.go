package format

import (
	URL "net/url"
	"errors"
)

func UrlDecode(url string) (string, error) {
	if url == ""{
		return "",errors.New("miss argument url")
	}
	u, err := URL.Parse(url)
	if err != nil {
		return "", err
	}
	return u.Path, err
}
