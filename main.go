package main

import (
	"fmt"

	"github.com/ipiao/metools/utils"
)

func main() {
	//var res = utils.GetAddress(`30.34308337010581,120.12410124292511`)
	var res2 = utils.GetPlanningRoute("39.848585906582,116.25989391769", "40.045728154863,116.31148692685")
	//fmt.Println(res)
	fmt.Println(res2)
}
