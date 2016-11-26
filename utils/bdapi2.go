package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Descripition struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type BDMapResult struct {
	Distance Descripition `json:"distance"`
	Duration Descripition `json:"duration"`
}

// 结构体
type BDMapInfo struct {
	Status  int           `json:"status"`
	Result  []BDMapResult `json:"result"`
	Message string        `json:"message"`
}

//
const (
	//ak = "alNeAGUdvXrKmcRyACrWYu9cfnANGQlw"
	ak         = "6zlc1bckc3iBxLaqpHoqo78mIfeZyGDp"
	output     = "json"
	coord_type = "gcj02" //bd09ll
)

func GetCoord(address string) string {
	var httpUrl = "http://api.map.baidu.com/geocoder/v2/"
	var data = make(url.Values)
	data["ak"] = []string{ak}
	data["output"] = []string{output}
	data["address"] = []string{address}
	data["city"] = []string{"北京市"}
	data["tactics"] = []string{"13"} //可选值：10 不走高速；11 最短时间；12 最短路径；13 最短 距离(不考虑路况); 
	resp, err := http.PostForm(httpUrl, data)
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}

// origins,destinations 格式 经度,纬度|经度,纬度,"|"分开支持多个出发点或终点
func GetDistanceAndTime(origins string, destinations string) *BDMapInfo {

	var httpUrl = "http://api.map.baidu.com/routematrix/v2/driving?"
	//	var httpUrl = "http://api.map.baidu.com/routematrix/v2/riding?"  //骑行
	//	var httpUrl = "http://api.map.baidu.com/routematrix/v2/walking?" //步行
	// "tactics" = 13 //10 不走高速；11 最短时间；12 最短路径；13 最短 距离(不考虑路况);
	httpUrl += "output=" + output + "&origins=" + origins + "&destinations=" + destinations + "&ak=" + ak

	resp, err := http.Get(httpUrl)
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var info BDMapInfo
	err = json.Unmarshal(res, &info)
	if err != nil {
		fmt.Println(err)
	}
	return &info
}
