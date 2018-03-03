package v2ex

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetAPIData(url string, v interface{}) error {
	log.Printf("[D] request url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("[D] resp: %s", resp.Body)
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(v)
	return err
}
