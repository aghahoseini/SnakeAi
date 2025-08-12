package food

import (
	"image/color"
	"math/rand"
	"snake/HelperFunctions"
	"snake/HelperModels"
	"snake/modules/canvas"
	"fmt"
)

type FoodSpawner struct {
	FoodPos                  HelperModels.Position
	painter                  *canvas.Painter
	FreeBoardTileTemplate    map[HelperModels.Position]bool
	MaxX                     int
	MaxY                     int
}

func NewFoodSpawner(paint *canvas.Painter , maxx int , maxy int ) *FoodSpawner {
	FoodSpawner := &FoodSpawner{}
	FoodSpawner.painter = paint
	FoodSpawner.MaxX = maxx
	FoodSpawner.MaxY = maxy

	FoodSpawner.FreeBoardTileTemplate = HelperFunctions.GenerateBoardTilePositionTemplate(FoodSpawner.MaxX,FoodSpawner.MaxY)

	return  FoodSpawner



	

}

func (fs *FoodSpawner) GenerateFreeTileList(snake_body []HelperModels.Position) []HelperModels.Position {

	freetiletemplate := make(map[HelperModels.Position]bool)
    for k, v := range fs.FreeBoardTileTemplate {
        freetiletemplate[k] = v
    }

    freetiles := []HelperModels.Position{}

    for _, body_part := range snake_body {
        delete(freetiletemplate, body_part)
    }

    for key := range freetiletemplate {
        freetiles = append(freetiles, key)
    }

    return freetiles
}

func (fs *FoodSpawner) SpawnFood(foodeaten bool , snake_body []HelperModels.Position) HelperModels.Position {


	if foodeaten {
		freeTiles := fs.GenerateFreeTileList(snake_body)
		fmt.Println("foodspawner/fs.FreeBoardTileTemplate",len(fs.FreeBoardTileTemplate))
		fmt.Println("foodspawner/len(freeTiles)",len(freeTiles))
    	idx := rand.Intn(len(freeTiles))

		fs.FoodPos = freeTiles[idx]
    	return freeTiles[idx]

	}

	return HelperModels.Position{X_Axis: -1 , Y_Axis: -1}


}




func (fs *FoodSpawner) Draw() {
	fs.painter.FillTile(fs.FoodPos, color.RGBA{255, 99, 71, 255})
}
