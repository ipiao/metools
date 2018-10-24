package thirdapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// 高德key
const (
	GdAk = "d9f366aab6d6c19c459f52fab1643e48"
)

// GDMapAddressToCoord 高德
type GDMapAddressToCoord struct {
	Status   string
	Count    string
	Info     string
	Geocodes []struct {
		Location string
		// FormattedAddress string
		// Province         string
		// City             string
		// CityCode         string
		// District         string
		// Township         []string
		// Street           []string
		// Number           []string
		// Adcode           string
		// Level            string
	}
}

// GdMapAddressToCoordJSON 主工具函数,返回凭借原始数据
func GdMapAddressToCoordJSON(address string) (string, error) {
	var baseURL = "http://restapi.amap.com/v3/geocode/geo?"
	var values = url.Values{}
	values.Add("output", "JSON")
	values.Add("key", GdAk)
	values.Add("address", address)
	httpURL := baseURL + values.Encode()
	cilent := http.Client{Timeout: time.Millisecond * 300}
	req, err := http.NewRequest("GET", httpURL, nil)
	if err != nil {
		return "", errors.New("bad request")
	}
	resp, err := cilent.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	return string(bs), err
}

// GdMapAddressToCoordStruct 主工具函数,返回凭借原始数据
func GdMapAddressToCoordStruct(address string) (*GDMapAddressToCoord, error) {
	var res = new(GDMapAddressToCoord)
	str, err := GdMapAddressToCoordJSON(address)
	if err != nil {
		log.Println("http调用错误：", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(str), res)
	if err != nil {
		log.Println("json解析错误", err)
		return nil, err
	}
	return res, nil
}
