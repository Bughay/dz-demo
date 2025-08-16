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

	bin1 := createBin("5", true, time.Now(), "mark")
	fmt.Println(bin1)

}
