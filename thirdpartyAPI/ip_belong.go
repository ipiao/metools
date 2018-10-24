package thirdapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// AliIP 阿里ip归属地查询
type AliIP struct {
	Status    string `json:"status,omitempty"`    //
	Info      string `json:"info,omitempty"`      //
	Infocode  string `json:"infocode,omitempty"`  //
	Province  string `json:"province,omitempty"`  //
	City      string `json:"city,omitempty"`      //
	AdCode    string `json:"adcode,omitempty"`    //
	Rectangle string `json:"rectangle,omitempty"` // 运营商
}

// AliIPBelongToJSON 阿里ip地址,返回凭借原始数据
func AliIPBelongToJSON(ip string) (string, error) {
	var baseURL = "http://iploc.market.alicloudapi.com/v3/ip?ip="
	var url = baseURL + ip
	cilent := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "APPCODE "+aliAppCode)
	resp, err := cilent.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	return string(bs), err
}

// AliIPBelongToStruct 主工具函数
func AliIPBelongToStruct(ip string) (*AliIP, error) {
	var res = new(AliIP)
	str, err := AliIPBelongToJSON(ip)
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
