package guid

import (
	"testing"
)

func BenchmarkGetGUID(b *testing.B) {
	b.StopTimer()
	InitGUID(0, 0)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if GetGUID() == 0 {
			b.Errorf("GUID generation failed!")
		}
	}
}

func TestGetGUID(t *testing.T) {
	InitGUID(0, 0)

	if GetGUID() == 0 {
		t.Errorf("GUID generation failed!")
	}
}
