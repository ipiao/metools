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
	//ak         = "6zlc1bckc3iBxLaqpHoqo78mIfeZyGDp"
	ak         = "RfPDxNiIPk6hUwdfXO2progMha5V5qej"
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
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
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

//--------------------- 获取规划路径
type PlanningRoute struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Type    int    `json:"type"`
	Info    struct {
		Copyright struct {
			Text     string `json:"text"`
			ImageUrl string `json:"imageUrl"`
		} `json:"copyright"`
	} `json:"info"`
	Result struct {
		Routers []Router `json:"routes"`
		Origin  struct {
			OriginPt Location      `json:"originPt"`
			Wd       string        `json:"wd"`
			Uid      int           `json:"uid"`
			AreaId   int           `json:"area_id"`
			CityName string        `json:"cname"`
			ListType int           `json:"listType"`
			Content  []ContentInfo `json:"content"`
		} `json:"origin"`

		Destination struct {
			OriginPt Location      `json:"originPt"`
			Wd       string        `json:"wd"`
			Uid      int           `json:"uid"`
			AreaId   int           `json:"area_id"`
			CityName string        `json:"cname"`
			ListType int           `json:"listType"`
			Content  []ContentInfo `json:"content"`
		} `json:"destination"`
		OriginInfo      PointInfo `josn:"originInfo"`
		DestinationInfo PointInfo `json:"destinationInfo"`
	} `json:"result"`
}

type PointInfo struct {
	OriginPt Location      `json:"originPt"`
	Wd       string        `json:"wd"`
	Uid      int           `json:"uid"`
	AreaId   int           `json:"area_id"`
	CityName string        `json:"cname"`
	ListType int           `json:"listType"`
	Content  []ContentInfo `json:"content"`
}

type ContentInfo struct {
	Name      string   `json:"name"`
	Address   string   `json:"address"`
	Uid       string   `json:"uid"`
	TelePhone string   `json:"telephone"`
	Location  Location `json:"location"`
}

type Router struct {
	Distance   int    `json:"distance"`
	Duration   int    `json:"duration"`
	ArriveTime string `json:"arrive_time"`
	Price      int    `json:"price"`
	Steps      []Step `json:"steps"`
}

type Step struct {
	Area                    int      `json:"area"`
	Direction               int      `json:"direction"`
	Distance                int      `json:"distance"`
	Duration                int      `json:"duration"`
	Instructions            string   `json:"instructions"`
	Path                    string   `json:"path"`
	Pois                    []Poi    `json:"pois"`
	Turn                    int      `json:"turn"`
	Type                    int      `json:"type"`
	StepOriginLocation      Location `json:"stepOriginLocation"`
	StepDestinationLocation Location `json:"stepDestinationLocation"`
	StepOriginInstruction   string   `json:"stepOriginInstruction"`
	TrafficCondition        int      `json:"traffic_condition"`
}

// tactics 10：不走高速；11：常规路线；12：最短路径；13：躲避拥堵
func GetPlanningRoute(origin, destination string, tactics ...string) *PlanningRoute {
	var tactic = "12"
	if len(tactics) > 0 && tactics[0] >= "10" && tactics[0] <= "13" {
		tactic = tactics[0]
	}
	var origin_region = GetAddress(origin).Result.AddressComponent.City
	var destination_region = GetAddress(destination).Result.AddressComponent.City
	fmt.Println(origin_region, destination_region)
	var httpUrl = "http://api.map.baidu.com/direction/v1?"
	var mode = "driving"
	httpUrl += "origin=" + origin + "&origin_region=" + origin_region + "&destination=" + destination + "&destination_region=" +
		destination_region + "&ak=" + ak + "&tactics=" + tactic + "&mode=" + mode + "&output=" + output
	resp, err := http.Get(httpUrl)
	fmt.Println(httpUrl)
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(res))
	var route PlanningRoute
	err = json.Unmarshal(res, &route)
	if err != nil {
		fmt.Println(err)
	}
	return &route
}

// 通过坐标查地址

type Location struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type AddressComponent struct {
	Country      string `json:"country"`
	CountryCode  int    `json:"country_code"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Adcode       string `json:"adcode"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
	Direction    string `json:"direction"`
	Distance     string `json:"distance"`
}

type Poi struct {
	Addr      string `json:"addr"`
	Cp        string `json:"cp"`
	Direction string `json:"direction"`
	Distance  string `json:"distance"`
	Name      string `json:"name"`
	PoiType   string `json:"poiType"`
	Point     Point  `json:"point"`
	Tag       string `json:"tag"`
	Tel       string `json:"tel"`
	Uid       string `json:"uid"`
	Zip       string `json:"zip"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type BDMapAddress struct {
	Status int `json:"status"`
	Result struct {
		Location         Location         `json:"location"`
		FormattedAddress string           `json:"formatted_address"`
		Business         string           `json:"business"`
		AddressComponent AddressComponent `json:"addressComponent"`
	} `json:"result"`
	//	Pois               []Poi    `json:"pois"`
	//	PoiRegions         []string `josn:"poiRegions"`
	//	SematicDescription string   `josn:"sematic_description"`
	//	CityCode           int      `json:"cityCode"`
}

// 坐标查地址 coord=经度，纬度
func GetAddress(coord string) *BDMapAddress {
	var httpUrl = "http://api.map.baidu.com/geocoder/v2/?"
	httpUrl += "output=" + output + "&location=" + coord + "&ak=" + ak + "&pois=1"
	resp, err := http.Get(httpUrl)
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var address BDMapAddress
	err = json.Unmarshal(res, &address)
	if err != nil {
		fmt.Println(err)
	}
	return &address
}
