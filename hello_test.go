package hello

import "testing"

func TestSayHello(t *testing.T) {

	subtest := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items:  []string{"Matt"},
			result: "Hello, Matt!",
		},
		{
			items:  []string{"Matt, Anne"},
			result: "Hello, Matt, Anne!",
		},
	}

	for _, st := range subtest {
		if s := SayHello(st.items); s != st.result {
			t.Errorf("wanted %s, got %s", st.result, s)
		}
	}

}
