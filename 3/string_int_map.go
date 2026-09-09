package main

import "maps"

type StringIntMap struct {
	source map[string]int
}

func NewStringIntMap() StringIntMap {
	return StringIntMap{
		source: make(map[string]int),
	}
}

func (m StringIntMap) Add(key string, value int) {
	m.source[key] = value
}

func (m StringIntMap) Remove(key string) {
	delete(m.source, key)
}

func (m StringIntMap) Copy() map[string]int {
	return maps.Clone(m.source)
}

func (m StringIntMap) Exists(key string) bool {
	_, ok := m.source[key]
	return ok
}

func (m StringIntMap) Get(key string) (int, bool) {
	v, ok := m.source[key]
	return v, ok
}
