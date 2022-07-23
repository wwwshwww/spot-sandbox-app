package yahoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ReverseGeoCodeEndPoint = "https://map.yahooapis.jp/geoapi/V1/reverseGeoCoder"
)

type ReverseGeoCodeResponse struct {
	ResultInfo struct {
		Count        int     `json:"Count"`
		Total        int     `json:"Total"`
		Start        int     `json:"Start"`
		Latency      float64 `json:"Latency"`
		Status       int     `json:"Status"`
		Description  string  `json:"Description"`
		Copyright    string  `json:"Copyright"`
		CompressType string  `json:"CompressType"`
	} `json:"ResultInfo"`
	Feature []struct {
		Geometry struct {
			Type        string `json:"Type"`
			Coordinates string `json:"Coordinates"`
		} `json:"Geometry"`
		Property struct {
			Country struct {
				Code string `json:"Code"`
				Name string `json:"Name"`
			} `json:"Country"`
			Address        string `json:"Address"`
			AddressElement []struct {
				Name  string `json:"Name"`
				Kana  string `json:"Kana"`
				Level string `json:"Level"`
				Code  string `json:"Code,omitempty"`
			} `json:"AddressElement"`
			Road []struct {
				Name        string `json:"Name"`
				Kana        string `json:"Kana"`
				PopularName string `json:"PopularName"`
			} `json:"Road"`
		} `json:"Property"`
	} `json:"Feature"`
}

func YahooReverseGeoCode(lat, lng float64) (*ReverseGeoCodeResponse, error) {
	u, err := url.Parse(ReverseGeoCodeEndPoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("lat", strconv.FormatFloat(lat, 'f', 8, 64))
	q.Set("lon", strconv.FormatFloat(lng, 'f', 8, 64))
	q.Set("output", "json")
	q.Set("appid", "APIKEY")
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result ReverseGeoCodeResponse
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return &result, nil
}
