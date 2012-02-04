// Copyright 2012 Markus Piéton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goguid is a GUID generation library based on noeqd / snowflake.
// All credit belongs to them!
//
// GUID generation and guarantees (GUID size = 64 bit)
//
//  - time            - 42 bit - millisecond precision (~ 100 years)
//  - machine id      - 10 bit - max. 1024 machines
//  - sequence number - 12 bit - rolls over after 4096 id's in one millisecond
//
// Usage
//   req_chan := make(chan chan int64)
//   quit := make(chan bool, 1)
//   
//   // Initialize the "server"
//   go ServGuid(0, 0, req_chan, quit)
//   
//   for i := 0; i < 10; i++ {
//       id := GetGUID(req_chan)
//       if id != 0 {
//         print(id)
//       }
//   }
//   
//   // Shut down the "server"
//   quit <- true
//
package guid

import (
	"time"
)

const (
	customEpoch = 1325376000000 // -> 2012-01-01 00:00:00 --> msec

	sequenceBits  = uint64(12)
	machineIdBits = uint64(10)

	machineIdShift = sequenceBits
	timestampShift = machineIdShift + machineIdBits

	sequenceMask = int64(-1) ^ (int64(-1) << sequenceBits)
)

var (
	lastTimestamp  int64
	machineId      int64
	sequenceNumber int64
)

// ServGuid should be used as a goroutine to serve the GUID.
func ServGuid(machineIdentifier, lastUsedTimestamp int, req <-chan chan int64, quit <-chan bool) {
	machineId = int64(machineIdentifier << machineIdShift)
	lastTimestamp = int64(lastUsedTimestamp)

	for {
		select {
		case receiver := <-req:
			receiver <- getNextGuid()
		case <-quit:
			return
		}
	}
	return
}

// GetGUID returns the next GUID.
// It requests a GUID over the chan that was initialized during ServGuid. 
func GetGUID(requestChannel chan<- chan int64) int64 {
	response := make(chan int64)
	requestChannel <- response

	return <-response
}

func customTimeInMilliseconds() int64 {
	return (time.Now().UnixNano() / 1e6)
}

func getNextGuid() int64 {
	timestamp := customTimeInMilliseconds()
	if lastTimestamp == timestamp {
		sequenceNumber = (sequenceNumber + 1) & sequenceMask
		if sequenceNumber == 0 {
			// sequence wrapped ... wait for the next millisecond
			for timestamp <= lastTimestamp {
				timestamp = customTimeInMilliseconds()
			}
		}
	} else {
		sequenceNumber = 0
	}

	if timestamp < lastTimestamp {
		return 0
	}
	lastTimestamp = timestamp

	id := (timestamp << timestampShift) | machineId | sequenceNumber

	return id
}
