package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins map[string]Bin
}

func createBinList() *BinList {
	binlist := &BinList{
		bins: make(map[string]Bin),
	}
	return binlist
}

func createBin(id string, private bool, createdAt time.Time, name string) *Bin {
	newBin := &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
	return newBin
}
func main() {

	bin1 := createBin("hi", true, time.Now(), "mark")
	fmt.Println(*bin1)
	binList := createBinList()
	binList.bins[bin1.id] = *bin1
	binList.bins["second"] = *bin1

	fmt.Println(*binList)

}
