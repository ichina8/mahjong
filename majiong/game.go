package majiong

import (
	"fmt"
)

const (
	PLAYER_MAX = 4
)

type game struct {
	desk Desk
}

func New(desk Desk) *game {
	Game := new(game)
	Game.desk = desk
	return Game
}

func (Game *game) InitDesk(playnum int) {
	Game.desk.PlayNum = playnum
	Game.desk.Player = make([]Player, PLAYER_MAX)
}

func (Game *game) InitHandPai(site int, handpai []byte) {
	Game.desk.Player[site].Handpai.Pai = make([]Pai, 0, HANDPAI_MAX)
	for _, val := range handpai {
		Game.desk.Player[site].Handpai.Pai = append(Game.desk.Player[site].Handpai.Pai, Pai{Val: val})
	}
}

func (Game *game) SetLaiZi(laizi byte) {
	Game.desk.LaiZi = laizi
}

func (Game *game) CheckHu(site int) {
	err := Game.desk.CheckHu(site)
	if err == nil {
		fmt.Println("胡牌成功", Game.desk.Player[site].HuType)
	} else {
		fmt.Println(err)
	}
}
