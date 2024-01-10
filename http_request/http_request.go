package httprequest

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

func RequestGet(url string, queryParams *url.Values, target interface{}) error {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return err
	}

	if queryParams != nil {
		request.URL.RawQuery = queryParams.Encode()
	}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Failed Fetch Htpp")
	}

	json.NewDecoder(response.Body).Decode(&target)

	return nil
}

func RequestPost(method string, url string, requestBody *map[string]interface{}, target *map[string]interface{}) error {
	var payload *bytes.Buffer = nil

	if requestBody != nil {
		payload = new(bytes.Buffer)
		err := json.NewEncoder(payload).Encode(requestBody)
		if err != nil {
			return err
		}
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest(method, url, payload)
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Failed Fetch Htpp")
	}

	json.NewDecoder(response.Body).Decode(&target)

	return nil
}

func RequestPostForm(method string, url string, multiPartWriter *multipart.Writer, requestBody bytes.Buffer, target *map[string]interface{}) error {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest(method, url, &requestBody)
	request.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	if err != nil {
		return err
	}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Failed Fetch Htpp")
	}

	json.NewDecoder(response.Body).Decode(&target)

	return nil
}
