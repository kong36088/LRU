package LRU

import (
	"container/list"
	"fmt"
)

const capacity = 10

type LRU struct {
	bucketMap map[*Bucket]*list.Element
	list      *list.List
	cap       int //容量
}

func New() *LRU {
	return new(LRU).Init()
}

//初始化
func (lru *LRU) Init() *LRU {
	lru.bucketMap = make(map[*Bucket]*list.Element)
	lru.list = list.New()
	lru.cap = capacity
	return lru
}

//插入
func (lru *LRU) Put(bucket *Bucket) {
	if length := lru.list.Len(); length >= lru.cap {
		backEle := lru.list.Back()
		delete(lru.bucketMap, backEle.Value.(*Bucket))
		lru.list.Remove(backEle)
	}

	var ele *list.Element
	if lru.list.Front() == nil {
		ele = lru.list.PushBack(bucket)
	} else {
		ele = lru.list.InsertBefore(bucket, lru.list.Front())
	}
	lru.bucketMap[bucket] = ele
}

//获取
func (lru *LRU) Get(bucket *Bucket) (*Bucket, bool) {
	v, ok := lru.bucketMap[bucket]
	if ! ok {
		return nil, ok
	}
	lru.list.MoveToFront(v)
	return v.Value.(*Bucket), true
}

//判断是否包含某个Bucket
func (lru *LRU) ContainsKey(bucket *Bucket) bool {
	_, ok := lru.bucketMap[bucket]
	return ok
}

//获取元素个数
func (lru *LRU) Len() int{
	return lru.list.Len()
}

func (lru *LRU) Empty() bool{
	return lru.Len() == 0
}

//打印内容
func (lru *LRU) String() string {
	var result string
	result += "list : "
	for e := lru.list.Front(); e != nil; e = e.Next() {
		result += fmt.Sprint(e.Value.(*Bucket)) + " "
	}

	result += "\nmap :\n"
	for k, v := range lru.bucketMap {
		result += fmt.Sprintf("key[%s] value[%s]\n", k, v.Value.(*Bucket))
	}
	return result
}

func main() {
	l := New()
	_, ok := l.Get(&Bucket{123})
	print(ok)
	return
}
