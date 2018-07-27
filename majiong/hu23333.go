package majiong

import (
	//"bytes"
	//"errors"
	"fmt"
)

const (
	HFO_YH  = 0x00000001
	HFO_PH  = 0x00000002
	HFO_PPH = 0x00000004
	HFO_QD  = 0x00000008
	HFO_13Y = 0x00000010

	HFO_YTL = 0x00000100
	HFO_BKD = 0x00000200
)

//胡牌接口
type hu interface {
	Hu(HandPai []byte) int
}

//胡牌
type hu23333 struct {
}

//碰碰胡
type huPPH struct {
}

//七对
type huQD struct {
}

//13幺
type hu13Y struct {
}

func (Hu23333 *hu23333) Hu(HandPai []byte) int {
	PaiList := make([]byte, len(HandPai))
	copy(PaiList, HandPai)
	//生成字典
	data := make([]byte, MJ_MAX)
	for _, paiVal := range PaiList {
		data[paiVal]++
	}
	//fmt.Println(data)
	grpDirc := new(DircPH).GenDirc(data)
	fmt.Println(len(grpDirc), grpDirc)
	//判断胡牌
	ret := AllHu(PaiList, len(PaiList), grpDirc, len(grpDirc))

	if ret == 0 {
		return HFO_PH
	} else {
		return 0
	}
}

func (HuPPH *huPPH) Hu(HandPai []byte) int {
	return 0
}

func (HuQD *huQD) Hu(HandPai []byte) int {
	return 0
}

func (Hu13Y *hu13Y) Hu(HandPai []byte) int {

	return 0
}

////////////////////////////////////////////////////////////
