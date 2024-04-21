package pkg

import (
	"io"
	"net/http"
)

func Get(url string) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, _ := io.ReadAll(resp.Body)

	return string(result), nil
}
