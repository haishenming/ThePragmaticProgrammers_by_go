package Tips

import "testing"

func TestTips(t *testing.T) {
	// 单元测试
	if i := Tips(1.0, 50); i != 0.5 {
		t.Error("Test False")
	} else {
		t.Log("Test OK")
	}

}


func BenchmarkTips(b *testing.B) {
	// 压力测试
	b.StartTimer()
	for i:=0; i<100000; i++ {
		Tips(1.5, 50)
	}
	}
