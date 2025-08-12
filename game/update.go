package game

import (
	"fmt"
	"os"
	"time"
)

func (g *SnakeGame) Update() error {

    now := time.Now()

	if now.Sub(g.lastUpdate) >= g.MoveInterval {
		g.lastUpdate = now


        fmt.Println("***************")
        fmt.Println("|  snake body  --> \n",g.SnakeController.Body)

		g.FoodSpawner.SpawnFood(g.SnakeController.FoodEaten,g.SnakeController.Body)

        fmt.Println("food posisiton  --> ",g.FoodSpawner.FoodPos)


		path := g.path_recommender.V1_Set_path(g.FoodSpawner.FoodPos)

        fmt.Println("recommended path ",path[0])

		for _, body := range g.SnakeController.Body {
			if body == path[0] {
				fmt.Println("path in body")
				os.Exit(0)
			}
		}


        g.SnakeController.Move_one_step(path[0],g.FoodSpawner.FoodPos)
        fmt.Println("***************")




		
		fmt.Println("\n\n\n\n")

	}
	


	return nil

}
