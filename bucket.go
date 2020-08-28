package token_bucket

import (
	"sync"
	"time"
)

type Bucket struct {
	thresholdValue  int64
	fillInterval    int64 // fillInterval holds the interval between each tick.
	availableTokens int64
	putSpeed        int64
	latestTick      int64
	mu              sync.Mutex
}

// Rate return every token needs how much Duration
func (bucket *Bucket) Rate() float64 {
	return 1e9 / float64(bucket.fillInterval)
}

func (bucket *Bucket) available(now time.Time) int64 {
	tick := now.UnixNano()
	bucket.mu.Lock()
	defer bucket.mu.Unlock()
	lastTick := bucket.latestTick
	bucket.latestTick = tick
	if bucket.availableTokens >= bucket.thresholdValue {
		bucket.availableTokens = bucket.thresholdValue
		return bucket.availableTokens
	}
	bucket.availableTokens += (tick - lastTick) / bucket.fillInterval
	if bucket.availableTokens > bucket.thresholdValue {
		bucket.availableTokens = bucket.thresholdValue
	}
	return bucket.availableTokens
}

func (bucket *Bucket) Take(tokenCount int64) (int64, bool) {
	available := bucket.available(time.Now())
	if available > tokenCount {
		bucket.availableTokens = available - tokenCount
		return tokenCount, true
	} else {
		bucket.availableTokens = 0
		return available, false
	}
}

func NewTokenBucket(thresholdValue, putSpeed int64) *Bucket {
	bucket := &Bucket{
		latestTick:      time.Now().UnixNano(),
		fillInterval:    1e9 / putSpeed,
		availableTokens: 0,
		thresholdValue:  thresholdValue,
		putSpeed:        putSpeed}
	return bucket
}
