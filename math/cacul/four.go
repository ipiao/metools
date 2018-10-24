package cacul

import "regexp"

const (
	numberReg = `^(-?\d+)(\.\d+)?$`
)

var numRegexp = regexp.MustCompile(numberReg)

func IsNumber(s string) bool {
	return numRegexp.MatchString(s)
}

func parseExp(s string) ([]float64, []string) {
	numRegexp.FindAllStringIndex()
}
