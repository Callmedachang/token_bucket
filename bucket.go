package main

import (
	"time"
)

type Bucket struct {
	Length              int
	ThresholdValue      int
	MinusThresholdValue int
	TokenList           int
	PutSpeed            int
}

/*
this function Put the token into the bucket at a certain rate
*/
func (bucket *Bucket) PushToken(putNum int) {
	for i := 0; i < putNum; i++ {
		if bucket.Length < bucket.ThresholdValue {
			bucket.Length++
			bucket.TokenList++
		}
	}
}

/*
request need use this function to get  licence for the services
true : means that this request could get the services
false: the reverse
*/
func (bucket *Bucket) GetToken(needNum int) bool {
	if bucket.Length > needNum {
		bucket.Length -= needNum
		bucket.TokenList -= needNum
		return true
	} else {
		if bucket.MinusThresholdValue > needNum {
			bucket.Length -= needNum
			bucket.TokenList -= needNum
			return true
		} else {
			return false
		}
	}
}

/*
Put a fixed number of tokens into the bucket every second
*/
func (bucket *Bucket) Start() {
	go func() {
		for true {
			bucket.PushToken(bucket.PutSpeed)
			time.Sleep(time.Second)
		}
	}()
}
func NewTokenBucket(thresholdValue, minusThresholdValue, putSpeed int) *Bucket {
	return &Bucket{ThresholdValue: thresholdValue, MinusThresholdValue: minusThresholdValue, PutSpeed: putSpeed}
}
func main() {
	/*
	init TokenBucket I called myBucket
	she will put 2 token every second
	*/
	myBucket := NewTokenBucket(5, 0, 2)
	myBucket.Start()
	for true {
		myBucket.GetToken(1)
		println(myBucket.TokenList)
		println(myBucket.Length)
		time.Sleep(time.Second)
	}
}
