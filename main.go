package main

import (
	"fmt"

	"github.com/ipiao/metools/utils"
)

func main() {
	var res = utils.GetDistanceAndTime(`30.34308337010581,120.12410124292511`,
		`30.342446000305196,120.11525567078871`)
	fmt.Println(res)
}
