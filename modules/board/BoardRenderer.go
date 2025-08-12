package board


import (
	"image/color"
	"snake/modules/canvas"

)

type BoardRenderer struct {
	painter  	  *canvas.Painter
	GridColor     color.Color
	BackgroundColor color.Color
	ScreenWidth     int
	ScreenHeight    int
	TileSize        int
}

func NewBoardRenderer(paint *canvas.Painter, gridColor color.Color ,backclr color.Color, screenwidth int , screenheight int , tilesize int ) *BoardRenderer {
	return &BoardRenderer{
		painter:       paint,
		GridColor:       gridColor,
		BackgroundColor: backclr,
		ScreenWidth: screenwidth,
		ScreenHeight: screenheight,
		TileSize: tilesize,
	}
}




func (br *BoardRenderer) CreateGridSystem() {


	br.painter.FillBackground(br.BackgroundColor)
	br.draw_vertical_grid_lines();
	br.draw_horizontal_grid_lines();



}


func (br *BoardRenderer) FillBackground() {

	br.painter.FillBackground(br.BackgroundColor)

}


func (br *BoardRenderer) draw_vertical_grid_lines(){

	for x := 0; x <= br.ScreenWidth; x += br.TileSize{
		br.painter.FillPixelsBelow(x,br.GridColor)
	}

}



func (br *BoardRenderer) draw_horizontal_grid_lines(){

	for y := 0; y <= br.ScreenHeight; y += br.TileSize {
		br.painter.FillPixelsAhead(y,br.GridColor)
	}

}

