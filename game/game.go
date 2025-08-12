package game

import (
	"snake/modules/canvas"
	"snake/modules/board"
	"snake/modules/snake"
	"snake/modules/food"
	"time"
	"snake/modules/PathRecommender"

)
type SnakeGame struct {
	Painter     *canvas.Painter
	BoardRenderer  *board.BoardRenderer
	SnakeController  *snake.SnakeController
	FoodSpawner  *food.FoodSpawner
	MoveInterval time.Duration
	lastUpdate  time.Time
	path_recommender *pathrecommender.PathRecommender



}

func NewSnakeGame( painter *canvas.Painter , boardrenderer *board.BoardRenderer, snake *snake.SnakeController , food_spawner *food.FoodSpawner , path_recommender *pathrecommender.PathRecommender  ) *SnakeGame{
    return &SnakeGame{
		Painter: painter,
		BoardRenderer: boardrenderer,
		SnakeController: snake,
		FoodSpawner: food_spawner,
		MoveInterval: 20 * time.Millisecond,
		path_recommender: path_recommender,

	}
}


