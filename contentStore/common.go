package contentStore

import (
	"io/ioutil"
	"net/http"
)

/**
 * this assumes no auth is needed to get data from url,
 * if aws s3 requires auth, aws sdk is required
 */
func GetObjectFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}