package util

import (
	"math/big"
	"math/rand"

	"github.com/google/uuid"
)

func ShortUID(length int) (uid string) {
	ubyte := []byte(uuid.New().String())
	idByte := new(big.Int).SetBytes(ubyte)
	rawID := idByte.String()
	start := rand.Intn(len(rawID) - length - 1)
	uid = rawID[start : length+start]
	return
}
