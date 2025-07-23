package entity

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Draw(screen *ebiten.Image)
	Update(world WorldView) bool
	Tag() string
}
