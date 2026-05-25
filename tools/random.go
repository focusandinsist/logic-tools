package random

import (
	"math/rand"
	"time"
)

// ShuffleQueue 洗牌队列
type ShuffleQueue struct {
	origin  []int32
	shuffle []int32
	ptr     int
	r       *rand.Rand
}

// NewShuffleQueue 使用指定数组创建洗牌队列
func NewShuffleQueue(arr []int32) *ShuffleQueue {
	src := make([]int32, len(arr))
	copy(src, arr)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	q := &ShuffleQueue{
		origin: src,
		r:      r,
	}
	q.Reset()
	return q
}

// NewShuffleQueueRange 创建 1~limit 的洗牌队列
func NewShuffleQueueRange(limit int32) *ShuffleQueue {
	arr := make([]int32, limit)
	for i := range arr {
		arr[i] = int32(i) + 1
	}
	return NewShuffleQueue(arr)
}

// NewShuffleQueueRangeN 从 min 到 max 中抽出 n 个卡牌形成牌堆
func NewShuffleQueueRangeN(min, max, n int32) *ShuffleQueue {
	if min > max {
		panic("invalid range: min > max")
	}
	if n <= 0 {
		panic("invalid n: must > 0")
	}
	total := max - min + 1
	if n > total {
		panic("invalid n: exceed range size")
	}
	arr := make([]int32, total)
	for i := range arr {
		arr[i] = min + int32(i)
	}
	q := NewShuffleQueue(arr)
	picked := append([]int32(nil), q.shuffle[:n]...) // 洗牌后取前 n 个
	return NewShuffleQueue(picked)
}

// Reset 重新洗牌重置
func (q *ShuffleQueue) Reset() {
	q.shuffle = make([]int32, len(q.origin))
	copy(q.shuffle, q.origin)
	q.r.Shuffle(len(q.shuffle), func(i, j int) {
		q.shuffle[i], q.shuffle[j] = q.shuffle[j], q.shuffle[i]
	})
	q.ptr = 0
}

// Next 取牌，取完自动重新洗牌新一轮
func (q *ShuffleQueue) Next() int32 {
	if q.ptr >= len(q.shuffle) {
		q.Reset()
	}
	val := q.shuffle[q.ptr]
	q.ptr++
	return val
}

// GetPtr 获取游标,即下一张牌的索引
func (q *ShuffleQueue) GetPtr() int32 {
	return int32(q.ptr)
}

// GetShuffle 获取当前轮洗牌后的完整牌序（返回副本）
func (q *ShuffleQueue) GetShuffle() []int32 {
	return append([]int32{}, q.shuffle...)
}

// GetShuffleInt64 获取当前轮洗牌后的完整牌序（int64 版本，返回副本）
func (q *ShuffleQueue) GetShuffleInt64() []int64 {
	ret := make([]int64, len(q.shuffle))
	for i, v := range q.shuffle {
		ret[i] = int64(v)
	}
	return ret
}
