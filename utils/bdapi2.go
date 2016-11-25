package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

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

//
func GetDistanceAndTime(origins string, destinations string) string {

	var httpUrl = "http://api.map.baidu.com/routematrix/v2/driving?"
	httpUrl += "output=" + output + "&origins=" + origins + "&destinations=" + destinations + "&ak=" + ak

	resp, err := http.Get(httpUrl)
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
