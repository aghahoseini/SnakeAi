package pathrecommender

import (
	"fmt"
	"snake/HelperFunctions"
	"snake/HelperModels"

)



func (pr *PathRecommender) v2_longest_path_to_tail() []HelperModels.Position {

	neighbors := HelperFunctions.GetAvailableNeighborTiles(pr.OrginalSnake.Body[0], pr.MaxX, pr.MaxY, pr.OrginalSnake.Body, pr.FoodPos)
	fmt.Println("len(neighbors) ",len(neighbors))
	path := []HelperModels.Position{}

	if len(neighbors) > 0 {
		fmt.Println("inside longest_path_to_tail  --> 1")
		dis := -999

		for _, n := range neighbors {
			fmt.Println("inside longest_path_to_tail  --> 2")

			if Distance(n, pr.OrginalSnake.GetTailPos()) > dis {
				fmt.Println("inside longest_path_to_tail  --> 3")
				v_snake7 := pr.create_virtual_snake_from_orginal_snake()
				fmt.Println("inside longest_path_to_tail  --> 4")
				v_snake7.move_one_step(n)
				fmt.Println("inside longest_path_to_tail  --> 5")
				pr.Path_h2_to_t = v_snake7.bfs_from_head_to_tail()
				fmt.Println("pr.Path_h2_to_t",pr.Path_h2_to_t)
				fmt.Println("inside longest_path_to_tail  --> 6")

				if len(pr.Path_h2_to_t) > 0 {
					fmt.Println("inside longest_path_to_tail  --> 7")
					path = append(path, n)
					fmt.Println("inside longest_path_to_tail  --> 8")
					dis = Distance(n, pr.OrginalSnake.GetTailPos())
					fmt.Println("inside longest_path_to_tail  --> 9")
				}
			}

		}
		fmt.Println("inside longest_path_to_tail  --> 10")
		//fmt.Println("path to tail - after going to   ", path[len(path)-1] ,"  ---> ", pr.Path_h2_to_t)


		if len(path) > 0 {
			fmt.Println("inside longest_path_to_tail  --> 11")
			fmt.Println("path to tail - after going to   ", path[len(path)-1] ,"  ---> ", pr.Path_h2_to_t)
			return []HelperModels.Position{path[len(path)-1]}
		}
	}

	fmt.Println("inside longest_path_to_tail  --> 12")
	return []HelperModels.Position{}

}


func (pr *PathRecommender) v2_shortest_path_to_tail() []HelperModels.Position {

	v_snake8 := pr.create_virtual_snake_from_orginal_snake()
	return v_snake8.bfs_from_head_to_tail()

}


func (pr *PathRecommender) V2_Set_path(fpos HelperModels.Position) []HelperModels.Position {

	pr.FoodPos = fpos
	pr.Path_h1_to_f = []HelperModels.Position{{}}
	pr.Path_h1_to_t = []HelperModels.Position{{}}
	pr.Path_h2_to_t = []HelperModels.Position{{}}

	neighbors := HelperFunctions.GetAvailableNeighborTiles(pr.OrginalSnake.Body[0], pr.MaxX, pr.MaxY, pr.OrginalSnake.Body, pr.FoodPos)

	if len(neighbors) > 0 {
		fmt.Println(" available neighbor exist")

		if pr.path_exist_to_food() && pr.can_reach_tail_after_eating_food() {

			fmt.Println("follow food")
			fmt.Println("full path to food --> ", pr.Path_h1_to_f)
			fmt.Println("next pos --> ", pr.Path_h1_to_f[0])
			fmt.Println("path to tail - after eating food --> ", pr.Path_h2_to_t)
			return pr.follow_food()

		} else {

			if Distance(pr.OrginalSnake.GetHeadPos(), pr.OrginalSnake.GetTailPos()) == 1 {

				fmt.Println("follow tail --> longest path")
				return pr.v2_longest_path_to_tail()

			} else {

				fmt.Println("follow tail --> shortest path")
				return pr.v2_shortest_path_to_tail()

			}

		}


		fmt.Println("option5")
		return []HelperModels.Position{}

	}else {
		fmt.Println("no available neighbor")
		return []HelperModels.Position{}
	}



}
