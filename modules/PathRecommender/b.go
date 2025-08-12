package pathrecommender

import (

	"snake/HelperModels"
	"snake/modules/snake"
	//"os"
)

type PathRecommender struct {
	v_snake                        VirtualSnake
	OrginalSnake                   *snake.SnakeController
	FoodPos                        HelperModels.Position
	MaxX                           int
	MaxY                           int
	Path_h1_to_f                   []HelperModels.Position
	Path_h1_to_t                   []HelperModels.Position
	Path_h2_to_t                   []HelperModels.Position
	snake_body_after_reaching_food []HelperModels.Position
}

func NewPathRecommender(snake *snake.SnakeController, maxx int, maxy int) *PathRecommender {
	return &PathRecommender{
		OrginalSnake: snake,
		MaxX:         maxx,
		MaxY:         maxy,
	}

}

func (pr *PathRecommender) create_virtual_snake_from_orginal_snake() *VirtualSnake {

	// Get the real snake's body
	snakeBody := pr.OrginalSnake.Body

	// Build a virtual snake with same body and TileEngine
	vSnake := VirtualSnake{
		Body:    make([]HelperModels.Position, len(snakeBody)),
		FoodPos: pr.FoodPos,
		MaxX:    pr.MaxX,
		MaxY:    pr.MaxY,
	}

	// Copy body (important: avoid referencing same slice)
	copy(vSnake.Body, snakeBody)

	return &vSnake

}

func (pr *PathRecommender) path_exist_to_food() bool {

	v_snake1 := pr.create_virtual_snake_from_orginal_snake()


	path_1 := v_snake1.bfs_from_head_to_food()
	pr.Path_h1_to_f = path_1

	if len(path_1) > 0 {
		return true
	}

	return false

}

func (pr *PathRecommender) can_reach_tail_after_eating_food() bool {

	path_2 := []HelperModels.Position{}
	v_snake2 := pr.create_virtual_snake_from_orginal_snake()

	for _, pos := range pr.Path_h1_to_f {
		v_snake2.move_one_step(pos)
	}

	path_2 = v_snake2.bfs_from_head_to_tail()
	pr.Path_h2_to_t = path_2

	if len(path_2) > 0 {
		return true
	}

	return false

}

func (pr *PathRecommender) follow_food() []HelperModels.Position {

	return pr.Path_h1_to_f

}
