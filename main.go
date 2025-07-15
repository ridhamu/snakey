package main

import (
	"bytes"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Point struct {
	x, y int
}

type Game struct {
	Snake      []Point
	Direction  Point
	Lastupdate time.Time
	Food       Point
	Gameover   bool
}

var (
	dirUp           = Point{x: 0, y: -1}
	dirDown         = Point{x: 0, y: 1}
	dirRight        = Point{x: 1, y: 0}
	dirLeft         = Point{x: -1, y: 0}
	mplusFaceSource *text.GoTextFaceSource
)

const (
	speed        = time.Second / 6
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 20
)

func (g Game) isBadCollision(newHead Point, snake *[]Point) bool {
	if newHead.x < 0 || newHead.y < 0 || newHead.x >= screenWidth/gridSize || newHead.y >= screenHeight/gridSize {
		return true
	}

	for _, bodySnake := range *snake {
		if bodySnake == newHead {
			return true
		}
	}

	return false
}

func (g *Game) UpdateSnake(snake *[]Point, direction *Point) {
	// get the current head
	head := (*snake)[0]

	// create the new head by adding the direction to the current head
	newHead := Point{
		x: head.x + direction.x,
		y: head.y + direction.y,
	}

	// check for bad collion e.g snake itself or out of bound
	if g.isBadCollision(newHead, &g.Snake) {
		g.Gameover = true
		return
	}

	// collision detection with the Food
	if newHead == g.Food {
		g.SpanwFood()
		*snake = append(
			[]Point{newHead},
			*snake...,
		)
	} else {
		// update the snake
		*snake = append(
			[]Point{newHead},
			(*snake)[:len(*snake)-1]...,
		)
	}
}

func (g *Game) SpanwFood() {
	g.Food = Point{
		x: rand.Intn(screenWidth / gridSize),
		y: rand.Intn(screenHeight / gridSize),
	}
}

func (g *Game) Update() error {
	if g.Gameover {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Direction = dirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Direction = dirLeft
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Direction = dirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Direction = dirRight
	}

	if time.Since(g.Lastupdate) < speed {
		return nil
	}
	g.Lastupdate = time.Now()

	g.UpdateSnake(&g.Snake, &g.Direction)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, p := range g.Snake {
		vector.DrawFilledRect(
			screen,
			float32(p.x*gridSize),
			float32(p.y*gridSize),
			gridSize,
			gridSize,
			color.White,
			true,
		)
	}

	// drawing food here
	vector.DrawFilledRect(
		screen,
		float32(g.Food.x*gridSize),
		float32(g.Food.y*gridSize),
		gridSize,
		gridSize,
		color.RGBA{255, 0, 0, 255},
		true,
	)

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
			screenWidth/2-w/2,
			screenHeight/2-h/2,
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
	return screenWidth, screenHeight
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

	g := &Game{
		Snake: []Point{
			{
				x: screenWidth / gridSize / 2,
				y: screenHeight / gridSize / 2,
			}},
		Direction: Point{
			x: 1,
			y: 0,
		},
	}

	g.SpanwFood()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snakey")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
