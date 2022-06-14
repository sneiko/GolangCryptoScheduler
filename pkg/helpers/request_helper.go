package helpers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetBodyFromResp(resp *http.Response) ([]byte, error) {
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body, nil
}
