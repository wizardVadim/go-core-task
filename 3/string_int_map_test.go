package main

import (
	"fmt"
	"maps"
	"testing"
)

func TestNewStringIntMap(t *testing.T) {
	tests := []struct {
		name     string
		wantType string
	}{
		{name: "Best way", wantType: "map[string]int"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()

			mapSrcType := fmt.Sprintf("%T", m.source)
			if mapSrcType != tt.wantType {
				t.Errorf("TestNewStringIntMap, got = %s, want = %s", mapSrcType, tt.wantType)
			}
		})
	}
}

func mustStringIntMap(t *testing.T) StringIntMap {
	t.Helper()

	return NewStringIntMap()
}

func TestAdd(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	tests := []struct {
		name       string
		inputKey   string
		inputValue int
	}{
		{name: "New key", inputKey: "key", inputValue: 100},
		{name: "Overwrite existing key", inputKey: "existing", inputValue: 200},
		{name: "Empty key", inputKey: "", inputValue: 100},
		{name: "Zero value", inputKey: "key", inputValue: 0},
		{name: "Negative value", inputKey: "key", inputValue: -100},
		{name: "Maximum int", inputKey: "key", inputValue: maxInt},
		{name: "Minimum int", inputKey: "key", inputValue: minInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mustStringIntMap(t)
			before := map[string]int{"existing": 42, "other": 7}
			for key, value := range before {
				m.source[key] = value
			}

			wantLen := len(before)
			if _, exists := before[tt.inputKey]; !exists {
				wantLen++
			}

			m.Add(tt.inputKey, tt.inputValue)

			if val, ok := m.source[tt.inputKey]; !ok || val != tt.inputValue {
				t.Errorf("source[%q] = (%d, %t), want (%d, true)", tt.inputKey, val, ok, tt.inputValue)
			}

			if len(m.source) != wantLen {
				t.Errorf("len(source) = %d, want %d", len(m.source), wantLen)
			}

			for key, value := range before {
				if key == tt.inputKey {
					continue
				}
				if val, ok := m.source[key]; !ok || val != value {
					t.Errorf("source[%q] = (%d, %t), want (%d, true)", key, val, ok, value)
				}
			}

			expected := make(map[string]int, len(m.source))
			for key, value := range m.source {
				expected[key] = value
			}
			m.Add(tt.inputKey, tt.inputValue)
			if !maps.Equal(m.source, expected) {
				t.Errorf("repeated Add changed source: got %v, want %v", m.source, expected)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name     string
		inputKey string
		initial  map[string]int
	}{
		{name: "Best way", inputKey: "key", initial: map[string]int{"key": 100, "other": 42}},
		{name: "Empty key", inputKey: "", initial: map[string]int{"": 1, "other": 42}},
		{name: "Not contains key", inputKey: "missing", initial: map[string]int{"key": 100, "": 1}},
		{name: "Zero value", inputKey: "key", initial: map[string]int{"key": 0, "other": 42}},
		{name: "Only key", inputKey: "key", initial: map[string]int{"key": 100}},
		{name: "Empty map", inputKey: "key", initial: map[string]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mustStringIntMap(t)
			for key, value := range tt.initial {
				m.source[key] = value
			}

			wantLen := len(tt.initial)
			if _, exists := tt.initial[tt.inputKey]; exists {
				wantLen--
			}

			m.Remove(tt.inputKey)

			if _, exists := m.source[tt.inputKey]; exists {
				t.Errorf("source[%q] still exists after Remove", tt.inputKey)
			}
			if len(m.source) != wantLen {
				t.Errorf("len(source) = %d, want %d", len(m.source), wantLen)
			}

			for key, value := range tt.initial {
				if key == tt.inputKey {
					continue
				}
				if val, ok := m.source[key]; !ok || val != value {
					t.Errorf("source[%q] = (%d, %t), want (%d, true)", key, val, ok, value)
				}
			}

			expected := maps.Clone(m.source)
			m.Remove(tt.inputKey)
			if !maps.Equal(m.source, expected) {
				t.Errorf("repeated Remove changed source: got %v, want %v", m.source, expected)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name    string
		initial map[string]int
		want    map[string]int
	}{
		{
			name: "Best way",
			initial: map[string]int{
				"key":  1,
				"key2": 2,
				"key3": 3,
			},
			want: map[string]int{
				"key":  1,
				"key2": 2,
				"key3": 3,
			},
		},
		{
			name: "With removing element",
			initial: map[string]int{
				"key":  1,
				"key2": 2,
				"key3": 3,
			},
			want: map[string]int{
				"key":  1,
				"key2": 2,
				"key3": 3,
			},
		},
		{
			name:    "Nil map",
			initial: nil,
			want:    nil,
		},
		{
			name:    "Empty map",
			initial: map[string]int{},
			want:    map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mustStringIntMap(t)
			for key, value := range tt.initial {
				m.source[key] = value
			}

			copied := m.Copy()
			if !maps.Equal(copied, tt.want) {
				t.Errorf("TestCopy unexpected copy, got = %v, want = %v", copied, tt.want)
			}
			if tt.name == "With removing element" {
				m.Remove("key2")
				_, ok := copied["key2"]
				if !ok {
					t.Errorf("TestCopy not a copy of origin array")
				}
			}
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		name    string
		initial map[string]int
		key     string
		want    bool
	}{
		{
			name: "Exists",
			initial: map[string]int{
				"key": 1,
			},
			key:  "key",
			want: true,
		},
		{
			name: "Not exists",
			initial: map[string]int{
				"key": 1,
			},
			key:  "key2",
			want: false,
		},
		{
			name:    "Empty map",
			initial: map[string]int{},
			key:     "key",
			want:    false,
		},
		{
			name: "Empty key",
			initial: map[string]int{
				"": 10,
			},
			key:  "",
			want: true,
		},
		{
			name:    "Nil map",
			initial: nil,
			key:     "key",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mustStringIntMap(t)
			for key, value := range tt.initial {
				m.source[key] = value
			}

			res := m.Exists(tt.key)
			if res != tt.want {
				t.Errorf("TestExists invalid result: got = %v, want = %v", res, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name      string
		initial   map[string]int
		key       string
		wantValue int
		wantOK    bool
	}{
		{name: "Existing key", initial: map[string]int{"key": 100, "other": 42}, key: "key", wantValue: 100, wantOK: true},
		{name: "Missing key", initial: map[string]int{"other": 42}, key: "key", wantValue: 0, wantOK: false},
		{name: "Zero value", initial: map[string]int{"key": 0}, key: "key", wantValue: 0, wantOK: true},
		{name: "Negative value", initial: map[string]int{"key": -100}, key: "key", wantValue: -100, wantOK: true},
		{name: "Empty key", initial: map[string]int{"": 10}, key: "", wantValue: 10, wantOK: true},
		{name: "Missing empty key", initial: map[string]int{"key": 100}, key: "", wantValue: 0, wantOK: false},
		{name: "Empty map", initial: map[string]int{}, key: "key", wantValue: 0, wantOK: false},
		{name: "Nil map", initial: nil, key: "key", wantValue: 0, wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := StringIntMap{source: maps.Clone(tt.initial)}

			value, ok := m.Get(tt.key)
			if value != tt.wantValue || ok != tt.wantOK {
				t.Errorf("Get(%q) = (%d, %t), want (%d, %t)", tt.key, value, ok, tt.wantValue, tt.wantOK)
			}

			if !maps.Equal(m.source, tt.initial) || (m.source == nil) != (tt.initial == nil) {
				t.Errorf("Get changed source: got %v, want %v", m.source, tt.initial)
			}
		})
	}
}
