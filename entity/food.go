package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/ridhamu/snakey/common"
	"github.com/ridhamu/snakey/math"
)

var _ Entity = (*Food)(nil)

type Food struct {
	position math.Point
}

func NewFood() *Food {
	return &Food{position: math.RandomPosition()}
}

func (f *Food) Update(world WorldView) bool {
	return false
}

func (f *Food) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		float32(f.position.X*common.GridSize),
		float32(f.position.Y*common.GridSize),
		common.GridSize,
		common.GridSize,
		color.RGBA{255, 0, 0, 255},
		true,
	)
}

func (f *Food) Respawn() {
	f.position = math.RandomPosition()
}

func (f *Food) Tag() string {
	return "food"
}
