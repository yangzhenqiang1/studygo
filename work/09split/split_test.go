package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) { //单元测试的方法名 必须以Test起头 参数为*testing.T
	got := Split("a,b,c", ",")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expect:%v,got:%v", want, got)
	}
}

func BenchmarkSplit(t *testing.B) { //单元测试的方法名 必须以Test起头 参数为*testing.T
	got := Split("a,b,c", ",")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expect:%v,got:%v", want, got)
	}
}
