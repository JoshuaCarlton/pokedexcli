package main

import (
	"fmt"
	"io"
	"net/http"
)

func get_bytes(config *config, url string) ([]byte, error) {
	bytes, ok := config.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
		}
		newBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		config.cache.Add(url, newBytes)
		bytes = newBytes
	}
	return bytes, nil
}
