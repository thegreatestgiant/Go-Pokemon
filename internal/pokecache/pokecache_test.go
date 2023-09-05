package pokecache

import (
	"reflect"
	"testing"
	"time"
)

func TestCache_Get(t *testing.T) {
	type fields struct {
		cache map[string]cacheEntry
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				cache: tt.fields.cache,
			}
			got, got1 := c.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Cache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCache_Add(t *testing.T) {
	type fields struct {
		cache map[string]cacheEntry
	}
	type args struct {
		key string
		val []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"Test for add",
			fields{
				make(map[string]cacheEntry),
			},
			args{
				"key1",
				[]byte("val1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				cache: tt.fields.cache,
			}
			c.Add(tt.args.key, tt.args.val)
		})
	}
}

func TestCache_reapLoop(t *testing.T) {
	type fields struct {
		cache map[string]cacheEntry
	}
	type args struct {
		interval time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				cache: tt.fields.cache,
			}
			c.reapLoop(tt.args.interval)
		})
	}
}

func TestCache_reap(t *testing.T) {
	type fields struct {
		cache map[string]cacheEntry
	}
	type args struct {
		interval time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				cache: tt.fields.cache,
			}
			c.reap(tt.args.interval)
		})
	}
}

func TestNewCache(t *testing.T) {
	type args struct {
		dur time.Duration
	}
	tests := []struct {
		name string
		args args
		want Cache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCache(tt.args.dur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
