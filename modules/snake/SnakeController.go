package snake

import (
	"image/color"
	"snake/HelperFunctions"
	"snake/HelperModels"
	"snake/modules/canvas"
)

type SnakeController struct {
    Body         []HelperModels.Position
    HeadColor    color.Color
    BodyColor    color.Color
    painter      *canvas.Painter
	Alive        bool
    Score        int
    MovesWithoutEating int
    MAX_MOVES_WITHOUT_EATING int
    FoodEaten bool

}


func NewSnake(body []HelperModels.Position, head_color color.Color , body_color color.Color, paint *canvas.Painter ) *SnakeController{
    s := &SnakeController{
        Body:         body,
		HeadColor:    head_color,
        BodyColor:    body_color,
		Alive:        true,
        Score:        0,
        painter: paint,
        FoodEaten: false,

    }

    
    s.MAX_MOVES_WITHOUT_EATING = 17 *17 *2

    return s
}



func (s *SnakeController) Move_one_step(newHead HelperModels.Position , food HelperModels.Position) {
    //fmt.Println("modules/snake/SnakeController.go/Move_one_step/s.Body",s.Body)
    if HelperFunctions.TileContainFood(newHead,food) {
        s.FoodEaten = true
        s.moveAndGrow(newHead)
    } else {
        s.FoodEaten = false 
        s.moveWithoutGrow(newHead)
    }
    //fmt.Println("modules/snake/SnakeController.go/Move_one_step/s.Body",s.Body)
}



func (s *SnakeController) moveAndGrow(newHead HelperModels.Position) {

    s.Body = append([]HelperModels.Position{newHead},s.Body... )
    s.Score++
    s.MovesWithoutEating = 0
}

func (s *SnakeController) moveWithoutGrow(newHead HelperModels.Position) {
    s.Body = append([]HelperModels.Position{newHead},s.Body... )
    s.Body = s.Body[:len(s.Body)-1 ]
    s.MovesWithoutEating++

}

func (vs *SnakeController) GetHeadPos() HelperModels.Position{
	return vs.Body[0]
}



func (vs *SnakeController) GetTailPos() HelperModels.Position{
	return vs.Body[len(vs.Body)-1]
}





func (s *SnakeController) Draw() {
    for i, pos := range s.Body {
        if i == 0 {
            s.painter.FillTileCentered(pos, s.HeadColor)
        } else {
            s.painter.FillTileCentered(pos, s.BodyColor) 
        }
    }
}


