package timebasedkvstore

import (
	"testing"
)

type TimeBasedKVStore struct {
	Item map[string][]Pair
}

type Pair struct {
	Value     string
	Timestamp int
}

func NewTimeBasedKVStore() TimeBasedKVStore {
	return TimeBasedKVStore{
		Item: make(map[string][]Pair),
	}
}

func (kv TimeBasedKVStore) Set(key, value string, timestamp int) {
	if _, ok := kv.Item[key]; !ok {
		kv.Item[key] = []Pair{}
	}
	kv.Item[key] = append(kv.Item[key], Pair{value, timestamp})
}

func (kv TimeBasedKVStore) Get(key string, timestamp int) string {
	var mostRecent string

	if pairs, ok := kv.Item[key]; ok {
		left := 0
		right := len(pairs) - 1

		for left <= right {
			mid := (left + right) / 2
			pair := pairs[mid]
			if pair.Timestamp <= timestamp {
				mostRecent = pair.Value
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return mostRecent
}

func TestTimeBasedKVStore(t *testing.T) {
	kvStore := NewTimeBasedKVStore()
	expected := "bar"
	kvStore.Set("foo", expected, 1)
	val := kvStore.Get("foo", 1)
	if val != expected {
		t.Errorf("expected %s got %s", expected, val)
	}

	expected = "bar2"
	kvStore.Set("foo", expected, 3)
	val = kvStore.Get("foo", 4)
	if val != expected {
		t.Errorf("expected %s got %s", expected, val)
	}

	expected = "bar3"
	kvStore.Set("foo", expected, 4)
	val = kvStore.Get("foo", 5)
	if val != expected {
		t.Errorf("expected %s got %s", expected, val)
	}
}
