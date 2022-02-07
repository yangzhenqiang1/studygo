package implode

import (
	"reflect"
	"testing"
)

func TestImplode(t *testing.T) {
	got := Implode([]string{"a", "b", "c"}, ":")
	want := "a:b:c"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v,want:%v", got, want)
	}
}

func BenchmarkImplode(b *testing.B) {
	got := Implode([]string{"a", "b", "c"}, ":")
	want := "a:b:c"
	if !reflect.DeepEqual(got, want) {
		b.Errorf("got:%v,want:%v", got, want)
	}
}
