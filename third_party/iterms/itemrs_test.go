package iterms

import "testing"

func TestSetIterms2Config(t *testing.T) {
	err := SetIterms2Config()

	if err != nil {
		t.FailNow()
	}
}
