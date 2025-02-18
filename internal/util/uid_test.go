package util

import "testing"

func TestShortUID(t *testing.T) {
	id := ShortUID(12)
	t.Log(id)
}
