package bins

import "time"

type Bin struct {
	Id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	Bins_container map[string]Bin
}

func CreateBinList() *BinList {
	binlist := &BinList{
		Bins_container: make(map[string]Bin),
	}
	return binlist
}

func CreateBin(id string, private bool, createdAt time.Time, name string) *Bin {
	newBin := &Bin{
		Id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
	return newBin
}
