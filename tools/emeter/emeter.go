package main

import (
	"fmt"
	"github.com/nik-johnson-net/go-smarthome"
	"os"
)

func main() {
	client := smarthome.NewClient(os.Args[1])

	info, err := client.SysInfo()
	if err != nil {
		panic(err)
	}

	if info.Children != nil {
		for _, child := range info.Children {
			meter, err := client.EMeter(child.ID)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s: %fW\n", child.Alias, meter.PowerW)
		}
	} else {
		meter, err := client.EMeter()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %fW\n", info.Alias, meter.PowerW)
	}
}
