package main

import (
	"sync"
	"time"
)

const TTL = 300

type Value struct {
	time int64
	TTL  int
}

type TTLMap struct {
	sync.RWMutex
	m map[string]Value
}

func addEntry(dataMap TTLMap, key string, value int64) {
	dataMap.Lock()
	dataMap.m[key] = Value{value, TTL}
	dataMap.Unlock()
}

func delEntry(dataMap TTLMap, key string) {
	dataMap.Lock()
	delete(dataMap.m, key)
	dataMap.Unlock()
}

func getEntry(dataMap TTLMap, key string) int64 {
	dataMap.RLock()
	v, ok := dataMap.m[key]
	dataMap.RUnlock()
	if !ok {
		return 0
	}
	return v.time
}

func expireEnteries(dataMap TTLMap) bool {
	for {
		time.Sleep(10 * time.Second)
		dataMap.Lock()
		for k, v := range dataMap.m {
			if v.TTL <= 0 {
				delete(dataMap.m, k)
			} else {
				v.TTL = v.TTL - 10
			}
		}
		dataMap.Unlock()
	}
}