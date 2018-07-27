package majiong

import (
	"bytes"
	//"errors"
	//"fmt"
)

const USE_LAIZI = 1

//麻将牌
const (
	//东南西北
	MJ_DONGFENG = 1
	MJ_NANFENG  = 2
	MJ_XIFENG   = 3
	MJ_BEIFENG  = 4

	//中发白
	MJ_HONGZHONG = 5
	MJ_BAIBAN    = 6
	MJ_FACAI     = 7

	//空值
	MJ_NOTHING = 10

	//条
	MJ_YITIAO  = 13
	MJ_JIUTIAO = 21

	//万
	MJ_YIWAN  = 24
	MJ_JIUWAN = 32

	//筒
	MJ_YITONG  = 35
	MJ_JIUTONG = 43

	//赖
	MJ_ANYTHING = 53

	//
	MJ_MAX = 54
)

//在手牌中查找真子集
func ZiJi(pai []byte, paicnt int, sub []byte, subcnt int) int {
	//查找直接不够， 直接返回
	if paicnt < subcnt {
		return paicnt
	}
	cnt := paicnt
	for _, val := range sub {
		idx := bytes.IndexByte(pai, val)
		if idx >= 0 {
			cnt--
			pai[idx], pai[cnt] = pai[cnt], pai[idx]
			pai = pai[:cnt]
		} else {
			if USE_LAIZI == 1 {
				idx := bytes.IndexByte(pai, MJ_ANYTHING)
				if idx >= 0 {
					cnt--
					pai[idx], pai[cnt] = pai[cnt], pai[idx]
					pai = pai[:cnt]
				} else {
					return -1
				}
			}
		}
	}
	if cnt+subcnt == paicnt {
		return paicnt
	} else {
		return -1
	}
}

//检测胡牌
func IsHuPai(pai []byte, cnt int) bool {
	switch cnt {
	case 0:
		return true
	case 2:
		if pai[0] == pai[1] {
			return true
		} else {
			if pai[0] == MJ_ANYTHING {
				return true
			}
			if pai[1] == MJ_ANYTHING {
				return true
			}
		}
	}
	return false
}

func Loop(pai []byte, paicnt int, grp [][]byte, grpcnt int) int {
	sub := grp[grpcnt-1]
	subcnt := len(grp[grpcnt-1])
	ret := ZiJi(pai, paicnt, sub, subcnt)
	if ret-subcnt > 0 {
		data := pai[0 : ret-subcnt]
		ret = AllHu(data, len(data), grp, grpcnt)
	}
	return ret
}

func OneHu(pai []byte, paicnt int, grp [][]byte, grpcnt int) int {
	if grpcnt == 0 {
		return -1
	}
	ret := Loop(pai, paicnt, grp, grpcnt)
	if ret == -1 {
		ret = OneHu(pai, paicnt, grp, grpcnt-1)
	}
	return ret
}

func AllHu(pai []byte, paicnt int, grp [][]byte, grpcnt int) int {
	ret := -1
	if grpcnt != 0 {
		pcnt, gcnt := paicnt, grpcnt
		ret = OneHu(pai, pcnt, grp, gcnt)

		if IsHuPai(pai, ret) {
			return 0
		}
		ret = AllHu(pai, paicnt, grp, grpcnt-1)
	}
	return ret
}
