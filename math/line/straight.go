package line

import "errors"

// PlainStraightLine 平面直线(二维)
// Ax+By+C=0
type PlainStraightLine struct {
	A float64
	B float64
	C float64
}

// GetStraightLineWithTwoPoint 两点一线
func GetStraightLineWithTwoPoint(p1, p2 *Point) (*PlainStraightLine, error) {
	if p1.Same(p1) {
		return nil, errors.New("two points can not be same")
	}
	var l = new(PlainStraightLine)
	// [x1,y1;x2,y2][A;B]=[-C;-C],用矩阵求A、B
	l.A = p1.Y - p2.Y
	l.B = p2.X - p1.X
	l.C = p1.X*p2.Y - p2.X*p1.Y
	l.Simple()
	return l, nil
}

// GetStraightLinePointSlope 点斜法
func GetStraightLinePointSlope(p *Point, r float64) *PlainStraightLine {
	var l = new(PlainStraightLine)
	if r == 0 {
		l.B = 1
		l.C = -l.B
		return l
	}
	l.B = 1
	l.A = -r
	l.C = r*p.X - p.Y
	return l
}

// GetPerpendicularLineWithTwoPoint 获取中垂线
func GetPerpendicularLineWithTwoPoint(p1, p2 *Point) (*PlainStraightLine, error) {
	if p1.Same(p1) {
		return nil, errors.New("two points can not be same")
	}
	var l = new(PlainStraightLine)
	if p1.Y == p2.Y {
		l.A = 1
		l.C = -(p1.X + p2.X) / 2
		return l, nil
	}
	p := &Point{X: (p1.X + p2.X) / 2, Y: (p1.Y + p2.Y) / 2}
	r := -(p2.X - p1.X) / (p2.Y - p1.Y)
	return GetStraightLinePointSlope(p, r), nil
}

// Simple 最简化
// Ax+By+C=0 => A/Bx+y+C/A=0
func (l *PlainStraightLine) Simple() {
	if l.A == 0 && l.B == 0 {
		return
	}
	if l.A == 0 {
		l.C /= l.B
		l.B = 1
	} else if l.B == 0 {
		l.C /= l.A
		l.A = 1
	} else {
		l.A /= l.B
		l.C /= l.B
		l.B = 1
	}
}

// ApproximatePoint 判断点是否近似在线上
func (l *PlainStraightLine) ApproximatePoint(p *Point, det float64) bool {
	return l.A*p.X+l.B*p.Y+l.C < det
}

// PointOnLine 判断点是否在线上
func (l *PlainStraightLine) PointOnLine(p *Point) bool {
	return l.A*p.X+l.B*p.Y+l.C == 0
}

// GetPointByX 根据X获取点坐标
func (l *PlainStraightLine) GetPointByX(x float64) (*Point, error) {
	if l.B == 0 {
		return nil, errors.New("l.B is 0,its not a certain line")
	}
	y := -(l.C + l.A*x) / l.B
	return &Point{X: x, Y: y}, nil
}

// GetPointByY 根据Y获取点坐标
func (l *PlainStraightLine) GetPointByY(y float64) (*Point, error) {
	if l.A == 0 {
		return nil, errors.New("l.A is 0,its not a certain line")
	}
	x := -(l.C + l.B*y) / l.A
	return &Point{X: x, Y: y}, nil
}
