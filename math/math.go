package math

import (
	"slices"
	"math/rand"

	"github.com/ridhamu/snakey/common"
)

type Point struct {
	X, Y int
}

var (
	DirUp    = Point{X: 0, Y: -1}
	DirDown  = Point{X: 0, Y: 1}
	DirRight = Point{X: 1, Y: 0}
	DirLeft  = Point{X: -1, Y: 0}
)

func (p Point) Equals(other Point) bool {
	return other.X == p.X && other.Y == p.Y
}

func (p Point) Add(new Point) Point {
	return Point{
		X: p.X + new.X,
		Y: p.Y + new.Y,
	}
}

func (p Point) IsBadCollision(points []Point) bool {
	if p.X < 0 || p.Y < 0 || p.X >= common.ScreenWidth/common.GridSize || p.Y >= common.ScreenHeight/common.GridSize {
		return true
	}

	return slices.Contains(points, p)
}

func RandomPosition() Point {
	return Point{
		X: rand.Intn(common.ScreenWidth / common.GridSize),
		Y: rand.Intn(common.ScreenHeight / common.GridSize),
	}
}
