package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var cacheAdapter = cache.New(5*time.Minute, 10*time.Minute)

func SetCache(k string, x interface{}, d time.Duration) {
	cacheAdapter.Set(k, x, d)
}

func GetCache(k string) (interface{}, bool) {
	return cacheAdapter.Get(k)
}

// SetDefaultCache 设置cache 无时间参数
func SetDefaultCache(k string, x interface{}) {
	cacheAdapter.SetDefault(k, x)
}

// DeleteCache 删除 cache
func DeleteCache(k string) {
	cacheAdapter.Delete(k)
}

// AddCache Add() 加入缓存
func AddCache(k string, x interface{}, d time.Duration) {
	cacheAdapter.Add(k, x, d)
}

// IncrementIntCache 对已存在的key 值自增n
func IncrementIntCache(k string, n int) (num int, err error) {
	return cacheAdapter.IncrementInt(k, n)
}
