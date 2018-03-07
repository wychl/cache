package cache

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	IsExist(key string) bool
	ClearAll() error
}