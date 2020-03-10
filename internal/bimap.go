package internal

import "sync"

type BiMap struct {
	mu sync.RWMutex
	kv map[interface{}]interface{}
	vk map[interface{}]interface{}
}

func (bi *BiMap) AddPair(key, value interface{}) {
	bi.mu.Lock()
	defer bi.mu.Unlock()

	bi.kv[key] = value
	bi.vk[value] = key
}

func (bi *BiMap) RemoveByKey(key interface{}) {
	bi.mu.Lock()
	defer bi.mu.Unlock()

	if value, ok := bi.kv[key]; ok {
		delete(bi.kv, key)
		delete(bi.vk, value)
	}
}

func (bi *BiMap) RemoveByValue(value interface{}) {
	bi.mu.Lock()
	defer bi.mu.Unlock()

	if key, ok := bi.vk[value]; ok {
		delete(bi.vk, value)
		delete(bi.kv, key)
	}
}

func (bi *BiMap) Value(key interface{}) (value interface{}, ok bool) {
	bi.mu.RLock()
	defer bi.mu.RUnlock()

	value, ok = bi.kv[key]
	return value, ok
}

func (bi *BiMap) Key(value interface{}) (key interface{}, ok bool) {
	bi.mu.RLock()
	defer bi.mu.RUnlock()

	key, ok = bi.vk[value]
	return key, ok
}

func NewBiMap() *BiMap {
	bi := &BiMap{
		kv: make(map[interface{}]interface{}),
		vk: make(map[interface{}]interface{}),
	}
	return bi
}