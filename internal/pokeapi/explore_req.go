package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAreaResp(pageNum *string) (AreaResp, error) {
	fullUrl := baseUrl + "/location-area/" + *pageNum

	// Check if url is in cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache Hit!")
		respJson := AreaResp{}

		err := json.Unmarshal(dat, &respJson)
		if err != nil {
			return AreaResp{}, fmt.Errorf("error unmarshaling json: %s", err)
		}

		return respJson, nil
	}

	fmt.Println("Cache Missed")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return AreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return AreaResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return AreaResp{}, err
	}

	respJson := AreaResp{}

	err = json.Unmarshal(dat, &respJson)
	if err != nil {
		return AreaResp{}, fmt.Errorf("error unmarshaling json: %s", err)
	}

	c.cache.Add(fullUrl, dat)

	return respJson, nil
}
