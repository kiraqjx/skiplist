package skiplist

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	s := getTestData()
	value := s.Get("hello")
	assert.Equal(t, value, "world")
}

func TestDelete(t *testing.T) {
	s := getTestData()
	s.Delete("hello")
	value := s.Get("hello")
	assert.Equal(t, value, "")
}

func TestPut(t *testing.T) {
	s := getTestData()
	s.Put("hello", "world1")
	value := s.Get("hello")
	assert.Equal(t, value, "world1")
}

func getTestData() SkipList {
	rand.Seed(time.Now().UnixNano())

	s := SkipListBuilder(7)

	dataNodes := make([]*Node, 7)

	foot := s.header.next[0]
	dataNodes[0] = foot

	data := Node{
		key:    "hello",
		value:  "world",
		is_del: false,
		next:   dataNodes,
	}

	s.header.next[0] = &data
	return s
}
