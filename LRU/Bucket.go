package LRU

import (
	"fmt"
)

type Bucket struct {
	Value interface{}
}

//打印bucket
func (bucket *Bucket) String() string {
	var result string
	result = fmt.Sprint(bucket.Value)
	return result
}