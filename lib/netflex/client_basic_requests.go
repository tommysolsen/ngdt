package netflex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c Client) Get(url string, payload interface{}) error {
	request, err := http.NewRequest("GET", strings.Join([]string{c.baseUrl, url}, "/"), nil)
	if err != nil {
		return fmt.Errorf("Unable to create request, %s", err)
	}
	request.SetBasicAuth(c.ApiData.PublicKey, c.ApiData.PrivateKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := c.client.Do(request)
	if err != nil {
		return fmt.Errorf("Unable to perform request: %s", err)
	}

	if response.StatusCode < 200 || response.StatusCode > 399 {
		return fmt.Errorf("http request failed: Got statuscode %d (%s)", response.StatusCode, response.Status)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	err = json.Unmarshal(buf.Bytes(), payload)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal, %s", err)
	}
	return nil
}

func (c Client) Post(url string, body interface{}, payload interface{}) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("post", strings.Join([]string{c.baseUrl, url}, "/"), bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(c.ApiData.PublicKey, c.ApiData.PrivateKey)

	response, err := c.client.Do(request)
	if err != nil {
		return fmt.Errorf("Unable to perform request: %s", err)
	}

	if response.StatusCode < 200 || response.StatusCode > 399 {
		return fmt.Errorf("http request failed: Got statuscode %d (%s)", response.StatusCode, response.Status)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	err = json.Unmarshal(buf.Bytes(), &payload)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal, %s", err)
	}
	return nil
}
