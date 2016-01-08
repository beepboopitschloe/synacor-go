package machine

import (
	"testing"
)

func TestPushAndPopStack(t *testing.T) {
	m := NewMachine()

	test_length := uint16(32768)

	expected := make([]uint16, test_length)

	for i := uint16(0); i < test_length; i++ {
		expected[i] = i * 4

		m.PushStack(expected[i])
	}

	for i := test_length - 1; i > 0; i-- {
		result := m.PopStack()

		if result != expected[i] {
			t.Fatalf("expected[%d] should have been %s, got %s\n", i, expected[i], result)
		}
	}
}
