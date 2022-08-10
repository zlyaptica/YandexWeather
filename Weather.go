package YandexWeather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client HTTP-client struct
type Client struct {
	client *http.Client
	apiKey string
}

// NewClient Create new client. Apikey - apikey from https://developer.tech.yandex.ru/services/
func NewClient(apikey string) *Client {
	return &Client{
		client: &http.Client{},
		apiKey: apikey,
	}
}

// GetWeatherInfo Get yandex weather method
func (c Client) GetWeatherInfo(client *Client, lat, lon float64) (Response, error) {
	requestURL := fmt.Sprintf("https://api.weather.yandex.ru/v2/informers?lat=%f&lon=%f&lang=ru_RU", lat, lon)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("X-Yandex-API-Key", client.apiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return Response{}, err
	}
	return response, nil
}
