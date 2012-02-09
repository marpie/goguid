package main

import (
	"fmt"
	guid "github.com/marpie/goguid"
	"time"
)

const (
	max_guids = int64(1e9)
)

func main() {
	fmt.Printf("GUID Benchmark\n--------------\n\n%d GUIDs ... way to go...", max_guids)

	// Initialize the package
	guid.InitGUID(0, 0)

	tick := time.Now().Unix()
	counter := int64(0)
	for i := int64(0); i < max_guids; i++ {
		if guid.GetGUID() != 0 {
			counter++
		}
	}
	tick = time.Now().Unix() - tick

	// get the GUIDs per sec...
	if tick > 0 {
		tick = counter / tick
	}

	fmt.Printf("\n%d valid GUIDs generated.\n", counter)
	fmt.Printf("%d/sec.\nDone.", tick)

	return
}
