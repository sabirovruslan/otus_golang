package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CacheTestSuit struct {
	suite.Suite
}

func TestCacheTestSuit(t *testing.T) {
	suite.Run(t, new(CacheTestSuit))
}

func (s *CacheTestSuit) TestCreateInstance() {
	c := NewCache(3)
	s.IsType(c, &lruCache{})
}

func (s *CacheTestSuit) TestSetAndGet() {
	c := NewCache(3)
	s.False(c.Set("1", 1))
	s.False(c.Set("2", 2))
	s.False(c.Set("3", 3))

	_, ok := c.Get("4")
	s.False(ok)

	s.True(c.Set("1", 1))

	_, ok = c.Get("1")
	s.True(ok)
	_, ok = c.Get("2")
	s.True(ok)
	_, ok = c.Get("3")
	s.True(ok)

	s.False(c.Set("4", 4))
	_, ok = c.Get("4")
	s.True(ok)

	_, ok = c.Get("1")
	s.False(ok)

	c.Clear()
	s.False(c.Set("4", 4))

}

func (s *CacheTestSuit) TestValues() {
	c := NewCache(5)

	wasInCache := c.Set("aaa", 100)
	s.False(wasInCache)

	wasInCache = c.Set("bbb", 200)
	s.False(wasInCache)

	val, ok := c.Get("aaa")
	s.True(ok)
	s.Equal(100, val)

	val, ok = c.Get("bbb")
	s.True(ok)
	s.Equal(200, val)

	wasInCache = c.Set("aaa", 300)
	s.True(wasInCache)

	val, ok = c.Get("aaa")
	s.True(ok)
	s.Equal(300, val)

	val, ok = c.Get("ccc")
	s.False(ok)
	s.Nil(val)
}
