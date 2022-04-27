package skiplist

import (
	"bytes"
	"encoding/binary"
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

func TestBenchData(t *testing.T) {
	s := getTestData()
	for i := 0; i < 10000; i++ {
		byteInt := intToBytes(i)
		s.Put(byteInt, byteInt)
	}

	for i := 0; i < 10000; i++ {
		if i%7 == 0 {
			s.Delete(intToBytes(i))
		}
	}

	for i := 0; i < 10000; i++ {
		value := s.Get(intToBytes(i))

		if i%7 == 0 {
			assert.Equal(t, value, []byte(nil))
		} else {
			assert.Equal(t, value, intToBytes(i))
		}
	}
}

func intToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
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
