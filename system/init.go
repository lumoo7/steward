package system

import (
	"fmt"
	"reflect"
	"steward/config"
	"steward/system/db"
	"steward/system/web"
	"sync"
)

type ModelCache struct {
	Lock  sync.Mutex
	Cache []any
}

var cache ModelCache

// Push 推入缓存
func (c *ModelCache) Push(m any) {
	fmt.Printf("cache push %v\n", reflect.TypeOf(m))
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
	cache.Push(m)
}

// getModelName 获取 model 名称
func getModelName(m any) string {
	return reflect.TypeOf(m).Elem().Name()
}

// InitModel 初始化 model
func InitModel() {

	fmt.Printf("cache len: %d\n", len(cache.Cache))

	for _, i := range cache.Cache {
		if err := db.DB().AutoMigrate(i); err != nil {
			panic(fmt.Errorf("auto migrate model '%s' err:%w", getModelName(i), err))
		}
	}
}

// Init 全局初始化
func Init() {
	config.InitConfig()
	db.InitMySQL()
	InitModel()
	web.InitWeb()
}
