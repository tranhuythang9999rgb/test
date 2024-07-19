package utils

import (
	"sync"
	"time"

	"math/rand"
)

var (
	mu sync.Mutex
)

func GenerateUniqueKey() int64 {
	mu.Lock()
	defer mu.Unlock()

	var length = 7
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	key := int64(0)
	for i := 0; i < length; i++ {
		key = key*10 + int64(seededRand.Intn(9)) + 1
	}

	return key
}
func GenTimeStemp() int64 {
	time_now := time.Now()
	return time_now.Unix()
}
