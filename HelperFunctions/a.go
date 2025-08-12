package HelperFunctions

import (
	"snake/HelperModels"
	"fmt"
)


func TileContainFood(tilePos  HelperModels.Position , foodPos HelperModels.Position) bool{
	if tilePos == foodPos {
		return true
	}
	return false
}


func TileContainSnakeBody(tilePos  HelperModels.Position , snake_body_parts []HelperModels.Position) bool{

	for _,body_part := range snake_body_parts {
		if tilePos == body_part {
			return true
		}
	}

	return false


}


func TileInBoard(tilePos  HelperModels.Position , MaxX int  ,  MaxY int) bool{

	return tilePos.X_Axis >= 0 && tilePos.X_Axis < MaxX&&
	tilePos.Y_Axis >= 0 && tilePos.Y_Axis < MaxY

}



func GenerateBoardTilePositionTemplate(MaxX int, MaxY int) map[HelperModels.Position]bool {
	//fmt.Println("HelperFunctions/MaxX",MaxX)
	//fmt.Println("HelperFunctions/MaxY",MaxY)
    grid := make(map[HelperModels.Position]bool) 

    for i := 0; i < MaxX; i++ {
        for j := 0; j < MaxY; j++ {
            grid[HelperModels.Position{X_Axis: i, Y_Axis: j}] = true 
        }
    }

	//fmt.Println("HelperFunctions/grid",grid)

    return grid
}



func  GenerateTileGrid( MaxX int  ,  MaxY int) []HelperModels.Position {
	grid := []HelperModels.Position{}

	for i := 0; i < MaxX; i++ {
		for j := 0; j < MaxY; j++ {
			grid = append(grid, HelperModels.Position{X_Axis: i, Y_Axis: j})
		}
	}

	return grid

}


func GetNeighborTiles(refrence_tile HelperModels.Position) []HelperModels.Position {
	return []HelperModels.Position{
		{X_Axis: refrence_tile.X_Axis + 1, Y_Axis: refrence_tile.Y_Axis},
		{X_Axis: refrence_tile.X_Axis - 1, Y_Axis: refrence_tile.Y_Axis},
		{X_Axis: refrence_tile.X_Axis, Y_Axis: refrence_tile.Y_Axis + 1},
		{X_Axis: refrence_tile.X_Axis, Y_Axis: refrence_tile.Y_Axis - 1},
	}
}



func GetAvailableNeighborTiles(ref HelperModels.Position,maxX, maxY int,snakeBody []HelperModels.Position,foodPos HelperModels.Position,) []HelperModels.Position {

    neighbors := GetNeighborTiles(ref)
    validNeighbors := []HelperModels.Position{}

    for _, n := range neighbors {
        if !TileInBoard(n, maxX, maxY) {
            fmt.Println(n, "--> out of grid")
            continue
        }

        if TileContainSnakeBody(n, snakeBody) {
        	fmt.Println(n, "--> in snake body")
            continue
        }
		/*
        if TileContainFood(n, foodPos) {
            fmt.Println(n, "--> in food tile")
            continue
        }
		*/
        fmt.Println(n, "--> valid neighbor")
        validNeighbors = append(validNeighbors, n)
    }

    return validNeighbors
}

