package pathrecommender

import (
	"fmt"
	"math/rand"
	"snake/HelperFunctions"
	"snake/HelperModels"
	//"snake/modules/snake"
	//"os"
)

func (pr *PathRecommender) v1_longest_path_to_tail() []HelperModels.Position {

	neighbors := HelperFunctions.GetAvailableNeighborTiles(pr.OrginalSnake.Body[0], pr.MaxX, pr.MaxY, pr.OrginalSnake.Body, pr.FoodPos)
	path := []HelperModels.Position{}

	if len(neighbors) > 0 {
		dis := -999

		for _, n := range neighbors {

			if Distance(n, pr.OrginalSnake.GetTailPos()) > dis {
				v_snake3 := pr.create_virtual_snake_from_orginal_snake()
				v_snake3.move_one_step(n)

				if len(v_snake3.bfs_from_head_to_tail()) > 0 {
					path = append(path, n)
					dis = Distance(n, pr.OrginalSnake.GetTailPos())
				}
			}

		}

		if len(path) > 0 {
			return []HelperModels.Position{path[len(path)-1]}
		}
	}

	return []HelperModels.Position{}

}

func (pr *PathRecommender) v1_any_safe_move() []HelperModels.Position {
	fmt.Println("start any safe move")
	neighbors := HelperFunctions.GetAvailableNeighborTiles(pr.OrginalSnake.Body[0], pr.MaxX, pr.MaxY, pr.OrginalSnake.Body, pr.FoodPos)
	path := []HelperModels.Position{}

	if len(neighbors) > 0 {

		randomIndex := rand.Intn(len(neighbors))
		path = append(path, neighbors[randomIndex])

		v_snake4 := pr.create_virtual_snake_from_orginal_snake()

		for _, move := range path {
			v_snake4.move_one_step(move)
		}

		if len(v_snake4.bfs_from_head_to_tail()) > 0 {
			fmt.Println("end any safe move 1")
			return path
		} else {

			v_snake5 := pr.create_virtual_snake_from_orginal_snake()
			fmt.Println("end any safe move 2")
			return v_snake5.bfs_from_head_to_tail()

		}

	}

	fmt.Println("end any safe move 3")
	return []HelperModels.Position{}

}

func (pr *PathRecommender) V1_Set_path(fpos HelperModels.Position) []HelperModels.Position {

	pr.FoodPos = fpos
	pr.Path_h1_to_f = []HelperModels.Position{{}}
	pr.Path_h1_to_t = []HelperModels.Position{{}}
	pr.Path_h2_to_t = []HelperModels.Position{{}}

	if pr.path_exist_to_food() && pr.can_reach_tail_after_eating_food() {

		fmt.Println("follow food")
		fmt.Println("full path to food --> ", pr.Path_h1_to_f)
		fmt.Println("next pos --> ", pr.Path_h1_to_f[0])
		fmt.Println("path to tail - after eating food --> ", pr.Path_h2_to_t)
		return pr.follow_food()

	} else {

		path := pr.v1_longest_path_to_tail()

		isEvenScore := pr.OrginalSnake.Score%2 == 0
		canMoveMore := pr.OrginalSnake.MovesWithoutEating < pr.OrginalSnake.MAX_MOVES_WITHOUT_EATING/2
		hasValidPath := len(path) > 0

		if hasValidPath && isEvenScore && canMoveMore {
			fmt.Println("option2")
			return path
		}

		path = pr.v1_any_safe_move()
		if len(path) > 0 {
			fmt.Println("option3")
			return path
		}

		v_snake6 := pr.create_virtual_snake_from_orginal_snake()
		path = v_snake6.bfs_from_head_to_tail()
		if len(path) > 0 {
			fmt.Println("option4")
			return path
		}

		fmt.Println("option5")
		return []HelperModels.Position{}

	}

	fmt.Println("option5")
	return []HelperModels.Position{}

}
