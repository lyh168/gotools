package randomUtils

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	numberRand     = rand.New(rand.NewSource(time.Now().UnixNano()))
	numberRandLock sync.Mutex
)

func NumericSequence(length int) string {
	if length <= 0 {
		return ""
	}
	bs := make([]byte, 0, length)
	numberRandLock.Lock()
	for i := 0; i < length; i++ {
		bs = strconv.AppendInt(bs, numberRand.Int63n(10), 10)
	}
	numberRandLock.Unlock()
	return string(bs)
}
