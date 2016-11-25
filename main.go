package main

import (
	"fmt"

	"github.com/ipiao/metools/utils"
)

func main() {
	fmt.Println(utils.GetCoord("浙江省杭州市萧山区伟七路南50米丰树物流园"))
	var origins []string
	var destinations []string
	var o = utils.GetCoord("杭州市拱墅区天堂E谷")
	fmt.Println(o)
	origins = append(origins, "120.11525567078871,30.342446000305196")
	var d = utils.GetCoord("杭州市祥园路德信北海公园")
	fmt.Println(d)
	destinations = append(destinations, "120.12410124292511，30.34308337010581")
	var res, err = utils.GetDistanceAndTime(origins, destinations)
	fmt.Println(res, err)
}
