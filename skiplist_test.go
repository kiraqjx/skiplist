package skiplist

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	s := getTestData()
	value := s.Get([]byte("hello"))
	assert.Equal(t, value, []byte("world"))
}

func TestDelete(t *testing.T) {
	s := getTestData()
	s.Delete([]byte("hello"))
	value := s.Get([]byte("hello"))
	assert.Equal(t, value, []byte(nil))
}

func TestPut(t *testing.T) {
	s := getTestData()
	s.Put([]byte("hello"), []byte("world1"))
	value := s.Get([]byte(("hello")))
	assert.Equal(t, value, []byte(("world1")))
}

func getTestData() SkipList {
	rand.Seed(time.Now().UnixNano())

	s := SkipListBuilder(7)

	dataNodes := make([]*Node, 7)

	foot := s.header.next[0]
	dataNodes[0] = foot

	data := Node{
		key:    []byte("hello"),
		value:  []byte("world"),
		is_del: false,
		next:   dataNodes,
	}

	s.header.next[0] = &data
	return s
}
