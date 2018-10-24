package line

import "github.com/ipiao/metools/math/vector"

// Point 平面点
type Point struct {
	X float64
	Y float64
}

// ToVector 平面点转换为向量
func (p *Point) ToVector() *vector.Vector {
	var v = vector.Vector([]float64{p.X, p.Y})
	return &v
}

// Same 判断点是否相同
func (p *Point) Same(p1 *Point) bool {
	return p1.X == p.X && p1.Y == p.Y
}
