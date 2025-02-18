package util

import (
	"testing"
)

func TestGenerateRandomKey(t *testing.T) {
	key, err := GenerateRandomKey(10)
	if err != nil {
		panic(err)
	}
	t.Logf("%s", key)
}
