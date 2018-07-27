// HuPai project main.go
package main

import (
	"HuPai/majiong"
	//"fmt"
	"time"
)

func InitPai() []byte {
	handpai := []byte{
		majiong.MJ_YITIAO + 0, majiong.MJ_YITIAO + 0, majiong.MJ_YITIAO + 0,
		majiong.MJ_YITIAO + 1, majiong.MJ_YITIAO + 2, majiong.MJ_YITIAO + 3,
		majiong.MJ_YITIAO + 4, majiong.MJ_YITIAO + 5, majiong.MJ_YITIAO + 6,
		majiong.MJ_YITIAO + 7, majiong.MJ_YITIAO + 8, majiong.MJ_YITIAO + 3,
		majiong.MJ_BEIFENG, majiong.MJ_BEIFENG}
	return handpai
}

func main() {
	var desk majiong.Desk
	mj := majiong.New(desk)
	mj.InitDesk(majiong.PLAYER_MAX)
	mj.SetLaiZi(majiong.MJ_YIWAN)
	handpai := InitPai()
	for i := 0; i < majiong.PLAYER_MAX; i++ {
		mj.InitHandPai(i, handpai)
		go mj.CheckHu(i)
	}

	time.Sleep(time.Second)
}
