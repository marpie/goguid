package main

import (
	"fmt"
	guid "goguid"
	"time"
)

const (
  max_guids = int64(1e6)
)

func main() {
	req_chan := make(chan chan int64)
	quit := make(chan bool, 1)

  fmt.Printf("GUID Benchmark\n--------------\n\n%d GUIDs ... way to go...", max_guids)

	// Initialize the "server"
	go guid.ServGuid(0, 0, req_chan, quit)

  tick := time.Now().Unix()
  counter := int64(0)
  for i := int64(0); i < max_guids; i++ {
    if guid.GetGUID(req_chan) != 0 {
      counter++
    }
  }
  tick = time.Now().Unix()-tick

	// Shut down the "server"
	quit <- true

  // get the GUIDs per sec...
  if tick > 0 {
    tick = counter / tick
  }

	fmt.Printf("\n%d valid GUIDs generated.\n", counter)
	fmt.Printf("%d/sec.\nDone.", tick)

	return
}
