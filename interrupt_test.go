package lua

import (
	"context"
	"testing"
	"time"
)

func Benchmark_Test_loop_close(b *testing.B) {
	var s1 = `
	function f()
		while true do
		print("Hello World")
		end
	end
	f()
`
	var luaVM = NewState()
	ctx, cancel := context.WithCancel(context.Background())
	luaVM.SetContext(ctx)
	go func() {
		err := luaVM.DoString(s1)
		if err != nil {
			b.Fail()
		}
	}()
	time.Sleep(1 * time.Second)
	cancel()
	b.Log("luaVM.cancel()")
	luaVM.Close()
	b.Log("luaVM.Close()")
	time.Sleep(2 * time.Second)
}
func Test_loop_close(t *testing.T) {
	var s1 = `
		function f()
			while true do
			-- print("Hello World")
			end
		end
		f()
	`
	var luaVM = NewState()
	ctx, cancel := context.WithCancel(context.Background())
	luaVM.SetContext(ctx)
	go func() {
		err := luaVM.DoString(s1)
		if err != nil {
			t.Fail()
		}
	}()
	time.Sleep(1 * time.Second)
	cancel()
	t.Log("luaVM.cancel()")
	luaVM.Close()
	t.Log("luaVM.Close()")
	time.Sleep(2 * time.Second)
}
