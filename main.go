package src

import (
	"fmt"
	"runtime"
	"LRU"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	lru := LRU.New()
	var buckets []*LRU.Bucket
	for i := 1; i <= 12; i++ {
		newBucket := &LRU.Bucket{Value: i}
		buckets = append(buckets, newBucket)
		lru.Put(newBucket)
	}
	lru.Get(buckets[2])   //检索bucket，bucket[2]会被标记为最近使用的bucket（放置到双向链表的链表头）
	lru.Put(&LRU.Bucket{Value: "newBucket"})   //插入新bucket，由于list已满，list末尾元素会被删除
	fmt.Print(lru)    //打印结果
}
