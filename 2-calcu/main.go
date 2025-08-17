package main

import (
	"app-3/bins"
	"fmt"
	"time"
)

func main() {

	bin1 := bins.CreateBin("hi", true, time.Now(), "mark")
	fmt.Println(*bin1)
	binList := bins.CreateBinList()
	binList.Bins_container[bin1.Id] = *bin1
	binList.Bins_container["second"] = *bin1

	fmt.Println(*binList)

}
