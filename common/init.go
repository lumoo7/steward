package common

import (
	"fmt"
	"reflect"
	"steward/common/module"
	"sync"
)

type ModelCache struct {
	Lock  sync.Mutex
	Cache []any
}

var cache ModelCache

// Push 推入缓存
func (c *ModelCache) Push(m any) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Cache = append(c.Cache, m)
}

// Pop 抛出缓存
func (c *ModelCache) Pop() any {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	res := c.Cache[len(c.Cache)-1]
	c.Cache = c.Cache[:len(c.Cache)-2]
	return res
}

// RegisterModel 注册 model
func RegisterModel(m any) {
	go cache.Push(m)
}

// getModelName 获取 model 名称
func getModelName(m any) string {
	return reflect.TypeOf(m).Elem().Name()
}

// InitModel 初始化 model
func InitModel() {
	for _, i := range cache.Cache {
		if err := module.DB().AutoMigrate(i); err != nil {
			panic(fmt.Errorf("auto migrate model '%s' err:%w", getModelName(i), err))
		}
	}
}
