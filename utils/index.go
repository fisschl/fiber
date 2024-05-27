package utils

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/gookit/goutil/envutil"
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

func LoadENV() error {
	config.AddDriver(yaml.Driver)
	err := config.LoadExists("config.yaml")
	if err != nil {
		return err
	}
	err = config.LoadData(envutil.EnvMap())
	return err
}
