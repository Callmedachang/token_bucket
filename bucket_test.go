package token_bucket

import (
	"log"
	"testing"
	"time"
)

func TestNewTokenBucket(t *testing.T) {
	/*
		init TokenBucket I called myBucket
		she will put 2 token every second
	*/
	myBucket := NewTokenBucket(5, 2)
	for true {
		log.Println(myBucket.Take(1))
		time.Sleep(time.Second)
	}
}
func TestNewTokenBucket2(t *testing.T) {
	s := int64(1)
	for {
		s = nextQuantum(s)
		log.Println(s)
		time.Sleep(time.Second)
	}
}
func nextQuantum(q int64) int64 {
	q1 := q * 11 / 10
	if q1 == q {
		q1++
	}
	return q1
}
