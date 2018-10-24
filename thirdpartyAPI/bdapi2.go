package thirdapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Descripition 描述
type Descripition struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

// BDMapResult 结果
type BDMapResult struct {
	Distance Descripition `json:"distance"`
	Duration Descripition `json:"duration"`
}

// BDMapInfo 结构体
type BDMapInfo struct {
	Status  int           `json:"status"`
	Result  []BDMapResult `json:"result"`
	Message string        `json:"message"`
}

//
const (
	//ak = "alNeAGUdvXrKmcRyACrWYu9cfnANGQlw"
	//ak         = "6zlc1bckc3iBxLaqpHoqo78mIfeZyGDp"
	ak        = "RfPDxNiIPk6hUwdfXO2progMha5V5qej"
	output    = "json"
	coordType = "gcj02" //bd09ll
)

// GetCoord 获取坐标
func GetCoord(address string) string {
	var httpURL = "http://api.map.baidu.com/geocoder/v2/"
	var data = make(url.Values)
	data["ak"] = []string{ak}
	data["output"] = []string{output}
	data["address"] = []string{address}
	data["city"] = []string{"北京市"}
	data["tactics"] = []string{"13"} //可选值：10 不走高速；11 最短时间；12 最短路径；13 最短 距离(不考虑路况); 
	resp, err := http.PostForm(httpURL, data)
	if err != nil {
		return ""
	}
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

// GetDistanceAndTime origins,destinations 格式 经度,纬度|经度,纬度,"|"分开支持多个出发点或终点
func GetDistanceAndTime(origins string, destinations string) *BDMapInfo {

	var httpURL = "http://api.map.baidu.com/routematrix/v2/driving?"
	//	var httpUrl = "http://api.map.baidu.com/routematrix/v2/riding?"  //骑行
	//	var httpUrl = "http://api.map.baidu.com/routematrix/v2/walking?" //步行
	// "tactics" = 13 //10 不走高速；11 最短时间；12 最短路径；13 最短 距离(不考虑路况);
	httpURL += "output=" + output + "&origins=" + origins + "&destinations=" + destinations + "&ak=" + ak

	resp, err := http.Get(httpURL)
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

// PlanningRoute 获取规划路径
type PlanningRoute struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Type    int    `json:"type"`
	Info    struct {
		Copyright struct {
			Text     string `json:"text"`
			ImageURL string `json:"imageUrl"`
		} `json:"copyright"`
	} `json:"info"`
	Result struct {
		Routers []Router `json:"routes"`
		Origin  struct {
			OriginPt Location      `json:"originPt"`
			Wd       string        `json:"wd"`
			UID      int           `json:"uid"`
			AreaID   int           `json:"area_id"`
			CityName string        `json:"cname"`
			ListType int           `json:"listType"`
			Content  []ContentInfo `json:"content"`
		} `json:"origin"`

		Destination struct {
			OriginPt Location      `json:"originPt"`
			Wd       string        `json:"wd"`
			UID      int           `json:"uid"`
			AreaID   int           `json:"area_id"`
			CityName string        `json:"cname"`
			ListType int           `json:"listType"`
			Content  []ContentInfo `json:"content"`
		} `json:"destination"`
		OriginInfo      PointInfo `josn:"originInfo"`
		DestinationInfo PointInfo `json:"destinationInfo"`
	} `json:"result"`
}

// PointInfo 点信息
type PointInfo struct {
	OriginPt Location      `json:"originPt"`
	Wd       string        `json:"wd"`
	UID      int           `json:"uid"`
	AreaID   int           `json:"area_id"`
	CityName string        `json:"cname"`
	ListType int           `json:"listType"`
	Content  []ContentInfo `json:"content"`
}

// ContentInfo 内容
type ContentInfo struct {
	Name      string   `json:"name"`
	Address   string   `json:"address"`
	UID       string   `json:"uid"`
	TelePhone string   `json:"telephone"`
	Location  Location `json:"location"`
}

// Router 路线
type Router struct {
	Distance   int    `json:"distance"`
	Duration   int    `json:"duration"`
	ArriveTime string `json:"arrive_time"`
	Price      int    `json:"price"`
	Steps      []Step `json:"steps"`
}

// Step 步骤
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

// GetPlanningRoute tactics 10：不走高速；11：常规路线；12：最短路径；13：躲避拥堵
func GetPlanningRoute(origin, destination string, tactics ...string) *PlanningRoute {
	var tactic = "12"
	if len(tactics) > 0 && tactics[0] >= "10" && tactics[0] <= "13" {
		tactic = tactics[0]
	}
	var originRegion = GetAddress(origin).Result.AddressComponent.City
	var destinationRegion = GetAddress(destination).Result.AddressComponent.City
	fmt.Println(originRegion, destinationRegion)
	var httpURL = "http://api.map.baidu.com/direction/v1?"
	var mode = "driving"
	httpURL += "origin=" + origin + "&origin_region=" + originRegion + "&destination=" + destination + "&destination_region=" +
		destinationRegion + "&ak=" + ak + "&tactics=" + tactic + "&mode=" + mode + "&output=" + output
	resp, err := http.Get(httpURL)
	fmt.Println(httpURL)
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

// Location 通过坐标查地址
type Location struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

// AddressComponent 地址
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

// Poi poi
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
	UID       string `json:"uid"`
	Zip       string `json:"zip"`
}

// Point 点
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// BDMapAddress 地址
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

// GetAddress 坐标查地址 coord=经度，纬度
func GetAddress(coord string) *BDMapAddress {
	var address BDMapAddress
	var httpURL = "http://api.map.baidu.com/geocoder/v2/?"
	httpURL += "output=" + output + "&location=" + coord + "&ak=" + ak + "&pois=1"
	resp, err := http.Get(httpURL)
	if err != nil {
		return &address
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = json.Unmarshal(res, &address)
	if err != nil {
		fmt.Println(err)
	}
	return &address
}
