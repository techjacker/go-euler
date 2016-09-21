package answers

import "testing"

func TestAnswer(t *testing.T) {
	res := Answer()
	expected := "hello"
	if res != expected {
		t.Errorf("wanted: %s\ngot: %s", expected, res)
	}
}
