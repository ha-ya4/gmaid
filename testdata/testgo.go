package testdata

import (
	"fmt"
	"encoding/json"
)

// StName structの名前。このファイルに定義されてるならなんでもいい。
const StName = "Point"

type danmaku struct {
	x int
}

// Point point
type Point struct {
	XxPoint int
	YyPoint int
}

func (p *Point) call() {
	fmt.Printf("x:%v y:%v", p.XxPoint, p.YyPoint)
}

func(p *Point) toString() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func add(p1, p2 Point) Point {
	return Point{
		XxPoint: p1.XxPoint+p2.XxPoint,
		YyPoint: p1.YyPoint+p2.YyPoint,
	}
}