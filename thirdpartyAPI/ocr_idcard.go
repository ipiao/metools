package thirdapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// OcrIDCardInput 印刷字识别-身份证-请求
type OcrIDCardInput struct {
	Image struct {
		DataType  int    `json:"dataType,omitempty"`
		DataValue string `json:"dataValue,omitempty"` //图片以base64编码的string
	} `json:"image,omitempty"`
	Configure struct {
		DataType  int `json:"dataType,omitempty"`
		DataValue struct {
			Side string `json:"side,omitempty"`
		} `json:"dataValue,omitempty"` //身份证正反面类型:face/back
	} `json:"configure,omitempty"`
}

// OcrIDCardInputs 印刷字识别-身份证-请求
type OcrIDCardInputs struct {
	Inputs []OcrIDCardInput `json:"inputs,omitempty"`
}

// OcrIDCardResp 印刷字识别-身份证-返回
type OcrIDCardResp struct {
	OutPuts []struct {
		OutputLabel string
		OutputMulti struct{}
		OutputValue struct {
			DataType  int
			DataValue string
		}
	}
}

// AliOcrIDCardJSON 阿里印刷字识别-身份证
func AliOcrIDCardJSON(inputs *OcrIDCardInputs) (string, error) {

	bs, err := json.Marshal(inputs)
	if err != nil {
		return "", errors.New("json解析错误")
	}
	var baseURL = "http://dm-51.data.aliyun.com/rest/160601/ocr/ocr_idcard.json"
	cilent := &http.Client{}
	body := bytes.NewReader(bs)
	req, _ := http.NewRequest("POST", baseURL, body)
	req.Header.Add("Authorization", "APPCODE "+aliAppCode)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	resp, err := cilent.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	rbs, err := ioutil.ReadAll(resp.Body)
	return string(rbs), err
}

// AliOcrIDCardStruct 主工具函数
func AliOcrIDCardStruct(inputs *OcrIDCardInputs) (*OcrIDCardResp, error) {
	var res = new(OcrIDCardResp)
	str, err := AliOcrIDCardJSON(inputs)
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

// TransPicToOcrIDCardInput 将图片转换为结构体
func TransPicToOcrIDCardInput(pic, picType string) (OcrIDCardInput, error) {
	var input = OcrIDCardInput{}
	// file, _ := os.OpenFile(pic, os.O_RDONLY, os.ModePerm)
	// var bs = make([]byte, 1024*50)
	// n, err := file.Read(bs)
	bs, err := ioutil.ReadFile(pic)
	if err != nil {
		return input, err
	}
	input.Image.DataType = 50
	input.Image.DataValue = base64.StdEncoding.EncodeToString(bs)
	input.Configure.DataType = 50
	input.Configure.DataValue.Side = picType
	return input, err
}

// TransIDCardsPicsToOcrIDCardInput 转换身份证
func TransIDCardsPicsToOcrIDCardInput(face, back string) (*OcrIDCardInputs, error) {
	f, err := TransPicToOcrIDCardInput(face, "face")
	if err != nil {
		return nil, err
	}
	s, err := TransPicToOcrIDCardInput(back, "back")
	if err != nil {
		return nil, err
	}
	var inputs = new(OcrIDCardInputs)
	inputs.Inputs = append(inputs.Inputs, f, s)
	return inputs, nil
}

// TransPicToOcrIDCardInputStr 将图片转换为str
func TransPicToOcrIDCardInputStr(pic, picType string) (string, error) {
	var input string
	bs, err := ioutil.ReadFile(pic)
	if err != nil {
		return input, err
	}
	input = fmt.Sprintf(`"{
		"inputs": [
			{
				"image": {
					"dataType": 50,
					"dataValue": "%s"
				},
				"configure": {
					"dataType": 50,
					"dataValue": "{\"side\":\"%s\"}"
				}
			}
		]
	}"`, base64.StdEncoding.EncodeToString(bs), picType)
	return input, err
}

// AliOcrIDCardToJSON 阿里印刷字识别-身份证
func AliOcrIDCardToJSON(inputs string) (string, error) {

	var baseURL = "http://dm-51.data.aliyun.com/rest/160601/ocr/ocr_idcard.json"
	cilent := &http.Client{}
	body := strings.NewReader(inputs)
	req, _ := http.NewRequest("POST", baseURL, body)
	req.Header.Add("Authorization", "APPCODE "+aliAppCode)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	resp, err := cilent.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	rbs, err := ioutil.ReadAll(resp.Body)
	return string(rbs), err
}
