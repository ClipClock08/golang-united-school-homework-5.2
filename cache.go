package cache

import "time"

type journal struct {
	mark string
}
type Cache struct {
	Marks map[string]journal
}

func NewCache() Cache {
	return Cache{Marks: map[string]journal{}}
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.Marks[key]
	if !ok {
		return "", false
	}

	return value.mark, true
}

func (c *Cache) Put(key, value string) {
	mark := journal{mark: value}
	c.Marks[key] = mark
}

/*func (c *Cache) Keys() []string {
}*/

func (c *Cache) PutTill(key, value string, deadline time.Time) {
}
