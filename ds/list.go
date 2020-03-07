package ds

import "sync"

type ZSet struct {
	list  SortedSet
	rwlock sync.RWMutex
}

func (zs *ZSet) ZRange(start, end int64, f func(float64, int64, interface{})) {
	zs.rwlock.RLock()
	defer zs.rwlock.RUnlock()
	zs.list.Range(start , end,  f)
}
func ( zs*ZSet) ZAdd(score float64, key int64, dat interface{}) {
	zs.rwlock.Lock()
	defer zs.rwlock.Unlock()
	zs.list.Set(score , key , dat)
}
