package thirdapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// AliJiSuZNWD 阿里智能问答
type AliJiSuZNWD struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Result struct {
		Type        string `json:"type,omitempty"`
		Content     string `json:"content,omitempty"`
		Relquestion string `json:"relquestion,omitempty"`
	} `json:"result,omitempty"`
}

// AliJiSuZNWDToJSON 阿里智能问答
func AliJiSuZNWDToJSON(question string) (string, error) {
	var baseURL = "http://jisuznwd.market.alicloudapi.com/iqa/query?question="
	var url = baseURL + question
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

// AliJiSuZNWDToStruct 主工具函数
func AliJiSuZNWDToStruct(question string) (*AliJiSuZNWD, error) {
	var res = new(AliJiSuZNWD)
	str, err := AliJiSuZNWDToJSON(question)
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
