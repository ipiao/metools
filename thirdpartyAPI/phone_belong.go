package thirdapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// AliMobile 阿里手机号码归属地查询
type AliMobile struct {
	ShowAPIResCode  int    `json:"showapi_res_code,omitempty"`
	ShowAPIResError string `json:"showapi_res_error,omitempty"`
	ShowAPIResBody  struct {
		RetCode      int    `json:"ret_code,omitempty"` //0为成功，其他失败。失败时不扣点数
		Province     string `json:"prov,omitempty"`     //省份
		AreaCode     string `json:"areaCode,omitempty"` //区号
		Name         string `json:"name,omitempty"`     //运营商名称
		PostCode     string `json:"postCode,omitempty"` //邮编
		ProvinceCode string `json:"provCode,omitempty"` //此地区身份证号开头几位
		Type         int    `json:"type,omitempty"`     //1移动    2电信    3联通
		City         string `json:"city,omitempty"`     //城市
	} `json:"showapi_res_body,omitempty"`
}

// AliMobileBelongToJSON 主工具函数,返回凭借原始数据
func AliMobileBelongToJSON(phoneNum string) (string, error) {
	var baseURL = "http://ali-mobile.showapi.com/6-1?num="
	var url = baseURL + phoneNum
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

// AliMobileBelongToStruct 主工具函数
func AliMobileBelongToStruct(phoneNum string) (*AliMobile, error) {
	var res = new(AliMobile)
	str, err := AliMobileBelongToJSON(phoneNum)
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
