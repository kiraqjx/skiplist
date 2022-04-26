package skiplist

import (
	"math/rand"
	"strings"
)

const MIN_LAYER_COUNT = 500

type SkipList struct {
	layer  int
	count  int64
	header *Node
}

type Node struct {
	key    string
	value  string
	is_del bool
	next   []*Node
}

func SkipListBuilder(layer int) SkipList {
	layerData := make([]*Node, layer)

	foot := Node{
		key:    "",
		value:  "",
		is_del: false,
	}
	for i := 0; i < layer; i++ {
		layerData[i] = &foot
	}

	head := Node{
		key:    "",
		value:  "",
		is_del: false,
		next:   layerData,
	}

	return SkipList{
		layer:  layer,
		header: &head,
	}
}

// get value from skip list by key
func (skiplist *SkipList) Get(key string) string {
	node := skiplist.getNode(key)
	if node == nil {
		return ""
	}
	return node.value
}

// delete value from skip list by key
func (skiplist *SkipList) Delete(key string) {
	node := skiplist.getNode(key)
	if node != nil {
		node.is_del = true
	}
	skiplist.count--
}

// put value by key
func (skiplist *SkipList) Put(key string, value string) {
	beforeNodes := skiplist.putPre(key)

	layerData := make([]*Node, skiplist.layer)

	addNode := Node{
		key:    key,
		value:  value,
		is_del: false,
		next:   layerData,
	}

	for i := 0; i < skiplist.layer-1; i++ {
		if i != 0 && skiplist.count < MIN_LAYER_COUNT {
			return
		}
		if i != 0 && rand.Intn(2) != 0 {
			return
		}

		layerBeforeNode := beforeNodes[i]
		layerNextNode := layerBeforeNode.next[i]
		if strings.Compare(layerNextNode.key, key) == 0 {
			layerNextNode.is_del = false
			layerNextNode.value = value
		} else {
			layerBeforeNode.next[i] = &addNode
			addNode.next[i] = layerNextNode
		}
	}
	skiplist.count++
}

// get node from skip list by key
func (skiplist *SkipList) getNode(key string) *Node {
	now_layer := skiplist.layer - 1
	now_node := skiplist.header.next[now_layer]
	before_one_node := skiplist.header

	for {
		if strings.Compare(now_node.key, "") == 0 {
			if now_layer == 0 {
				return nil
			} else {
				now_layer--
				now_node = skiplist.header.next[now_layer]
				continue
			}
		}

		if now_node.is_del {
			now_node = now_node.next[now_layer]
			continue
		}

		is_eq := strings.Compare(now_node.key, key)

		if is_eq == 0 {
			if now_node.is_del {
				return nil
			}
			return now_node
		}

		if is_eq == 1 {
			now_layer--
			now_node = before_one_node.next[now_layer]
			continue
		} else {
			before_one_node = now_node
			now_node = now_node.next[now_layer]
		}
	}
}

// put prepare from getting before-node in every layer
func (skiplist *SkipList) putPre(key string) []*Node {
	beforeEveryLayer := make([]*Node, skiplist.layer)
	now_layer := skiplist.layer - 1
	now_node := skiplist.header.next[now_layer]
	before_one_node := skiplist.header

	for {
		if strings.Compare(now_node.key, "") == 0 {
			if now_layer == 0 {
				return beforeEveryLayer
			} else {
				beforeEveryLayer[now_layer] = before_one_node
				now_layer--
				now_node = skiplist.header.next[now_layer]
				continue
			}
		}

		if now_node.is_del {
			now_node = now_node.next[now_layer]
			continue
		}

		is_eq := strings.Compare(now_node.key, key)

		if is_eq == 0 {
			beforeEveryLayer[now_layer] = before_one_node
			if now_layer == 0 {
				return beforeEveryLayer
			}
			now_layer--
			now_node = before_one_node.next[now_layer]
			continue
		}

		if is_eq == 1 {
			beforeEveryLayer[now_layer] = before_one_node
			now_layer--
			now_node = before_one_node.next[now_layer]
			continue
		} else {
			before_one_node = now_node
			now_node = now_node.next[now_layer]
		}
	}
}
