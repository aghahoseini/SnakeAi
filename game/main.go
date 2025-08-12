package game

import (
	//"fmt"
	"image/color"
	"log"
	"snake/modules/PathRecommender"
	"snake/modules/board"
	"snake/modules/canvas"
	"snake/modules/config"

	"snake/HelperModels"
	"snake/modules/food"
	"snake/modules/snake"

	"github.com/hajimehoshi/ebiten/v2"
)

func Start() {
	loader := config.NewConfigLoader("settings.toml")
	settings, _ := loader.LoadSettings()

	screenWidth        := settings["ScreenWidth"].(int)
	screenHeight       := settings["ScreenHeight"].(int)
	TileSize           := settings["TileSize"].(int)
	GridColor          := settings["GridColor"].(color.RGBA)
	SnakeBodyColor     := settings["SnakeBodyColor"].(color.RGBA)
	SnakeHeadColor     := settings["SnakeHeadColor"].(color.RGBA)
	//SnakeTailColor     := settings["SnakeTailColor"].(color.RGBA)
	//FoodColor          := settings["FoodColor"].(color.RGBA)
	BackgroundColor    := settings["BackgroundColor"].(color.RGBA)

	MaxX := screenWidth / TileSize
	MaxY := screenHeight / TileSize


    Painter := canvas.NewPainter(TileSize,screenWidth,screenHeight,nil)
    BoardRenderer := board.NewBoardRenderer(Painter,GridColor,BackgroundColor,screenWidth,screenHeight,TileSize)
	InitialSnakePosition := []HelperModels.Position{
		{X_Axis: 8 , Y_Axis: 8},
		{X_Axis: 9 , Y_Axis: 8},
		{X_Axis: 10 , Y_Axis: 8},
	}


    snake := snake.NewSnake(InitialSnakePosition,SnakeHeadColor,SnakeBodyColor,Painter)
	foodspawner := food.NewFoodSpawner(Painter,MaxX,MaxY)

	pathrecommender := pathrecommender.NewPathRecommender(snake,MaxX,MaxY)





	game := NewSnakeGame(Painter,BoardRenderer,snake,foodspawner,pathrecommender)

	ebiten.SetWindowSize(screenWidth, screenWidth)


	
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}