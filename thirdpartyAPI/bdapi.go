package metools

//import (
//	"errors"
//	"fmt"
//	"net/http"
//	"net/url"
//	"strings"
//)

//var (
//	ak   = "alNeAGUdvXrKmcRyACrWYu9cfnANGQlw"
//	coor = "bd09II"
//	//coor   = "gcj02" //国测局坐标
//	output = "json"
//)

//type CoordResult struct {
//	Lng string
//	Lat string
//}

//type BDResult1 struct {
//	Status     int
//	Result     CoordResult
//	Precise    int
//	Confidence int
//	Level      string
//}

//// 百度api，查找坐标
//func GetCoord(address string) string {

//	var httpUrl = "http://api.map.baidu.com/geocoder/v2/"

//	var data = make(url.Values)
//	data["ak"] = []string{ak}
//	data["coor"] = []string{coor}
//	data["output"] = []string{output}
//	data["address"] = []string{address}
//	res, err := http.PostForm(httpUrl, data)
//	if err != nil {
//		fmt.Println(err)
//	}
//	var result = make([]byte, 1024)
//	res.Body.Read(result)
//	//fmt.Println(string(result))
//	return string(result)
//}

//// 查询起点和终点之间的距离，和时间
//func GetDistanceAndTime(origins []string, destinations []string) (string, error) {
//	var httpUrl = "http://api.map.baidu.com/routematrix/v2/driving?" //驾车
//	if len(origins)*len(origins) > 50 {
//		return "", errors.New("原点数与终点数乘积不能超过50")
//	}
//	//	var httpUrl = "http://api.map.baidu.com/routematrix/v2/riding?"  //骑行
//	//	var httpUrl = "http://api.map.baidu.com/routematrix/v2/walking?" //步行
//	var data = make(url.Values)
//	data["origins"] = []string{strings.Join(origins, "|")}
//	data["destinations"] = []string{strings.Join(destinations, "|")}
//	data["tactics"] = []string{"13"} //10 不走高速；11 最短时间；12 最短路径；13 最短 距离(不考虑路况); 
//	data["ak"] = []string{ak}
//	data["coord_type"] = []string{coor}
//	data["output"] = []string{output}
//	res, err := http.PostForm(httpUrl, data)
//	if err != nil {
//		fmt.Println(err)
//	}
//	var result = make([]byte, 1024)
//	res.Body.Read(result)
//	fmt.Println(string(result))
//	return string(result), nil
//}
