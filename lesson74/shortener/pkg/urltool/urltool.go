package urltool

import (
	"net/url"
	"path"
)

func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}
	return path.Base(myUrl.Path), nil
}
