package utils

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/btcsuite/btcd/btcutil/base58"
	"time"
)

func UUID() string {
	milli := time.Now().UnixMilli()
	timeBuff := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBuff, uint64(milli))
	randBuff := make([]byte, 8)
	_, _ = rand.Read(randBuff)
	buff := append(timeBuff, randBuff...)
	return base58.Encode(buff)
}
