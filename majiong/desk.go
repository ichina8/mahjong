package majiong

import (
	//"bytes"
	"errors"
	"fmt"
	"time"
)

const HANDPAI_MAX = 14

type Pai struct {
	Val byte
	//Add Other Element
	//Idx int8u
}

type Handpai struct {
	Pai []Pai
}

type Player struct {
	Handpai Handpai
	//Add Other Element
	HuPai  byte
	HuType int
}

type Desk struct {
	Player  []Player
	PlayNum int
	LaiZi   byte
	//Add Other Element
}

func (desk *Desk) GetHandPaiList(Site int, Flag bool) (bool, []byte) {
	data := make([]byte, 0)
	hasLaizi := false
	for _, pai := range desk.Player[Site].Handpai.Pai {
		if Flag && desk.LaiZi != 0 && pai.Val == desk.LaiZi {
			data = append(data, MJ_ANYTHING)
			hasLaizi = true
		} else {
			data = append(data, byte(pai.Val))
		}
	}
	return hasLaizi, data
}

func (desk *Desk) CheckHu(Site int) error {
	t := time.Now()
	desk.Player[Site].HuType |= new(hu23333).HuPai(desk, Site)
	desk.Player[Site].HuType |= new(huPPH).HuPai(desk, Site)
	desk.Player[Site].HuType |= new(huQD).HuPai(desk, Site)
	desk.Player[Site].HuType |= new(hu13Y).HuPai(desk, Site)
	elapsed := time.Since(t)
	fmt.Println("共消耗时间: ", elapsed)

	if desk.Player[Site].HuType > 0 {
		desk.Player[Site].HuType |= new(isJZH).HuPai(desk, Site)
		desk.Player[Site].HuType |= new(isYTL).HuPai(desk, Site)
		desk.Player[Site].HuType |= new(isBKD).HuPai(desk, Site)
		return nil
	} else {
		return errors.New("胡牌失败")
	}
}

////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
//胡牌接口(进一步封装)
type huPai interface {
	HuPai(desk *Desk, Site int) int
}

//见字胡
type isJZH struct {
}

//边卡吊
type isBKD struct {
}

//一条龙
type isYTL struct {
}

//清一色
type isQYS struct {
}

//进一步封装
func (Hu23333 *hu23333) HuPai(desk *Desk, Site int) int {
	var ret = 0
	hasLaizi, PaiList := desk.GetHandPaiList(Site, true)
	ret |= Hu23333.Hu(PaiList)
	if ret > 0 {
		if hasLaizi {
			_, PaiList = desk.GetHandPaiList(Site, false)
			ret1 := Hu23333.Hu(PaiList)
			if ret1 > 0 {
				ret |= HFO_YH
			}
		} else {
			ret |= HFO_YH
		}
	}
	return ret
}

func (HuPPH *huPPH) HuPai(desk *Desk, Site int) int {

	return 0
}

func (HuQD *huQD) HuPai(desk *Desk, Site int) int {

	return 0
}

func (Hu13Y *hu13Y) HuPai(desk *Desk, Site int) int {

	return 0
}

func (IsYTL *isYTL) HuPai(desk *Desk, Site int) int {
	flag := []byte{0, 0, 0}
	_, PaiList := desk.GetHandPaiList(Site, true)
	data := make([]byte, MJ_MAX)
	for _, paiVal := range PaiList {
		data[paiVal]++
	}
	for i := MJ_YITIAO; i <= MJ_JIUTIAO; i++ {
		if data[i] > 0 {
			flag[0]++
		}
	}
	for i := MJ_YIWAN; i <= MJ_JIUWAN; i++ {
		if data[i] > 0 {
			flag[1]++
		}
	}
	for i := MJ_YITONG; i <= MJ_JIUTONG; i++ {
		if data[i] > 0 {
			flag[2]++
		}
	}

	//检测判断一条龙
	for i := 0; i < 3; i++ {
		if flag[i]+data[MJ_ANYTHING] >= 9 {
			var sub = make([]byte, 9)
			for j := 0; j < 9; j++ {
				sub[j] = MJ_YITIAO + byte(i*11)
			}
			cnt := ZiJi(PaiList, len(PaiList), sub, len(sub))
			if cnt-len(sub) > 0 {
				subList := PaiList[0 : cnt-len(sub)]
				ret := new(hu23333).Hu(subList)
				if ret > 0 {
					return HFO_YTL
				}
			}
		}
	}
	return 0
}

func (IsBKD *isBKD) HuPai(desk *Desk, Site int) int {
	return 0
}

func (IsJZH *isJZH) HuPai(desk *Desk, Site int) int {
	return 0
}
