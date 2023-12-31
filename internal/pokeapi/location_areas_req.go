package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl *string) (LocationAreasResp, string, error) {
	const endpoint = "/location-area/?offset=0&limit=20"
	fullUrl := baseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	// Check if url is in cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache Hit!")
		respJson := LocationAreasResp{}

		err := json.Unmarshal(dat, &respJson)
		if err != nil {
			return LocationAreasResp{}, fullUrl, fmt.Errorf("error unmarshaling json: %s", err)
		}

		return respJson, fullUrl, nil
	}

	fmt.Println("Cache Missed")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, fullUrl, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, fullUrl, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fullUrl, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, fullUrl, err
	}

	respJson := LocationAreasResp{}

	err = json.Unmarshal(dat, &respJson)
	if err != nil {
		return LocationAreasResp{}, fullUrl, fmt.Errorf("error unmarshaling json: %s", err)
	}

	c.cache.Add(fullUrl, dat)

	return respJson, fullUrl, nil
}
