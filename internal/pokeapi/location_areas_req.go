package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) getLocationAreas() (LocationAreasResp, error) {
	const fullUrl = baseUrl + "/location-area"
	req, err := http.NewRequest("Get", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	respJson := LocationAreasResp{}

	err = json.Unmarshal(dat, &respJson)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("error unmarshaling json: %s", err)
	}

	return respJson, nil
}
