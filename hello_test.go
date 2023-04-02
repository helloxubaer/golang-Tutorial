package hello

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello tes!"
	got := SayHello("test")

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
