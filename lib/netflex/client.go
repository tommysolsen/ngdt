package netflex

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tommysolsen/ngdt/lib/filetools"
)

type Client struct {
	ApiData          filetools.ApiJSON
	baseUrl          string
	client           http.Client
	templateArray    []Template
	templatesFetched bool
}

func New(data filetools.ApiJSON) Client {
	return Client{
		ApiData: data,
		baseUrl: "https://api.netflexapp.com/v1",
		client:  http.Client{},
	}
}

func (c *Client) LoadTemplates() ([]Template, error) {
	if c.templatesFetched == false && len(c.templateArray) == 0 {
		var tmp []Template
		err := c.Get("foundation/templates", &tmp)
		if err != nil {
			return nil, err
		}
		(*c).templateArray = tmp
		(*c).templatesFetched = true
		return tmp, nil
	}
	return c.templateArray, nil

}

func (c *Client) PostLabel(label string) (Boolean, error) {
	payload := &PostLabelPayload{
		Label: base64.StdEncoding.EncodeToString([]byte(label)),
	}
	payloadString, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("POST", c.baseUrl+"/foundation/labels", bytes.NewBuffer(payloadString))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.ApiData.PublicKey, c.ApiData.PrivateKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 409 {
		return false, fmt.Errorf("Expected error code to be 200 or 409, got: %d", resp.StatusCode)
	}
	return resp.StatusCode == 409, err
}

func (c Client) GetTemplateById(id int64) (*Template, error) {
	templates, err := c.LoadTemplates()
	if err != nil {
		return nil, err
	}
	for _, templ := range templates {
		if id == int64(templ.ID) {
			return &templ, nil
		}
	}
	return nil, nil
}
