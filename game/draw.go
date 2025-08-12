package game

import (

	"github.com/hajimehoshi/ebiten/v2"
	//"image/color"
	//"fmt"
)




func (g *SnakeGame) Draw(screen *ebiten.Image) {


    g.Painter.Screen = screen
    
    g.BoardRenderer.FillBackground()
	g.BoardRenderer.CreateGridSystem()

	g.FoodSpawner.Draw()
	g.SnakeController.Draw()

	



}
