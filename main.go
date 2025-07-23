package main

import (
	"bytes"
	"errors"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/ridhamu/snakey/common"
	"github.com/ridhamu/snakey/entity"
	"github.com/ridhamu/snakey/game"
	"github.com/ridhamu/snakey/math"
)

type Game struct {
	world      *game.World
	Lastupdate time.Time
	Gameover   bool
}

var (
	mplusFaceSource *text.GoTextFaceSource
)

func (g *Game) Update() error {
	if g.Gameover {
		return nil
	}

	playerRaw, ok := g.world.GetFirstEntity("player")
	if !ok {
		return errors.New("entity player was not found")
	}

	player := playerRaw.(*entity.Player)

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		player.SetDirection(math.DirUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.SetDirection(math.DirLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		player.SetDirection(math.DirDown)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.SetDirection(math.DirRight)
	}

	if time.Since(g.Lastupdate) < common.Speed {
		return nil
	}
	g.Lastupdate = time.Now()

	for _, entity := range g.world.Entities() {
		if entity.Update(g.world) {
			g.Gameover = true
			return nil
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, entity := range g.world.Entities() {
		entity.Draw(screen)
	}

	// draw the game over text
	if g.Gameover {
		// draw the text here
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   48,
		}
		t := "Game Over"
		w, h := text.Measure(
			t,
			face,
			face.Size,
		)

		op := &text.DrawOptions{}
		op.GeoM.Translate(
			common.ScreenWidth/2-w/2,
			common.ScreenHeight/2-h/2,
		)

		op.ColorScale.ScaleWithColor(color.White)

		text.Draw(
			screen,
			t,
			face,
			op,
		)

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth, common.ScreenHeight
}

func main() {
	t, err := text.NewGoTextFaceSource(
		bytes.NewReader(
			fonts.MPlus1pRegular_ttf,
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	mplusFaceSource = t

	world := game.NewWorld()
	world.AddEntity(
		entity.NewPlayer(math.Point{
			X: common.ScreenWidth / common.GridSize / 2,
			Y: common.ScreenHeight / common.GridSize / 2,
		}, math.DirRight),
	)

	world.AddEntity(entity.NewFood())

	g := &Game{
		world: world,
	}

	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetWindowTitle("Snakey")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
