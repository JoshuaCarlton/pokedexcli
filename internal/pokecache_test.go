package internal

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "dosent really matter",
			val: []byte("test data"),
		},
		{
			key: "http://somewhere/on/the/internet",
			val: []byte("data from somewhere"),
		},
	}
	for _, c := range cases {
		myCache := NewCache(interval)
		myCache.Add(c.key, c.val)
		val, ok := myCache.Get(c.key)
		if !ok {
			t.Errorf("expected to get a value")
			t.Fail()
		}
		if string(val) != string(c.val) {
			t.Errorf("expected value to equal input")
			t.Fail()
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
