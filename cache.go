package cache

import "time"

type journal struct {
	mark string
	dl   time.Time
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

	if value.dl.IsZero() {
		delete(c.Marks, key)
		return "", false
	}

	return value.mark, true
}

func (c *Cache) Put(key, value string) {
	mark := journal{mark: value}
	c.Marks[key] = mark
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0, len(c.Marks))

	for k := range c.Marks {
		_, ok := c.Get(k)
		if ok {
			keys = append(keys, k)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.Marks[key] = journal{value, deadline}
}
