package pathrecommender

import (
	"fmt"
	"snake/HelperFunctions"
	"snake/HelperModels"
)


type VirtualSnake struct {
	Body []HelperModels.Position
	FoodPos HelperModels.Position
	MaxX int
	MaxY int

}


func (vs *VirtualSnake) GetHeadPos() HelperModels.Position{
	return vs.Body[0]
}

func (vs *VirtualSnake) GetTailPos() HelperModels.Position{
	return vs.Body[len(vs.Body)-1]
}

func (vs *VirtualSnake) DeleteTail() {
	if len(vs.Body) > 0 {
		vs.Body = vs.Body[:len(vs.Body)-1]
	}
}

func (vs *VirtualSnake) AddTail(pos HelperModels.Position) {
	vs.Body = append(vs.Body, pos)
}


func (vs *VirtualSnake) IsTileInMyBody(tile HelperModels.Position) bool{

	for _,pos := range vs.Body {
		if pos == tile {
			return true
		}
	}
	return false

}



func (vs *VirtualSnake) move_one_step(pos HelperModels.Position) {



	if HelperFunctions.TileContainFood(pos , vs.FoodPos ) {
	    vs.Body = append([]HelperModels.Position{pos}, vs.Body...) 
	}else {
	    vs.Body = append([]HelperModels.Position{pos}, vs.Body...) 
		vs.Body = vs.Body[:len(vs.Body)-1]

    }

	

}




func (vs *VirtualSnake) bfs(start HelperModels.Position, end HelperModels.Position) []HelperModels.Position {

	//fmt.Println("PathRecommender/a.go/VirtualSnake/bfs/start",start)
	//fmt.Println("PathRecommender/a.go/VirtualSnake/bfs/end",end)


	queue := []HelperModels.Position{}

	queue = append(queue, start)
	
	visited := map[HelperModels.Position]bool{}

	grid := HelperFunctions.GenerateTileGrid(vs.MaxX,vs.MaxY)
	//fmt.Println("PathRecommender/a.go/VirtualSnake/bfs/Grid",grid)

	for _, pos := range grid {
		visited[pos] = false
	}

	visited[start] = true

	prev := map[HelperModels.Position]interface{}{}
	for _, pos := range grid {
		prev[pos] = nil
	}


	for len(queue) > 0 {


		node := queue[0]
		queue = queue[1:]

		neighbors := HelperFunctions.GetNeighborTiles(node)


		for _, next_node := range neighbors {

			if HelperFunctions.TileInBoard(next_node,vs.MaxX,vs.MaxY) && !vs.IsTileInMyBody(next_node) &&  !visited[next_node] {
				

				queue = append(queue, next_node)
				visited[next_node] = true
				prev[next_node] = node			

			}

		}


	}




	path := []HelperModels.Position{}

	p_node := end



	start_node_found := false


	for !start_node_found {


		if prev[p_node] == nil {
			return []HelperModels.Position{}
		}

		p_node = prev[p_node].(HelperModels.Position)


		if p_node == start {
			path = append(path, end)

			return  path

		}
		path = append([]HelperModels.Position{p_node}, path...)

	}

			

	return []HelperModels.Position{}


}

func (vs *VirtualSnake) bfs_from_head_to_food() []HelperModels.Position {

	//fmt.Println("PathRecommender/a.go/VirtualSnake/bfs_from_head_to_food/vs.GetHeadPos()",vs.GetHeadPos())
	//fmt.Println("PathRecommender/a.go/VirtualSnake/bfs_from_head_to_food/vs.FoodPos",vs.FoodPos)


	abc := vs.bfs(vs.GetHeadPos() , vs.FoodPos)
	return abc

}


func (vs *VirtualSnake) bfs_from_head_to_tail() []HelperModels.Position {

	fmt.Println("pathrecommender/set_path/bfs_from_head_to_tail/vs.Body",vs.Body)
	tail_pos := vs.GetTailPos()
	vs.DeleteTail()
	fmt.Println("pathrecommender/set_path/bfs_from_head_to_tail/vs.Body after delete tail",vs.Body)
	abc := vs.bfs(vs.GetHeadPos() , tail_pos)
	vs.AddTail(tail_pos)
	fmt.Println("pathrecommender/set_path/bfs_from_head_to_tail/vs.Body after add tail",vs.Body)
	fmt.Println("pathrecommender/set_path/bfs_from_head_to_tail/return path from bfs",abc)
	return abc



}


func Distance(pos1, pos2 HelperModels.Position) int {
	dx := pos2.X_Axis - pos1.X_Axis
	dy := pos2.Y_Axis - pos1.Y_Axis

	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}
