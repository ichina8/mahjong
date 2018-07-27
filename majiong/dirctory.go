package majiong

//"bytes"
//"errors"
//"fmt"

//
type Dirc interface {
	GenDirc(Pai []byte) [][]byte
}

type DircPH struct {
}

type DircQD struct {
}

func (dircPH *DircPH) GenDirc(Pai []byte) [][]byte {
	var i, j byte
	var Data [][]byte
	if USE_LAIZI == 1 {
		Kezi := []byte{MJ_ANYTHING, MJ_ANYTHING, MJ_ANYTHING}
		Data = append(Data, Kezi)
	}

	//加入刻字/顺子<利用赖子充当>
	for i = MJ_DONGFENG; i <= MJ_JIUTONG; i++ {
		if Pai[i] > 0 && Pai[i] < 3 && Pai[i]+Pai[MJ_ANYTHING] >= 3 {
			Kezi := []byte{i, i, i}
			Data = append(Data, Kezi)
		}

		if i >= MJ_YITIAO {
			var count byte = 0
			for j = 0; j < 3; j++ {
				if Pai[int(i+j)] > 0 {
					count++
				}
			}
			if count > 0 && count < 3 && count+Pai[MJ_ANYTHING] >= 3 {
				ShunZi := []byte{i, i + 1, i + 2}
				Data = append(Data, ShunZi)
			}
		}
	}

	//加入刻字/顺子<不用赖充当>
	for i = MJ_DONGFENG; i <= MJ_JIUTONG; i++ {
		if Pai[i] >= 3 {
			Kezi := []byte{i, i, i}
			Data = append(Data, Kezi)
		}
		if i >= MJ_YITIAO {
			var count byte = 0
			for j = 0; j < 3; j++ {
				if Pai[int(i+j)] > 0 {
					count++
				}
			}
			if count >= 3 {
				ShunZi := []byte{i, i + 1, i + 2}
				Data = append(Data, ShunZi)
			}
		}
	}
	return Data
}

//
func (dircPH *DircQD) GenDirc(Pai []byte) [][]byte {
	var Data [][]byte
	//to do
	return Data

}
