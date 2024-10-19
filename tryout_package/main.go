package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/sony/sonyflake"
)

func main() {

	fmt.Println(net.InterfaceAddrs())

	fmt.Printf("%b\n", 132)

	sf_settings := sonyflake.Settings{
		StartTime:      time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC),
		MachineID:      func() (uint16, error) { return uint16(132), nil },
		CheckMachineID: nil}

	sf, err := sonyflake.New(sf_settings)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ids := []uint64{}
	for i := 300; i >= 0; i-- {
		k, _ := sf.NextID()
		ids = append(ids, k)
		fmt.Println(k)
		fmt.Println(sonyflake.Decompose(k))
		fmt.Printf("%b\n", k)
		fmt.Println("--------------------------------")
	}

	seen := make(map[uint64]int)

	for j, i := range ids {
		fmt.Println(j, " ", i)
		if k, e := seen[i]; e {
			fmt.Println("seen ", i, " at ", k, " and ", j)
		}
	}

	fmt.Println("complete")
}
