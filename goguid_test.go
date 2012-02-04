package guid

import "testing"

func TestGetNextGuid(t *testing.T) {
	machineId = int64(0 << machineIdShift)
	lastTimestamp = int64(0)

	for i := 0; i < 10; i++ {
		if getNextGuid() == 0 {
			t.Errorf("GUID generation failed!")
		}
	}
}

func testGuidServer(t *testing.T) {
	req_chan := make(chan chan int64)
	quit := make(chan bool, 1)

	go ServGuid(0, 0, req_chan, quit)

	for i := 0; i < 10; i++ {
		id := GetGUID(req_chan)
		if id == 0 {
			t.Errorf("failed.")
		}
	}

	quit <- true

}
