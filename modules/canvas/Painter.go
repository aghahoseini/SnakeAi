package canvas

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"snake/HelperModels"
)

type Painter struct {
	TileSize     int
	ScreenWidth  int
	ScreenHeight int
	Screen       *ebiten.Image
}

func NewPainter(tileSize, screenWidth, screenHeight int, screen *ebiten.Image) *Painter {
	return &Painter{
		TileSize:     tileSize,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		Screen:       screen,
	}
}

func (p *Painter) FillTileCentered(tilePos HelperModels.Position, clr color.Color) {
    startX := tilePos.X_Axis * p.TileSize
    startY := tilePos.Y_Axis * p.TileSize

    margin := float64(p.TileSize) * 0.15
    innerStartX := startX + int(margin)
    innerEndX := startX + p.TileSize - int(margin)
    innerStartY := startY + int(margin)
    innerEndY := startY + p.TileSize - int(margin)

    for y := innerStartY; y < innerEndY; y++ {
        for x := innerStartX; x < innerEndX; x++ {
            p.Screen.Set(x, y, clr)
        }
    }
}

func (p *Painter) FillTile(tilePos HelperModels.Position, clr color.Color) {
	startX := tilePos.X_Axis * p.TileSize
	startY := tilePos.Y_Axis * p.TileSize

	for y := 0; y < p.TileSize; y++ {
		for x := 0; x < p.TileSize; x++ {
			p.Screen.Set(startX+x, startY+y, clr)
		}
	}
}

func (p *Painter) FillPixel(pixelPos HelperModels.Position, clr color.Color) {
	p.Screen.Set(pixelPos.X_Axis, pixelPos.Y_Axis, clr)
}

func (p *Painter) FillBackground(clr color.Color) {
	p.Screen.Fill(clr)
}

func (p *Painter) FillPixelsBelow(pixelX int, clr color.Color) {
	for y := 0; y < p.ScreenHeight; y++ {
		pixelPos := HelperModels.Position{X_Axis: pixelX, Y_Axis: y}
		p.FillPixel(pixelPos, clr)
	}
}

func (p *Painter) FillPixelsAhead(pixelY int, clr color.Color) {
	for x := 0; x < p.ScreenWidth; x++ {
		pixelPos := HelperModels.Position{X_Axis: x, Y_Axis: pixelY}
		p.FillPixel(pixelPos, clr)
	}
}
